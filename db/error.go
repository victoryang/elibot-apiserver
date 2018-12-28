package db

import (
	"github.com/mattn/go-sqlite3"
)

const (
	ErrModuleTag = 100000

	ErrNone = 0
)

func ErrCode(e error) int {
	err := e.(sqlite3.Error)

	return int(err.ExtendedCode) + ErrModuleTag
}