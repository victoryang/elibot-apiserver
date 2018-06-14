package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include <stdlib.h>
// #include "include/nv.h"
import "C"
import (
	"fmt"
	"sync"
	"encoding/json"
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
		return new(NvRam)
	}
}

func getAxisCount() int32 {
	return int32(C.get_axis_count());
}

func getMainFile() string {
	return C.GoString(C.get_main_file())
}

func getCurCoordinate() int32{
	return int32(C.get_cur_coordinate())
}

func getNV() []byte{
	nvram := NvRam{
		ProjectName: 		getMainFile(),
		CurCoordinate:		getCurCoordinate(),
	}
	b, _ := json.Marshal(nvram)
	return b
}

func watchNV() ([]byte, string) {
	value := C.watch_nv()
	return getNV(), fmt.Sprint(uint64(value))
}

func InitNVRam() {
	C.init_nv_ram()
}