//go:build windows

package platform

import (
	"fmt"
	"syscall"
)

func Interrupt() error {
	dll, err := syscall.LoadDLL("kernel32.dll")
	if err != nil {
		return fmt.Errorf("unable to load dll: %v", err)
	}
	procedure, err := dll.FindProc("GenerateConsoleCtrlEvent")
	if err != nil {
		return fmt.Errorf("unable to find procedure: %v", err)
	}
	result, _, err := procedure.Call(syscall.CTRL_BREAK_EVENT, uintptr(syscall.Getpid()))
	if result == 0 {
		return fmt.Errorf("unable to interrupt process: %v", err)
	}
	return nil
}
