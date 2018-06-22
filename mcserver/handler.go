package mcserver

import (
	"net"
	"time"

    Log "elibot-apiserver/log"
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

func WriteMessage(conn net.Conn, command string) error {
    _, e := conn.Write([]byte(command))
    if e != nil {
        Log.Error("Error to send message because of ", e.Error())
        return e
    }
    return nil
}

func ReadMessage(conn net.Conn) (string, error){
    buf := make([]byte, BUFSIZE)
    n , err := conn.Read(buf)
    if err != nil {
        Log.Error("Error to read message because of ", err)
        return "", err
    }
    return string(buf[:n]), nil
}

func HandleCommand(ctx context.Context, conn net.Conn, command string) (string, error) {
	if _, err := read(conn); err!=nil {
		return "", err
	}

	err := writeline(conn, command)
	if err!=nil {
		return "", err
	}

	return readline(conn)
}