//go:build windows

package platform

import (
	"go_websocket/ui"
	"time"
)

func ResizeListenner() {
	for {
		ui.Resize(nil)
		time.Sleep(time.Millisecond * 150)
	}
}
