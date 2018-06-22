package mcserver

import (
	"net"
	"time"
	"context"

    //Log "elibot-apiserver/log"
)

const (
    CMDLINE = "mcserver>"
    BUFSIZE = 255
)

const (
	timeoutDuration = 1*time.Second
)

func HandleCommand(conn net.Conn, command string) (string, error) {
	err := writeline(conn, command)
	if err!=nil {
		return "", err
	}

	res, err := readline(conn)
	if err!=nil {
		return "", err
	}
	return parse(res), nil
}