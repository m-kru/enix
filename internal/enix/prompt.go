package enix

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"

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
	Config *cfg.Config
	Colors *cfg.Colorscheme
	Keys   *cfg.Keybindings

	Screen tcell.Screen
	Frame  frame.Frame

	// History of executed commands.
	History []string

	Window *Window

	Line   *line.Line
	Cursor *cursor.Cursor
	View   view.View

	ShadowText string

	State PromptState
}

func (p *Prompt) Clear() {
	for x := 0; x < p.Frame.Width; x++ {
		p.Frame.SetContent(x, 0, ' ', p.Colors.Default)
	}
	p.Screen.Show()
}

func (p *Prompt) ShowError(msg string) {
	x := 0
	for _, r := range msg {
		p.Frame.SetContent(x, 0, r, p.Colors.Error)
		x++
	}
	for {
		if x == p.Frame.Width {
			break
		}
		p.Frame.SetContent(x, 0, ' ', p.Colors.Prompt)
		x++
	}
	p.Screen.Show()
	p.State = InError
}

func (p *Prompt) ShowInfo(msg string) {
	x := 0
	for _, r := range msg {
		p.Frame.SetContent(x, 0, r, p.Colors.Default)
		x++
	}
	for {
		if x == p.Frame.Width {
			break
		}
		p.Frame.SetContent(x, 0, ' ', p.Colors.Default)
		x++
	}
	p.Screen.Show()
	p.State = InError
}

// Currently assume text + shadow text always fits screen width.
func (p *Prompt) Activate(text, shadowText string) {
	p.Line = line.FromString(text)

	p.Cursor = &cursor.Cursor{
		Config: p.Config,
		Line:   p.Line,
		Idx:    len(text),
		BufIdx: len(text),
	}

	p.ShadowText = shadowText

	p.State = InText
	if len(shadowText) > 0 {
		p.State = InShadow
	}

	p.View.Line = 1
	p.View.Column = 1
	p.View.Width = p.Frame.Width - 2 - 1 // - 1 because of cursor
	p.View.Height = 1

	p.Render()
}

func (p *Prompt) Render() {
	p.Frame.SetContent(0, 0, '>', p.Colors.Prompt)
	p.Frame.SetContent(1, 0, ' ', p.Colors.Prompt)

	if !p.View.IsVisible(p.Cursor.View()) {
		p.View = p.View.MinAdjust(p.Cursor.View())
	}

	p.Line.Render(p.Config, p.Colors, p.Frame.Line(2, 0), p.View)

	if len(p.ShadowText) > 0 {
		for i, r := range p.ShadowText {
			p.Frame.SetContent(i+2+p.Line.Len(), 0, r, p.Colors.PromptShadow)
		}
	}

	p.Cursor.Render(p.Config, p.Colors, p.Frame.Line(2, 0), p.View)

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

func (p *Prompt) Down() {
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

func (p *Prompt) Left() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
	case InText:
		p.Cursor.Left()
	}
}

func (p *Prompt) Right() {
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

func (p *Prompt) WordStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.WordStart()
	case InText:
		p.Cursor.WordStart()
	}
}

func (p *Prompt) WordEnd() {
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

func (p *Prompt) LineStart() {
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
		p.Window.Render()
	case *tcell.EventKey:
		cmd, _ := p.Keys.ToCmd(ev)

		switch cmd {
		case "backspace":
			p.Backspace()
		case "del":
			p.Delete()
		case "down":
			p.Down()
		case "left":
			p.Left()
		case "right":
			p.Right()
		case "word-start":
			p.WordStart()
		case "word-end":
			p.WordEnd()
		case "line-start":
			p.LineStart()
		case "enter":
			return p.Enter()
		case "esc", "quit":
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
	name, args, _ := strings.Cut(strings.TrimSpace(p.Line.String()), " ")

	var err error
	var ret EventReceiver = p

	switch name {
	case "":
		// Do nothing
	case "cmd":
		p.Activate("", "")
		return p
	case "cmd-info":
		p.ShowInfo(args)
		return p.Window
	case "cmd-error":
		p.ShowError(args)
		return p.Window
	case "down":
		err = cmd.Down(args, p.Window.CurrentTab)
		ret = p.Window
	case "left":
		err = cmd.Left(args, p.Window.CurrentTab)
		ret = p.Window
	case "right":
		err = cmd.Right(args, p.Window.CurrentTab)
		ret = p.Window
	case "up":
		err = cmd.Up(args, p.Window.CurrentTab)
		ret = p.Window
	case "end":
		err = cmd.End(args, p.Window.CurrentTab)
		ret = p.Window
	case "cursor-count":
		p.ShowInfo(fmt.Sprintf("%d", p.Window.CurrentTab.Cursors.Count()))
		return p.Window
	case "tab-count":
		p.ShowInfo(fmt.Sprintf("%d", p.Window.Tabs.Count()))
		return p.Window
	case "tab-width":
		err = cmd.CfgTabWidth(args, p.Config)
		ret = p.Window
	default:
		p.ShowError(
			fmt.Sprintf(
				"invalid or unimplemented command '%s', if unimplemented report on https://github.com/m-kru/enix",
				name,
			),
		)
		return p.Window
	}

	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return p.Window
	}

	p.Clear()

	if ret == p.Window {
		p.Window.Render()
	}

	return ret
}
