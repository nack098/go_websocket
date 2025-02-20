package system

import (
	"fmt"
	result "go_websocket/type"
	"go_websocket/ui"
	"os"

	"golang.org/x/term"
)

func startListenner(_ any) result.Result {
	startTerminateListener()
	startInputListener()
	startResizeListenner()
	return result.Ok(nil)
}

func systemInit(_ any) result.Result {
	fd := int(os.Stdin.Fd())
	var err error
	oldState, err = term.MakeRaw(fd)
	if err != nil {
		return result.Err(fmt.Errorf("could not set terminal to raw mode: %v", err))
	}

	hideCursor()

	window = &Window{
		isRunning:   true,
		renderCount: 0,
	}
	return result.Ok(nil)
}

func mainLoop(_ any) result.Result {
	for window.isRunning {
		select {
		case <-listener:
			window.renderCount++
			ui.Update()
			clearScreen()
			ui.Render()
		default:
			if window.renderCount == 0 {
				window.renderCount++
				ui.Update()
				clearScreen()
				ui.Render()
			}
			continue
		}
	}
	return result.Ok(nil)
}

func cleanup() {
	if err := terminalRestore().Error(); err != nil {
		fmt.Println(err)
	}
	clearScreen()
	showCursor()
	fmt.Print("\033[2J")
}

func Start() {
	defer cleanup()
	fmt.Print("\033[2J")
	if err := systemInit(
		nil,
	).Bind(
		ui.Init,
	).Bind(
		startListenner,
	).Bind(
		mainLoop,
	).Error(); err != nil {
		fmt.Println(err)
	}
}
