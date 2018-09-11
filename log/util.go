package log

import (
	"os"
)

const (
	MaxLogFileSize = 2 << 20
)

func checkFileSizeExceededMax(filename string) bool {
	if info, err := os.Stat(filename); err== nil {
		if info.Size() > MaxLogFileSize {
			return true
		}
	}

	return false
}