package ui

import (
	"os"
	"syscall"
)

type quitButton struct {
	name string
}

func (b *quitButton) render() {
}

func (b *quitButton) action() {
	// TODO: Handle Error
	proc, _ := os.FindProcess(os.Getpid())

	proc.Signal(syscall.SIGTERM)
}

func (b *quitButton) getName() *string {
	return &b.name
}
