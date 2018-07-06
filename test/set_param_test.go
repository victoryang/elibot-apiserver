package main 

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"fmt"
)

func main() {
	data := []byte(`{data:40}`)
	req,_ := http.NewRequest("PUT", "192.168.1.217:9000/v2/robot/repository/param?md_id=param.body.robot_type&&index=0", bytes.NewBuffer(data))
	client := http.Client{}
	rsp, err := client.Do(req)
	if err!=nil {
		fmt.Println(err)
	} else {
		defer rsp.Body.Close()
		buf, err := ioutil.ReadAll(rsp.Body)
		if err!=nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(buf))
		}
	}

	return
}