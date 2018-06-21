package mcserver

import (
	"net"

    Log "elibot-apiserver/log"
)

const (
    CMDLINE = "mcserver>"
    BUFSIZE = 128
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
    return ReadMessage(conn)
}

func (h *Handler) HandleCommand(command string) (string, error) {
	if _, err := CleanupBeforeSendCommand(h.conn); err!=nil {
		return "", err
	}

	err := WriteMessage(h.conn, command)
	if err!=nil {
		return "", err
	}

	var res string
	res, err = ReadMessage(h.conn)
	return res, err
}

func NewHandler(c interface{}) *Handler{
	return &Handler{
		conn: 	c.(net.Conn),
	}
}