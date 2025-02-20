package ui

import (
	"fmt"
	"strings"
)

type mainMenu struct {
	index int
	items []menuItem
}

func (m *mainMenu) increment() {
	if m.index+1 >= len(m.items) {
		m.index = 0
	} else {
		m.index++
	}
}

func (m *mainMenu) decrement() {
	if m.index-1 < 0 {
		m.index = len(m.items) - 1
	} else {
		m.index--
	}
}

func (m *mainMenu) HandleInput(key *string) {
	switch *key {
	case "\x1b[A", "k": // Up
		m.increment()
	case "\x1b[B", "j": // Down
		m.decrement()
	case "\r", "\n":
		m.items[m.index].action() // Enter
	}
}

func (m *mainMenu) banner() string {
	banner := `+-------------------------+
|  🕮 Chat room simulator  |
+-------------------------+`
	buf_split := strings.Split(banner, "\n")
	var buf string
	for i, line := range buf_split {
		buf += fmt.Sprintf("\033[%d;%dH\033[2K%s", ui.surfaceHeight/2+i+1-len(buf_split), ui.surfaceWidth/2-14, line)
	}
	return buf
}

func (m *mainMenu) selection() string {
	var buf string
	for i, item := range m.items {
		str := *item.getName()
		if i == m.index {
			// str = "☛  " + str
			str = ">>" + str + "<<"
		}
		buf += fmt.Sprintf("\033[%d;%dH\033[2K%s", ui.surfaceHeight/2+i+3-len(m.items), (ui.surfaceWidth-len(str))/2-2, str)
	}
	return buf
}

func (m *mainMenu) renderItems() {
	// First render
	if !isBannerRender {
		fmt.Print(m.banner())
		fmt.Print(m.selection())
		isBannerRender = true
		lastSelection = m.index
		return
	}

	// Updating
	if m.index != lastSelection {
		fmt.Print(m.selection())
		lastSelection = m.index
	}
}

func (m *mainMenu) Update() {
}

func (m *mainMenu) Render() {
	m.renderItems()
}

func (m *mainMenu) get() UI {
	if MainMenu == nil {
		MainMenu = &mainMenu{
			index: 0,
			items: []menuItem{
				&startButton{
					"Start",
				},
				&quitButton{
					"Exit",
				},
			},
		}
	}
	return MainMenu
}

var (
	MainMenu       *mainMenu = nil
	isBannerRender           = false
	lastSelection            = 0
)
