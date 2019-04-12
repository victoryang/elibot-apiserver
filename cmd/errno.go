package cmd

import (
	"fmt"
)

const (
	ERR_NONE = iota
	ERR_OPEN_LOG_FILE_FAIL
	ERR_AUTH_INIT_FAIL
	ERR_START_MCSERVER
	ERR_START_APISERVER
	ERR_START_GRPCSERVER
	ERR_START_WSSERVER
	ERR_START_RESERVER
	ERR_START_PARAMSERVER
)

func returnError(errno int) error {
	return fmt.Errorf("Fail to start server due to: -" + string(errno))
}
