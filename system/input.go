package system

import (
	"fmt"
	"go_websocket/ui"
	"os"
)

func readRaw() (string, error) {
	buffer := make([]byte, 12)
	n, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("cannot read input to buffer: %v", err)
	}

	return string(buffer[:n]), nil
}

func getInput() {
	for window.isRunning {
		buf, err := readRaw()
		if err != nil {
			fmt.Println(err)
			window.isRunning = false
		}
		// SIGINT Substitution
		if buf == "\x03" {
			window.isRunning = false
		}
		ui.HandleInput(&buf)
	}
}

func startInputListener() {
	go getInput()
}
