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
		return bytes.NewBuffer(make([]byte, 0, bufferSize))
	},
}
var crc_nv int = 0
var mutex sync.Mutex

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
		r = append(r, int32(C.get_zero_encode(C.int(i))))
	}
	return r
}

func getNVAndCompare() []byte{
	mutex.Lock()
	defer mutex.Unlock()

	nvram := NvRam{
		ProjectName: 		getMainFile(),
		CurCoordinate:		getCurCoordinate(),
		Zero_encode:		getZeroEncode(),
	}

	var crc int
	now, err := json.Marshal(nvram)
	if err!=nil {
		crc = 0
	}
	crc = int(crc32.ChecksumIEEE(now))

	buf := NVRamPool.Get().(*bytes.Buffer)
	if buf == nil {
		if crc == 0 {
			return []byte("")
		} else {
			crc_nv = 0
		}
	}

	cache := buf.Bytes()
	if crc == crc_nv {
		NVRamPool.Put(buf)
		return nil
	}

	buf.Reset()
	_, err = buf.Write(now)
	if err != nil {
		crc_nv = 0
		NVRamPool = sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0, bufferSize*2))
			},
		}
		return []byte("")
	}

	crc_nv = crc
	NVRamPool.Put(buf)
	return now
}

func watchNV(modified chan []byte) {
	if res := getNVAndCompare(); res!=nil {
		modified <- res
	}
}

func InitNVRam() {
	C.init_nv_ram()
}