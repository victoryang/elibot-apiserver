package test

import (
	"testing"
	"net/http"
	"fmt"
)

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/arc")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}