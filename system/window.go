package system

import (
	"go_websocket/system/platform"
)

type Window struct {
	isRunning bool
}

func startResizeListenner() {
	go platform.ResizeListenner()
}

var window *Window = nil
