//go:build windows

package platform

import (
	"golang.org/x/sys/windows"
)

func GetTermSize() (int, int, error) {
	handle, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		return 0, 0, err
	}

	var csbi windows.ConsoleScreenBufferInfo
	err = windows.GetConsoleScreenBufferInfo(handle, &csbi)
	if err != nil {
		return 0, 0, err
	}
	width := int(csbi.Window.Right - csbi.Window.Left + 1)
	height := int(csbi.Window.Bottom - csbi.Window.Top + 1)
	return width, height, nil
}
