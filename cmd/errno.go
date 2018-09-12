package cmd

import (
	"fmt"
)

const (
	ERR_NONE = iota
	ERR_OPEN_LOG_FILE_FAIL
	ERR_START_MCSERVER
	ERR_START_APISERVER
	ERR_START_GRPCSERVER
	ERR_START_WSSERVER
	ERR_START_SHMSERVER
)

func returnError(errno int) {
	return fmt.Errorf("Fail to start server due to: -" + errno)
}