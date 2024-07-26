package enix

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"

	"github.com/gdamore/tcell/v2"
)

type PromptState int

const (
	InText PromptState = iota
	InShadow
	InError
)

// Prompt represents command line prompt.
type Prompt struct {
	Screen tcell.Screen
	Width  int // Screen Width
	Y      int // Tcell Y coordinate of prompt line.

	// History of executed commands.
	History []string

	Colors *cfg.Colorscheme
	Keys   *cfg.Keybindings

	Window *Window

	Line      *line.Line
	Cursor    *cursor.Cursor
	ViewRange line.Range

	ShadowText string

	State PromptState
}

func (p *Prompt) Clear() {
	for x := 0; x < p.Width; x++ {
		p.Screen.SetContent(x, p.Y, ' ', nil, p.Colors.Default)
	}
	p.Screen.Show()
}

func (p *Prompt) ShowError(msg string) {
	x := 0
	for _, r := range msg {
		p.Screen.SetContent(x, p.Y, r, nil, p.Colors.Error)
		x++
	}
	for {
		if x == p.Width {
			break
		}
		p.Screen.SetContent(x, p.Y, ' ', nil, p.Colors.Prompt)
		x++
	}
	p.Screen.Show()
	p.State = InError
}

func (p *Prompt) ShowInfo(msg string) {
	x := 0
	for _, r := range msg {
		p.Screen.SetContent(x, p.Y, r, nil, p.Colors.Default)
		x++
	}
	for {
		if x == p.Width {
			break
		}
		p.Screen.SetContent(x, p.Y, ' ', nil, p.Colors.Default)
		x++
	}
	p.Screen.Show()
	p.State = InError
}

// Currently assume text + shadow text always fits screen width.
func (p *Prompt) Activate(text, shadowText string) {
	p.Line = &line.Line{
		Screen: p.Screen,
		Colors: p.Colors,
		Buf:    text,
	}

	p.Cursor = &cursor.Cursor{
		Screen: p.Screen,
		Colors: p.Colors,
		Line:   p.Line,
		BufIdx: len(text),
	}

	p.ShadowText = shadowText

	p.State = InText
	if len(shadowText) > 0 {
		p.State = InShadow
	}

	p.ViewRange.Lower = 0
	p.ViewRange.Upper = p.Width - 2 - 1

	p.Render()
}

func (p *Prompt) Render() {
	p.Screen.SetContent(0, p.Y, '>', []rune{' '}, p.Colors.Prompt)

	if p.Cursor.BufIdx < p.ViewRange.Lower {
		p.ViewRange.Lower = p.Cursor.BufIdx
		p.ViewRange.Upper = p.ViewRange.Lower + p.Width - 2 - 1
	} else if p.Cursor.BufIdx > p.ViewRange.Upper {
		p.ViewRange.Upper = p.Cursor.BufIdx
		p.ViewRange.Lower = p.ViewRange.Upper - p.Width + 2 + 1
	}

	p.Line.Render(2, p.Y, p.Width-2, p.ViewRange.Lower)

	if len(p.ShadowText) > 0 {
		for i, r := range p.ShadowText {
			p.Screen.SetContent(i+2+p.Line.Len(), p.Y, r, nil, p.Colors.PromptShadow)
		}
	}

	p.Cursor.Render(2, p.Y, p.ViewRange.Lower)

	p.Screen.Show()
}

func (p *Prompt) Backspace() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Backspace()
	}
}

func (p *Prompt) Delete() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Delete()
	}
}

func (p *Prompt) CursorDown() {
	switch p.State {
	case InShadow:
		p.Line.Append(p.ShadowText)
		p.ShadowText = ""
		p.State = InText
	case InText:
		// Implement history handling here.
		panic("unimplemeted")
	}
}

func (p *Prompt) CursorLeft() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Left()
	}
}

func (p *Prompt) CursorRight() {
	switch p.State {
	case InShadow:
		p.Line.Append(p.ShadowText)
		p.Cursor.BufIdx += len(p.ShadowText)
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Right()
	}
}

func (p *Prompt) CursorWordStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.WordStart()
	case InText:
		p.Cursor.WordStart()
	}
}

func (p *Prompt) CursorWordEnd() {
	switch p.State {
	case InShadow:
		p.Line.Append(p.ShadowText)
		p.ShadowText = ""
		p.State = InText
		p.Cursor.WordEnd()
	case InText:
		p.Cursor.WordEnd()
	}
}

func (p *Prompt) CursorLineStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.LineStart()
	case InText:
		p.Cursor.LineStart()
	}
}

func (p *Prompt) Enter() EventReceiver {
	if p.State == InShadow {
		p.Line.Append(p.ShadowText)
		p.ShadowText = ""
		p.State = InText
	}

	return p.Exec()
}

func (p *Prompt) HandleRune(r rune) {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.HandleRune(r)
	case InText:
		p.Cursor.HandleRune(r)
	}
}

func (p *Prompt) RxEvent(ev tcell.Event) EventReceiver {
	switch ev := ev.(type) {
	case *tcell.EventResize:
		p.Window.Resize()
	case *tcell.EventKey:
		switch p.Keys.ToCmd(ev) {
		case "backspace":
			p.Backspace()
		case "del":
			p.Delete()
		case "cursor-down":
			p.CursorDown()
		case "cursor-left":
			p.CursorLeft()
		case "cursor-right":
			p.CursorRight()
		case "cursor-word-start":
			p.CursorWordStart()
		case "cursor-word-end":
			p.CursorWordEnd()
		case "cursor-line-start":
			p.CursorLineStart()
		case "enter":
			return p.Enter()
		case "escape", "quit":
			p.Clear()
			return p.Window
		default:
			switch ev.Key() {
			case tcell.KeyRune:
				p.HandleRune(ev.Rune())
			}
		}
	}

	p.Render()

	return p
}

// Exec executes command.
func (p *Prompt) Exec() EventReceiver {
	cmd, args, _ := strings.Cut(strings.TrimSpace(p.Line.Buf), " ")

	switch cmd {
	case "cmd":
		p.Activate("", "")
		return p
	case "cmd-info":
		p.ShowInfo(args)
		return p.Window
	case "cmd-error":
		p.ShowError(args)
		return p.Window
	default:
		p.ShowError(
			fmt.Sprintf(
				"invalid or unimplemented command '%s', if unimplemented report on https://github.com/m-kru/enix",
				cmd,
			),
		)
		return p.Window
	}

	p.Clear()

	return p
}
