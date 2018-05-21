package mcserver

import (
	"fmt"
	"net"
)

func WriteMessage(conn net.Conn, command string) error {
    _, e := conn.Write([]byte(command))
    if e != nil {
        fmt.Println("Error to send message because of ", e.Error())
        return e
    }
    return nil
}

func ReadMessage(conn net.Conn) (string, error){
    buf := make([]byte, 1024)
    n , err := conn.Read(buf)
    if err != nil {
        fmt.Println("Error to read message because of ", err)
        return "", nil
    }
    return string(buf[:n]), nil
}

func testCommandLine(conn *net.Conn) bool {
	res, err := ReadMessage(conn)
	if err!=nil || res!="mcserver>" {
		fmt.Println("Read error: ", err)
		return false
	}

	return true
}