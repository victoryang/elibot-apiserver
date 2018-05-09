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

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/arcparams")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/bookprograms")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/extaxis")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/interference")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/metadata")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/parameter")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/ref")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/toolframe")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/userframe")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/zeropoints")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}