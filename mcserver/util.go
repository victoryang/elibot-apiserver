package mcserver

import (
    "net"
    "bufio"
    Log "elibot-apiserver/log"
)

func writeline(conn net.Conn, cmd string) error {
    writer := bufio.NewWriter(conn)
    _, err := writer.WriteString(cmd)
    if err != nil {
        Log.Error("Error to send message because of ", err.Error())
        return err
    }
    return nil
}

func readline(conn net.Conn) string {
    scanner := bufio.NewScanner(conn)
    if scanner.Scan() {
        Log.Debug("read from conn: ", scanner.Text())
        return scanner.Text()
    }
    return ""
}

func read(conn net.Conn) (string, error) {
    buf := make([]byte, BUFSIZE)
    reader := bufio.NewReader(conn)
    n, err := reader.Read(buf)
    if err != nil {
        Log.Error("Error to read message because of ", err)
        return "", err
    }
    Log.Debug("read from conn: ", string(buf[:n]))
    return string(buf[:n]), nil
}

/* A safe send*/
func SafeSendResponseToChannel(ch chan Response, resp Response) {
    defer func() {
        if err := recover(); err!=nil {
            Log.Error("Client closed response channel... drop it")
        }
    }()
    ch <- resp
}

func Deprecated_SafeSendResponseToChannel(ch chan Response, resp Response) {
    if _, ok := <- ch; ok {
        ch <- resp
    } else {
        Log.Error("Client closed response channel... drop it")
    }
}