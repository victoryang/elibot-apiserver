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

var mcserver *MCserver = nil

func handleCommand(conn net.Conn, command string) (string, error) {
	err := WriteMessage(conn, command)
	if err!=nil {
		return "", err
	}

	return ReadMessage(conn)
}

func OnCommandRecived() (string, error) {
	cmd := test + end
	if testCommandLine(mcserver.Conn) {
		return handleCommand(mcserver.Conn, cmd)
	}
	return "", errors.New("Not in a proper situation\n")
}

func (mc *MCserver) Connect() error {
	if mc == nil {
		return errors.New("Server does not exist")
	}
	var err error
	address := mc.Addr + mc.Port
    mc.Conn, err = net.Dial("tcp", address)
    if err != nil {
        fmt.Println("dial error:", err)
        return err
    }

    fmt.Println("mcserver connected")
    return nil
}

func (mc *MCserver) Close() {
	mc.Conn.Close()
} 

func NewMCServer() *MCserver {
	mcserver := &MCserver {
		Addr:	"192.168.1.106",
		Port:	":8055",
	}
	return mcserver
}