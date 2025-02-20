//go:build !windows

package platform

import (
	"os"

	"golang.org/x/term"
)

func GetTermSize() (int, int, error) {
	return term.GetSize(int(os.Stdin.Fd()))
}
