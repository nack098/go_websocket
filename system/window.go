package system

import (
	"fmt"
	"go_websocket/system/platform"
)

type Window struct {
	isRunning bool
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func startResizeListenner() {
	go platform.ResizeListenner()
}

var window *Window = nil
