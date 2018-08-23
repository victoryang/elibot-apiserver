package mcserver

import (
	"net"
	"time"

    //Log "elibot-apiserver/log"
)

const (
    CMDLINE = "mcserver>"
    BUFSIZE = 255
)

const (
	timeoutDuration = 5*time.Second
)

func freeConn(c net.Conn) {
	c.Close()
	c = nil
}

func setConnDeadline(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(timeoutDuration))
}

func checkConnTimeout(err error) bool {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return true
	}
	return false
}

func HandleCommand(conn net.Conn, command string) (string, error) {
	setConnDeadline(conn)

	err := writeConn(conn, command)
	if err!=nil {
		if ok := checkConnTimeout(err); !ok {
			freeConn(conn)
		}
		return "", err
	}

	res, err := readConn(conn)
	if err!=nil {
		if ok := checkConnTimeout(err); !ok {
			freeConn(conn)
		}
		return "", err
	}
	return parse(res), nil
}