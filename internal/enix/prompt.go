package enix

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"

	"github.com/gdamore/tcell/v2"
)

type PromptState int

const (
	InText PromptState = iota
	InShadow
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

	Window     *Window
	CurrentTab *Tab

	Line   *line.Line
	Cursor *cursor.Cursor

	ShadowText string

	State PromptState
}

func (p *Prompt) Clear() {
	for x := 0; x < p.Width; x++ {
		p.Screen.SetContent(x, p.Y, ' ', nil, p.Colors.Default)
	}
	p.Screen.Show()
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
		X:      2 + len(text) + len(shadowText),
		Y:      p.Y,
		Colors: p.Colors,
		Line:   p.Line,
		BufIdx: len(text),
	}

	p.ShadowText = shadowText

	p.State = InText
	if len(shadowText) > 0 {
		p.State = InShadow
	}

	p.Render()
}

func (p *Prompt) Render() {
	p.Screen.SetContent(0, p.Y, '>', []rune{' '}, p.Colors.Prompt)

	p.Line.Render(2, p.Y, p.Width, 0)

	if len(p.ShadowText) > 0 {
		for i, r := range p.ShadowText {
			p.Screen.SetContent(i+2+p.Line.Len(), p.Y, r, nil, p.Colors.PromptShadow)
		}
	}

	p.Cursor.Render()

	p.Screen.Show()
}

func (p *Prompt) HandleBackspace() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.X = p.Line.Len() + 2
	case InText:
		p.Cursor.HandleBackspace()
	}
}

func (p *Prompt) HandleLeft() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.X = p.Line.Len() + 2
	case InText:
		p.Cursor.HandleLeft()
	}
}

func (p *Prompt) HandleRight() {
	switch p.State {
	case InShadow:
		// Do nothing
	case InText:
		p.Cursor.HandleRight()
	}
}

func (p *Prompt) HandleRune(r rune) {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.X = p.Line.Len() + 2
		p.Cursor.HandleRune(r)
	case InText:
		p.Cursor.HandleRune(r)
	}
}

func (p *Prompt) RxEvent(ev tcell.Event) EventReceiver {
	switch ev := ev.(type) {
	/*
		case *tcell.EventResize:
			p.Window.Render()
	*/
	case *tcell.EventKey:
		switch p.Keys.ToCmd(ev) {
		case "escape":
			p.Clear()
			return p.Window
		case "quit":
			return nil
		default:
			switch ev.Key() {
			case tcell.KeyBackspace | tcell.KeyBackspace2:
				p.HandleBackspace()
			case tcell.KeyLeft:
				p.HandleLeft()
			case tcell.KeyRight:
				p.HandleRight()
			case tcell.KeyRune:
				p.HandleRune(ev.Rune())
			}
		}
	}

	p.Render()

	return p
}
