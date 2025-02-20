//go:build !windows

package platform

import "syscall"

func Interrupt() error {
	return syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}
