package lang

import (
	"fmt"
	"regexp"
	"sort"
	"unicode/utf8"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
)

type bracketPosition struct {
	lineNum int
	runeIdx int
}

type Highlighter struct {
	Regions    []*Region
	Highlights []highlight.Highlight // Reused on every call to Analyze.

	matchingBrackets []bracketPosition

	lineNum         int
	firstVisLineNum int
	lastVisLineNum  int

	region *Region // Current region

	startTokens []RegionToken // Region start tokens detected in a given line.
}

// DefaultHighlighter returns a highlighter highlighting only a cursor word.
func DefaultHighlighter() *Highlighter {
	return &Highlighter{
		Regions:          []*Region{DefaultRegion()},
		Highlights:       make([]highlight.Highlight, 0, 1024),
		matchingBrackets: nil,
		lineNum:          0,
		firstVisLineNum:  0,
		lastVisLineNum:   0,
		region:           nil,
		startTokens:      []RegionToken{},
	}
}

func NewHighlighter(lang string) (*Highlighter, error) {
	if lang == "" {
		return DefaultHighlighter(), nil
	}

	langDef, err := readFiletypeDefFromJSON(lang)
	if err != nil {
		return DefaultHighlighter(),
			fmt.Errorf("creating highlighter for %s filetype: %v", lang, err)
	}
	if langDef == nil {
		return DefaultHighlighter(), nil
	}

	hlr, err := langDefIntoHighlighter(langDef)
	if err != nil {
		return DefaultHighlighter(),
			fmt.Errorf("%s highlighter: %v", lang, err)
	}

	return hlr, nil
}

// The line argument  must be the first line of file.
// StartLineIdx is the index of the first visible line.
// EndLineIdx is the index of the last visible line.
// Highlights is overwritten by each call to Analyze.
func (hlr *Highlighter) Analyze(
	line *line.Line, // First tab line
	firstVisLineNum int,
	lastVisLineNum int,
	cursors []*cursor.Cursor,
) {
	hlr.Highlights = hlr.Highlights[:0]

	if len(hlr.Regions) == 0 {
		return
	}

	hlr.reset(cursors, firstVisLineNum, lastVisLineNum)

	for !hlr.done() {
		hlr.analyzeLine(line)
		line = line.Next
	}
}

// Resets highlighter for a new analysis.
func (hlr *Highlighter) reset(
	cursors []*cursor.Cursor,
	firstVisLineNum int,
	lastVisLineNum int,
) {
	hlr.matchingBrackets = make([]bracketPosition, 0, 8)
	hlr.lineNum = 1
	hlr.firstVisLineNum = firstVisLineNum
	hlr.lastVisLineNum = lastVisLineNum
	hlr.region = hlr.Regions[0]

	cursorWordRegex := ""

	for _, cur := range cursors {
		var mbCur *cursor.Cursor

		r := cur.Rune()
		switch r {
		case '[', ']':
			mbCur = cur.MatchBracket(hlr.firstVisLineNum, hlr.lastVisLineNum)
		case '{', '}':
			mbCur = cur.MatchCurly(hlr.firstVisLineNum, hlr.lastVisLineNum)
		case '(', ')':
			mbCur = cur.MatchParen(hlr.firstVisLineNum, hlr.lastVisLineNum)
		}

		if mbCur != nil && firstVisLineNum <= mbCur.LineNum && mbCur.LineNum <= lastVisLineNum {
			mbPos := bracketPosition{
				lineNum: mbCur.LineNum,
				runeIdx: mbCur.RuneIdx,
			}
			hlr.matchingBrackets = append(hlr.matchingBrackets, mbPos)
			continue
		}

		cursorWord := ""
		if cfg.Cfg.HighlightCursorWord {
			cursorWord = cur.GetWord()
			if cursorWord == "" {
				continue
			}
			if len(cursorWordRegex) > 0 {
				cursorWordRegex += `|\b`
			} else {
				cursorWordRegex += `\b`
			}
			cursorWordRegex += cursorWord + `\b`

		}
	}

	// Sort matching delimiters
	less := func(i, j int) bool {
		di := hlr.matchingBrackets[i]
		dj := hlr.matchingBrackets[j]

		if di.lineNum < dj.lineNum {
			return true
		} else if di.lineNum == dj.lineNum {
			return di.runeIdx < dj.runeIdx
		}

		return false
	}
	sort.Slice(hlr.matchingBrackets, less)

	// Reuse the compiled regex if the cursor word pattern is unchanged.
	re := hlr.Regions[0].CursorWord
	if cursorWordRegex == "" {
		re = nil
	} else if re == nil || re.String() != cursorWordRegex {
		var err error
		re, err = regexp.Compile(cursorWordRegex)
		if err != nil {
			re = nil
		}
	}

	for _, r := range hlr.Regions {
		r.CursorWord = re
	}
}

func (hlr *Highlighter) done() bool {
	return hlr.lineNum > hlr.lastVisLineNum
}

