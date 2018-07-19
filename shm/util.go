package shm

import (
	"strings"
)

func reformatString(input string) string {
	return strings.Replace(input, "\\", "", -1)
}