package mcserver

import (
	"fmt"
	"net"
    "strings"
)

const (
    CMDLINE = "mcserver>"
)

func parse(s string) string {
    var res string
    ss := strings.Split(s, "\n")
    for _, v := range ss {
        if v != CMDLINE {
            res+=v
        }
    }
    return res
}

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
        return "", err
    }
    return parse(string(buf[:n])), nil
}

func consumeCommandLineIf (conn net.Conn) bool {
	res, err := ReadMessage(conn)
	if err!=nil || res == "" || res == CMDLINE{
        return true
    }

	return false
}