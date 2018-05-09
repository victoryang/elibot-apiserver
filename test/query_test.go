package test

import (
	"testing"
	"net/http"
	"fmt"
)

func Test_Get_root(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arc(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/arc")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_arcparams(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/arcparams")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_bookprograms(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/bookprograms")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_extaxis(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/extaxis")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_interference(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/interference")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_metadata(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/metadata")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_parameterbyid(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/parameterbyid")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_parameterbygroup(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/parameterbygroup")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_ref(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/ref")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_toolframe(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/toolframe")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_userframe(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/userframe")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}

func Test_Get_all_zeropoints(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:9000/v1/zeropoints")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		fmt.Println(resp)
	}
}