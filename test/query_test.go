/*
	Hints:
	go clean -testcache:
		To clean the test cache

	go test -v test/query_test.go -run Get_root(Regexp):
		To test specific case
*/

package test

import (
	"testing"
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
)

const ipAddr = "http://127.0.0.1"
const port = ":9000"
const version = "/v1"

func GetIP(p string) string {
	return ipAddr + port + version + p
}

func ExtractResponse(r *http.Response) {
	buf, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(buf))
}

func Test_Get_root(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(ipAddr + port + "/")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_arc(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/arc"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_arcparams(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/arcparams")+"?group=arc.welder")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_bookprograms(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/bookprograms"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_enum(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/enum"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_extaxis(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/extaxis"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_interference(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/interference"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_metadata(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/metadata")+"?lang=zh_cn")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_params(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/params"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_parameterbyid(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/parameterbyid")+"?md_id=param.speed.speed_min_joint")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_parameterbygroup(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/parameterbygroup")+"?group=param.speed")
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_ref(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/ref"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_toolframe(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/toolframe"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_userframe(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/userframe"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

func Test_Get_all_zeropoints(t *testing.T) {
	start := time.Now()
	resp, err := http.Get(GetIP("/zeropoints"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		d := time.Since(start)
		ExtractResponse(resp)
		fmt.Println("time: ", d)
	}
}

/*func Test_Backup(t *testing.T) {
	resp, err := http.Post(GetIP("/db/backup"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		ExtractResponse(resp)
	}
}*/

func Test_Backup_List(t *testing.T) {
	resp, err := http.Get(GetIP("/db/backup"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		ExtractResponse(resp)
	}
}

/*func Test_Restore(t *testing.T) {
	resp, err := http.Post(GetIP("/db/restore"))
	if err!=nil {
		t.Error("test failed")
	} else {
		t.Log("test pass")
		ExtractResponse(resp)
	}
}*/