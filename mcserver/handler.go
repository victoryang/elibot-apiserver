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

type Handler struct {
	conn 		net.Conn
}

func (h *Handler)HandleCommand(ctx context.Context, command string) (string, error) {
	err := writeline(h.conn, command)
	if err!=nil {
		return "", err
	}

	res, err := readline(h.conn)
	if err!=nil {
		return "", err
	}
	return parse(res), nil
}

func NewHandler(c net.Conn) *Handler {
	return &Handler {
		conn: 	c,
	}
}