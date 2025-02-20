//go:build windows

package platform

import (
	"fmt"
	"syscall"
)

func Interrupt() error {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	generateConsoleCtrlEvent := kernel32.NewProc("GenerateConsoleCtrlEvent")
	result, _, err := generateConsoleCtrlEvent.Call(syscall.CTRL_BREAK_EVENT, 0)
	if result == 0 {
		return fmt.Errorf("unable to interrupt process: %v", err)
	}
	return nil
}
