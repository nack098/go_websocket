package ui

import (
	"fmt"
	result "go_websocket/type"
	"go_websocket/ui/platform"
)

type uiState struct {
	current       UI
	previous      UI
	surfaceWidth  int
	surfaceHeight int
}

type UI interface {
	HandleInput(*string)
	Update()
	Render()

	setRender(bool)
	get() UI
}

type UIEnum int

var ui *uiState = nil

const (
	MAIN_MENU UIEnum = iota
)

func getUI(currentUI UIEnum) (UI, error) {
	switch currentUI {
	case MAIN_MENU:
		return MainMenu.get(), nil
	default:
		return nil, fmt.Errorf("unknown menu state")
	}
}

func Init(_ any) result.Result {
	width, height, err := platform.GetTermSize()
	if err != nil {
		return result.Err(fmt.Errorf("could not get terminal size: %v", err))
	}
	menu, _ := getUI(MAIN_MENU)

	ui = &uiState{
		current:       menu,
		previous:      menu,
		surfaceWidth:  width,
		surfaceHeight: height,
	}

	return result.Ok(nil)
}

func Resize(_ any) result.Result {
	width, height, err := platform.GetTermSize()
	if err != nil {
		return result.Err(fmt.Errorf("could not get terminal size: %v", err))
	}

	ui.surfaceWidth = width
	ui.surfaceHeight = height

	fmt.Print("\033[2J")
	ui.current.setRender(false)
	return result.Ok(nil)
}

func Update() {
	ui.current.Update()
}

func Render() {
	ui.current.Render()
}

func HandleInput(key *string) {
	ui.current.HandleInput(key)
}
