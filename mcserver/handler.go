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
	if _, err := readline(h.conn); err!=nil {
		return "", err
	}

	err := writeline(h.conn, command)
	if err!=nil {
		return "", err
	}

	return readline(h.conn)
}

func NewHandler(c net.Conn) *Handler {
	return &Handler {
		conn: 	c,
	}
}