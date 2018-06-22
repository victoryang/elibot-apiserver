package mcserver

import (
	"net"

    Log "elibot-apiserver/log"
)

const (
    CMDLINE = "mcserver>"
    BUFSIZE = 255
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

func CleanupBeforeSendCommand (conn net.Conn) (string, error) {
    return read(conn)
}

func (h *Handler) HandleCommand(command string) (string, error) {
	if _, err := read(h.conn); err!=nil {
		return "", err
	}

	err := writeline(h.conn, command)
	if err!=nil {
		return "", err
	}

	return readline(h.conn), nil
}

func NewHandler(c interface{}) *Handler{
	return &Handler{
		conn: 	c.(net.Conn),
	}
}