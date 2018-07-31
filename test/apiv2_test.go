package test

import (
	"testing"
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"bytes"
	"encoding/json"
)

const (
	serverip = "http://192.168.1.217:9000"
)

func ExtractResponseToString(r *http.Response) string {
	buf, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return string(buf)
}

func Test_SetParam(t *testing.T) {
	start := time.Now()
	data := []byte(`{data:40}`)
	req,_ := http.NewRequest(
		"PUT", 
		"192.168.1.217:9000/v2/robot/repository/param?md_id=param.body.robot_type&&index=0", 
		bytes.NewBuffer(data),
	)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err!=nil {
		t.Error("test failed")
	} else {
		defer resp.Close()
		d := time.Since(start)
		fmt.Println(ExtractResponseToJson(resp))
		fmt.Println("time: ", d)
	}
}

