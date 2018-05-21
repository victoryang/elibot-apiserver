package mcserver

import (
	"errors"
	"fmt"
	"net"
)

var test = "testGo 0 1"
var end = "\n"

type MCserver struct {
	Addr		string
	Port		string
	Conn 		net.Conn
}

func handleCommand(conn net.Conn, command string) (string, error) {
	err := WriteMessage(conn, command)
	if err!=nil {
		return "", err
	}

	return ReadMessage(conn)
}

func (mc *MCserver) OnCommandRecived() (string, error) {
	cmd := test + end
	if testCommandLine(mc.Conn) {
		return handleCommand(mc.Conn, cmd)
	}
	return "", errors.New("Not in a proper situation")
}

func (mc *MCserver) Connect() error {
	address := mc.Addr + mc.Port
    mc.Conn, err := net.Dial("unix", address)
    if err != nil {
        fmt.Println("dial error:", err)
        return err
    }
    defer mc.Conn.Close()

    fmt.Println("mcserver connected")
    return nil
}

func NewServer() *MCserver {
	mcserver := &MCserver {
		Addr:	"192.168.1.106",
		Port:	":8055",
	}
	return mcserver
}