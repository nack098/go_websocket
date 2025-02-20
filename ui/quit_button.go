package ui

import (
	"fmt"
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

	err := proc.Signal(syscall.SIGTERM)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (b *quitButton) getName() *string {
	return &b.name
}
