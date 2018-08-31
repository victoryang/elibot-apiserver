package accesslog

import (
	"os"
)

const (
	MaxLogFileSize = 2 << 20
)

func checkFileSizeExceededMax (file *os.File) bool {
	if info, err := file.Stat(); err==nil {
		if info.Size() > MaxLogFileSize {
			return true
		}
	}

	return false
}