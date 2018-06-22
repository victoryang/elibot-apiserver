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
	timeoutDuration = 5*time.Second
)

func freeConn(c net.conn) {
	c.Close()
	c = nil
}

func setConnDeadline(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(timeoutDuration))
}

func HandleCommand(conn net.Conn, command string) (string, error) {
	setConnDeadline(conn)

	err := writeline(conn, command)
	if err!=nil {
		freeConn(conn)
		return "", err
	}

	res, err := readline(conn)
	if err!=nil {
		freeConn(conn)
		return "", err
	}
	return parse(res), nil
}