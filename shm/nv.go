package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include <stdlib.h>
// #include <nv.h>
import "C"
import (
	"bytes"
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
var crc_nv int = 0

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
	num := int(getAxisCount())
	r := make([]int32, 0)
	for i:=0; i<num; i++ {
		r = append(r, int32(C.get_zero_encode(i)))
	}
	return r
}

func getNVAndCompare() []byte{
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
	if crc_now != crc_nv {
		buf.Reset()
		_, err := buf.Write(now)
		if err != nil {
			crc_nv = 0
			return nil
		}

		NVRamPool.Put(buf.Bytes())
		crc_nv = crc_now
		return now
	}
	return nil
}

func watchNV(modified chan []byte) {
	modified <- getNVAndCompare()
}

func InitNVRam() {
	C.init_nv_ram()
}