package enix

import (
	"fmt"
	"strings"
	"unicode"

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
	p.Screen.HideCursor()
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
	p.Frame.SetContent(0, 0, ':', p.Colors.Prompt)

	if !p.View.IsVisible(p.Cursor.View()) {
		p.View = p.View.MinAdjust(p.Cursor.View())
	}

	p.Line.Render(p.Config, p.Colors, p.Frame.Line(1, 0), p.View)

	if len(p.ShadowText) > 0 {
		for i, r := range p.ShadowText {
			p.Frame.SetContent(i+1+p.Line.Len(), 0, r, p.Colors.PromptShadow)
		}
	}

	p.Cursor.Render(p.Config, p.Colors, p.Frame.Line(1, 0), p.View, true)

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

func (p *Prompt) PrevWordStart() {
	switch p.State {
	case InShadow:
		p.ShadowText = ""
		p.State = InText
		p.Cursor.PrevWordStart()
	case InText:
		p.Cursor.PrevWordStart()
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

func (p *Prompt) Enter() TcellEventReceiver {
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
		p.Cursor.InsertRune(r)
	case InText:
		p.Cursor.InsertRune(r)
	}
}

func (p *Prompt) RxTcellEvent(ev tcell.Event) TcellEventReceiver {
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
		case "enter":
			return p.Enter()
		case "esc", "quit":
			p.Clear()
			return p.Window
		case "left":
			p.Left()
		case "line-start":
			p.LineStart()
		case "right":
			p.Right()
		case "prev-word-start":
			p.PrevWordStart()
		case "word-end":
			p.WordEnd()
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
func (p *Prompt) Exec() TcellEventReceiver {
	name, argStr, _ := strings.Cut(strings.TrimSpace(p.Line.String()), " ")
	args := strings.Fields(argStr)
	tab := p.Window.CurrentTab

	var err error
	updateView := true

	// If the first word is a number, then treat it as a goto command.
	if unicode.IsDigit([]rune(name)[0]) {
		err = cmd.Goto(strings.Fields(p.Line.String()), tab)
		goto errorCheck
	}

	switch name {
	case "":
	case "add-cursor":
		err = cmd.AddCursor(args, tab)
		// Do nothing
	case "cmd":
		p.Activate("", "")
		return p
	case "cmd-error":
		p.ShowError(argStr)
		return p.Window
	case "cmd-info":
		p.ShowInfo(argStr)
		return p.Window
	case "cursor-count":
		p.ShowInfo(fmt.Sprintf("%d", tab.Cursors.Count()))
		return p.Window
	case "down":
		err = cmd.Down(args, tab)
	case "end":
		err = cmd.End(args, tab)
	case "goto":
		err = cmd.Goto(args, tab)
	case "left":
		err = cmd.Left(args, tab)
	case "line-end":
		err = cmd.LineEnd(args, tab)
	case "line-start":
		err = cmd.LineStart(args, tab)
	case "right":
		err = cmd.Right(args, tab)
	case "rune":
		err = cmd.Rune(args, tab)
	case "quit", "q":
		err = cmd.Quit(args, tab, false)
		if err == nil {
			return nil
		}
	case "quit!", "q!":
		_ = cmd.Quit(args, tab, true)
		return nil
	case "save":
		err = cmd.Save(args, tab, p.Config.TrimOnSave)
	case "space":
		err = cmd.Space(args, tab)
	case "tab":
		err = cmd.Tab(args, tab)
	case "tab-count":
		p.ShowInfo(fmt.Sprintf("%d", p.Window.Tabs.Count()))
		return p.Window
	case "tab-width":
		err = cmd.CfgTabWidth(args, p.Config)
	case "trim":
		err = cmd.Trim(args, tab)
	case "up":
		err = cmd.Up(args, tab)
	case "view-down":
		err = cmd.ViewDown(args, tab)
		updateView = false
	case "view-up":
		err = cmd.ViewUp(args, tab)
		updateView = false
	case "prev-word-start":
		err = cmd.PrevWordStart(args, tab)
	case "word-end":
		err = cmd.WordEnd(args, tab)
	case "word-start":
		err = cmd.WordStart(args, tab)
	default:
		p.ShowError(
			fmt.Sprintf(
				"invalid or unimplemented command '%s', if unimplemented report on https://github.com/m-kru/enix",
				name,
			),
		)
		return p.Window
	}

errorCheck:
	if err != nil {
		p.ShowError(fmt.Sprintf("%v", err))
		return p.Window
	}

	if updateView {
		tab.UpdateView()
	}

	p.Clear()

	return p.Window
}
