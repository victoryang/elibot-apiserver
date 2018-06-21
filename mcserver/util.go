package mcserver

import (
	"net"
    "strings"

    Log "elibot-apiserver/log"
)

const (
    CMDLINE = "mcserver>"
    BUFSIZE = 128
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
    return parse(string(buf[:n])), nil
}

func consumeCommandLineIf (conn net.Conn) bool {
	res, err := ReadMessage(conn)
	if err!=nil || res == "" || res == CMDLINE{
        return true
    }

	return false
}

func SafeSendResponseToChannel(ch chan Response, resp Response) {
    if _, ok := <- ch; ok {
        ch <- resp
    } else {
        Log.Error("Client closed response channel... drop it")
    }
}