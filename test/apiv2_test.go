package test

import (
	"testing"
	"net/http"
	"net/url"
	"fmt"
	"time"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

const (
	serverip = "http://192.168.1.253:9000"
)

func ExtractResponseToString(r *http.Response) string {
	buf, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		fmt.Println(err)
		return ""
	}
	return string(buf)
}

func SendToServer(method string, url string, data []byte) {
	start := time.Now()
	req,_ := http.NewRequest(method, url, bytes.NewBuffer(data))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err!=nil {
		fmt.Error("test failed")
	} else {
		defer resp.Close()
		d := time.Since(start)
		fmt.Println(ExtractResponseToString(resp))
		fmt.Println("time: ", d)
	}
	return
}

func Test_SetArcParam(t *testing.T) {
	u,_ := url.Parse(serverip)
	file_no := "0"
	md_id := "arc.welder.prepareAspirationTime"
	u.Path = "v1/robot/repository/arcparam/" + file_no + "/" + md_id

	SendToServer("PUT", u.String(), []byte(`{"index":"0", "value":"40"}`))
}

func Test_SetParam(t *testing.T) {
	SendToServer("PUT", "192.168.1.217:9000/v1/robot/repository/params/param.body.robot_type", []byte(`{"index":"0", "value":"40"}`))
}

func Test_SetParam(t *testing.T) {
	SendToServer("PUT", "192.168.1.217:9000/v1/robot/repository/params/param.body.robot_type", []byte(`{"index":"0", "value":"40"}`))
}

func Test_SetParam(t *testing.T) {
	SendToServer("PUT", "192.168.1.217:9000/v1/robot/repository/params/param.body.robot_type", []byte(`{"index":"0", "value":"40"}`))
}

