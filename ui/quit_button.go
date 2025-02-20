package ui

import (
	"fmt"
	"go_websocket/ui/platform"
)

type quitButton struct {
	name string
}

func (b *quitButton) render() {
}

func (b *quitButton) action() {
	err := platform.Interrupt()
	if err != nil {
		fmt.Println(err)
	}
}

func (b *quitButton) getName() *string {
	return &b.name
}
