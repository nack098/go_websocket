package system

import (
	"go_websocket/system/platform"
)

type Window struct {
	isRunning   bool
	renderCount int
}

func startResizeListenner() {
	go platform.ResizeListenner()
}

var window *Window = nil