func (hlr *Highlighter) analyzeLine(line *line.Line) {
	if len(line.Buf) == 0 {
		hlr.lineNum++
		return
	}

	if len(hlr.Regions) == 1 {
		hlr.analyzeLineOneRegionOnly(line)
		return
	}

	startTokensValid := false
	startTokIdx := 0 // Current list index of the start token
	var endTokens map[string][]RegionToken
	bufIdx := 0 // Current line buffer index
	runeOffset := 0
	startTokEndIdx := 0

	for bufIdx < len(line.Buf) {
		region := hlr.region
		endIdx := len(line.Buf)

		if hlr.region.Name == "Default" {
			if !startTokensValid {
				hlr.startTokens = hlr.startTokens[:0]
				lineStartTokens(line.Buf, bufIdx, hlr.Regions, &hlr.startTokens)
				startTokensValid = true
			}
			for i := startTokIdx; i < len(hlr.startTokens); i++ {
				tok := hlr.startTokens[i]
				if bufIdx < tok.startBufIdx {
					endIdx = tok.startBufIdx
					break
				} else if bufIdx == tok.startBufIdx {
					hlr.region = tok.region
					startTokIdx = i + 1
					startTokEndIdx = tok.endBufIdx
					break
				}
			}
		}

		if hlr.region.Name != "Default" {
			region = hlr.region

			if endTokens == nil {
				endTokens = make(map[string][]RegionToken)
			}

			var endToks []RegionToken
			var ok bool
			endToks, ok = endTokens[hlr.region.Name]
			if !ok {
				endToks = lineEndTokens(line.Buf, bufIdx, hlr.region)
				endTokens[hlr.region.Name] = endToks
			}

			for _, tok := range endToks {
				if tok.startBufIdx >= startTokEndIdx || !startTokensValid {
					endIdx = tok.endBufIdx
					hlr.region = hlr.Regions[0]
					break
				}
			}
		}

		if hlr.lineNum < hlr.firstVisLineNum {
			bufIdx = endIdx
			continue
		}

		runeOffset = hlr.highlightRegion(line.Buf, hlr.lineNum, bufIdx, endIdx, runeOffset, region)
		bufIdx = endIdx
	}

	hlr.lineNum++
}

func (hlr *Highlighter) analyzeLineOneRegionOnly(line *line.Line) {
	if hlr.lineNum < hlr.firstVisLineNum {
		hlr.lineNum++
		return
	}

	hlr.highlightRegion(line.Buf, hlr.lineNum, 0, len(line.Buf), 0, hlr.Regions[0])
	hlr.lineNum++
}

// highlightRegion returns the rune index at the end of the region.
func (hlr *Highlighter) highlightRegion(
	line []byte,
	lineNum int,
	startBufIdx int,
	endBufIdx int,
	runeOffset int,
	region *Region,
) int {
	matches := region.match(line[startBufIdx:endBufIdx])

	regStyle := cfg.Style.Get(region.Style)
	runeIdx := 0 // Current rune index

	highlightMatchingBrackets := func(hl highlight.Highlight) highlight.Highlight {
		mb := hlr.matchingBrackets[0]
		for len(hlr.matchingBrackets) > 0 && hl.CoversCell(mb.lineNum, mb.runeIdx) {
			mbhl := highlight.Highlight{
				LineNum:      mb.lineNum,
				StartRuneIdx: mb.runeIdx,
				EndRuneIdx:   mb.runeIdx + 1,
				Style:        cfg.Style.MatchingBracket,
			}
			hls := hl.Split(mbhl)
			for i := range len(hls) - 1 {
				hlr.Highlights = append(hlr.Highlights, hls[i])
			}
			hl = hls[len(hls)-1]
			hlr.matchingBrackets = hlr.matchingBrackets[1:]

			if len(hlr.matchingBrackets) > 0 {
				mb = hlr.matchingBrackets[0]
			}
		}

		return hl
	}

	mIdx := 0
	for mIdx < len(matches) {
		m := matches[mIdx]

		hl := highlight.Highlight{
			LineNum:      lineNum,
			StartRuneIdx: runeOffset + runeIdx,
			EndRuneIdx:   runeOffset + m.start,
			Style:        regStyle,
		}

		if runeIdx == m.start {
			hl.EndRuneIdx = runeOffset + m.end
			hl.Style = m.style
			mIdx++
			runeIdx = m.end
		} else {
			runeIdx = m.start
		}

		if len(hlr.matchingBrackets) > 0 && hlr.matchingBrackets[0].lineNum == lineNum {
			hl = highlightMatchingBrackets(hl)
		}

		hlr.Highlights = append(hlr.Highlights, hl)
	}

	endRuneIdx := runeOffset + utf8.RuneCount(line[startBufIdx:endBufIdx])
	if runeIdx < endRuneIdx {
		hl := highlight.Highlight{
			LineNum:      lineNum,
			StartRuneIdx: runeIdx,
			EndRuneIdx:   endRuneIdx,
			Style:        cfg.Style.Get(region.Style),
		}

		if len(hlr.matchingBrackets) > 0 && hlr.matchingBrackets[0].lineNum == lineNum {
			hl = highlightMatchingBrackets(hl)
		}

		hlr.Highlights = append(hlr.Highlights, hl)
	}

	return endRuneIdx
}
