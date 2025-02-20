package system

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func terminalRestore() {
	if oldState != nil {
		term.Restore(int(os.Stdin.Fd()), oldState)
	}
}

func terminateListener(signal chan os.Signal) {
	<-signal
	window.isRunning = false
}

func startTerminateListener() {
	termSig := make(chan os.Signal, 1)
	signal.Notify(termSig, syscall.SIGTERM, syscall.SIGINT)
	go terminateListener(termSig)
}

func hideCursor() {
	fmt.Print("\033[?25l")
}

func showCursor() {
	fmt.Print("\033[?25h")
}

var oldState *term.State = nil
