package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include <stdlib.h>
// #include <include/nv.h>
import "C"
import (
	"bytes"
	"fmt"
	"sync"
	"encoding/json"
	"hash/crc32"
)

type NvRam struct {
	ProjectName			string 		`json:"projectname,omitempty"`
	Sys_technics		int32		`json:"sys_technics,omitempty"`	
	CurCoordinate		int32		`json:"curcoordinate,omitempty"`
	ManualSpeedRate		int32		`json:"manualspeedRate,omitempty"`
	CycleMode			int32		`json:"cyclemode,omitempty"`
	ToolNum				int32		`json:"toolnum,omitempty"`
	UserNum				int32		`json:"usernum,omitempty"`
	Zero_encode			[]int32		`json:"zero_encode,omitempty"`
	System_period		float64		`json:"system_period,omitempty"`
	System_ctrl_mode	int32		`json:"system_ctrl_mode,omitempty"`
	Origin				[]float64	`json:"origin,omitempty"`
}

var NVRamPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}
var crc_old int = 0

func getAxisCount() int32 {
	return int32(C.get_axis_count());
}

func getMainFile() string {
	return C.GoString(C.get_main_file())
}

func getCurCoordinate() int32{
	return int32(C.get_cur_coordinate())
}

func getZeroEncode() []int32 {
	num := getAxisCount()
	r := make([]int32, 0)
	for i=0; i<num; i++ {
		r = append(r, C.get_zero_encode(i))
	}
	return r
}

func getAndCompare() []byte{
	nvram := NvRam{
		ProjectName: 		getMainFile(),
		CurCoordinate:		getCurCoordinate(),
	}

	buf := NVRamPool.Get().(*bytes.Buffer)
	old := buf.Bytes()
	now, err := json.Marshal(nvram)
	if err!=nil {
		if old!=nil {
			return old
		} else {
			return nil
		}
	}
	crc_now := int(crc32.ChecksumIEEE(now))
	if crc_now != crc_old {
		buf.Reset()
		NVRamPool.Put(buf.Write(now))
		crc_old = crc_now
		return now
	}
	return nil
}

func watchNV(modified chan []byte) {
	modified <- getAndCompare()
}

func InitNVRam() {
	C.init_nv_ram()
}