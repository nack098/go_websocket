//go:build !windows

package platform

import (
	"fmt"
	"go_websocket/ui"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ResizeListenner() {
	resizeSig := make(chan os.Signal, 1)
	signal.Notify(resizeSig, syscall.SIGWINCH)
	for {
		<-resizeSig
		if err := ui.Resize(nil).Error(); err != nil {
			fmt.Print(err)
		}
		time.Sleep(time.Millisecond * 50)
	}
}
