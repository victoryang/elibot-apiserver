package mcserver

import (
    "net"
    "bufio"
    "strings"

    Log "elibot-apiserver/log"
)

func parse(s string) string {
    return strings.Trim(s, CMDLINE)
}

func parse1(s string) string {
    var res string
    ss := strings.Split(s, "\n")
    for _, v := range ss {
        if v != CMDLINE {
            res+=v
        }
    }
    return res
}

/* Important: should call flush() to send buffer to network*/
func writeline(conn net.Conn, cmd string) error {
    writer := bufio.NewWriter(conn)
    _, err := writer.WriteString(cmd)
    if err != nil {
        Log.Error("Error to send message because of ", err.Error())
        return err
    }

    Log.Debug("writeline to conn: ", cmd)
    return writer.Flush()
}

func originalwrite(conn net.Conn, command string) error {
    _, e := conn.Write([]byte(command))
    if e != nil {
        Log.Error("Error to send message because of ", e.Error())
        return e
    }
    return nil
}

func originalread(conn net.Conn) (string, error){
    buf := make([]byte, BUFSIZE)
    n , err := conn.Read(buf)
    if err != nil {
        Log.Error("Error to read message because of ", err)
        return "", err
    }
    return string(buf[:n]), nil
}

func readline(conn net.Conn) (string, error) {
    reader := bufio.NewReader(conn)
    res, err := reader.ReadString('\n')
    if err != nil {
        Log.Error("Error to send message because of ", err.Error())
        return "", err
    }
    Log.Debug("readline from conn: ", res)
    return res, nil
}

func readall(conn net.Conn) (string, error) {
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