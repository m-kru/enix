package lang

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

// One of bufIdx or runeIdx is not required, not sure yet which one.
type matchingDelimiterPosition struct {
	lineNum int
	bufIdx  int
	runeIdx int
}

type Highlighter struct {
	Regions []*Region

	matchingDelims []matchingDelimiterPosition

	lineNum         int
	firstVisLineNum int
	lastVisLineNum  int

	region *Region // Current region

	startTokens []RegionToken // Region start tokens detected in a given line.
}

// DefaultHighlighter returns a highlighter highlighting only a cursor word.
func DefaultHighlighter() *Highlighter {
	return &Highlighter{
		Regions:         []*Region{DefaultRegion()},
		matchingDelims:  nil,
		lineNum:         0,
		firstVisLineNum: 0,
		lastVisLineNum:  0,
		region:          nil,
		startTokens:     []RegionToken{},
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

	hl, err := langDefIntoHighlighter(langDef)
	if err != nil {
		return DefaultHighlighter(),
			fmt.Errorf("%s highlighter: %v", lang, err)
	}

	return hl, nil
}

// Resets highlighter for a new analysis.
func (hl *Highlighter) reset(
	cursors []*cursor.Cursor,
	firstVisLineNum int,
	lastVisLineNum int,
) {
	hl.matchingDelims = make([]matchingDelimiterPosition, 0, 8)
	hl.lineNum = 1
	hl.firstVisLineNum = firstVisLineNum
	hl.lastVisLineNum = lastVisLineNum
	hl.region = hl.Regions[0]

	cursorWordRegex := ""

	for _, cur := range cursors {
		var mdCur *cursor.Cursor

		r := cur.Rune()
		switch r {
		case '[', ']':
			mdCur = cur.MatchBracket(hl.firstVisLineNum, hl.lastVisLineNum)
		case '{', '}':
			mdCur = cur.MatchBracket(hl.firstVisLineNum, hl.lastVisLineNum)
		case '(', ')':
			mdCur = cur.MatchParen(hl.firstVisLineNum, hl.lastVisLineNum)
		}

		if mdCur != nil && firstVisLineNum <= mdCur.LineNum && mdCur.LineNum <= lastVisLineNum {
			mdPos := matchingDelimiterPosition{
				lineNum: mdCur.LineNum,
				bufIdx:  mdCur.Line.BufIdx(mdCur.RuneIdx),
				runeIdx: mdCur.RuneIdx,
			}
			hl.matchingDelims = append(hl.matchingDelims, mdPos)
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
		di := hl.matchingDelims[i]
		dj := hl.matchingDelims[j]

		if di.lineNum < dj.lineNum {
			return true
		} else if di.lineNum == dj.lineNum {
			return di.runeIdx < dj.runeIdx
		}

		return false
	}
	sort.Slice(hl.matchingDelims, less)

	// Reset cursor word regex
	for _, r := range hl.Regions {
		r.CursorWord = nil
	}

	// Compile cursor word regex
	re, err := regexp.Compile(cursorWordRegex)
	if err == nil {
		for _, r := range hl.Regions {
			r.CursorWord = re
		}
	}
}

func (hl *Highlighter) done() bool {
	return hl.lineNum > hl.lastVisLineNum
}

func (hl *Highlighter) analyzeLine(line *line.Line, hls *[]highlight.Highlight) {
	if len(line.Buf) == 0 {
		hl.lineNum++
		return
	}

	if len(hl.Regions) == 1 {
		hl.analyzeLineOneRegionOnly(line, hls)
		return
	}

	startTokensValid := false
	startTokIdx := 0 // Current index of the start token
	var endTokens map[string][]RegionToken
	bufIdx := 0 // Current line buffer index

	for bufIdx < len(line.Buf) {
		region := hl.region
		endIdx := len(line.Buf)

		if hl.region.Name == "Default" {
			if !startTokensValid {
				hl.startTokens = hl.startTokens[:0]
				lineStartTokens(line.Buf, bufIdx, hl.Regions, &hl.startTokens)
				startTokensValid = true
			}
			for i := startTokIdx; i < len(hl.startTokens); i++ {
				tok := hl.startTokens[i]
				if bufIdx < tok.startBufIdx {
					endIdx = tok.startBufIdx
					break
				} else if bufIdx == tok.startBufIdx {
					hl.region = tok.region
					startTokIdx = i + 1
					break
				}
			}
		}

		if hl.region.Name != "Default" {
			region = hl.region

			if endTokens == nil {
				endTokens = make(map[string][]RegionToken)
			}

			var endToks []RegionToken
			var ok bool
			endToks, ok = endTokens[hl.region.Name]
			if !ok {
				endToks = lineEndTokens(line.Buf, bufIdx, hl.region)
				endTokens[hl.region.Name] = endToks
			}

			for _, tok := range endToks {
				if bufIdx < tok.startBufIdx {
					endIdx = tok.endBufIdx
					hl.region = hl.Regions[0]
					break
				}
			}
		}

		if hl.lineNum < hl.firstVisLineNum {
			bufIdx = endIdx
			continue
		}

		hl.highlightRegion(line.Buf, hl.lineNum, bufIdx, endIdx, region, hls)
		bufIdx = endIdx
	}

	hl.lineNum++
}

func (hl *Highlighter) analyzeLineOneRegionOnly(line *line.Line, hls *[]highlight.Highlight) {
	hl.lineNum++

	if hl.lineNum < hl.firstVisLineNum {
		return
	}

	hl.highlightRegion(line.Buf, hl.lineNum, 0, len(line.Buf), hl.Regions[0], hls)
}

func (hl *Highlighter) highlightRegion(
	line []byte,
	lineNum int,
	startBufIdx int,
	endBufIdx int,
	region *Region,
	hls *[]highlight.Highlight,
) {
	matches := region.match(line[startBufIdx:endBufIdx])

	runeOffset := util.ByteIdxToRuneIdx(line, startBufIdx)

	runeIdx := 0 // Current rune index
	for _, m := range matches {
		if runeIdx < m.start {
			hl := highlight.Highlight{
				LineNum:      lineNum,
				StartRuneIdx: runeOffset + runeIdx,
				EndRuneIdx:   runeOffset + m.start,
				Style:        cfg.Style.Get(region.Style),
			}
			*hls = append(*hls, hl)
		}

		hl := highlight.Highlight{
			LineNum:      lineNum,
			StartRuneIdx: runeOffset + m.start,
			EndRuneIdx:   runeOffset + m.end,
			Style:        m.style,
		}
		*hls = append(*hls, hl)
		runeIdx = m.end
	}

	endRuneIdx := util.ByteIdxToRuneIdx(line, endBufIdx)
	if runeIdx < endRuneIdx {
		hl := highlight.Highlight{
			LineNum:      lineNum,
			StartRuneIdx: runeIdx,
			EndRuneIdx:   endRuneIdx,
			Style:        cfg.Style.Get(region.Style),
		}
		*hls = append(*hls, hl)
	}
}
