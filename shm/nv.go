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

	Log "elibot-apiserver/log"
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
		return bytes.NewBuffer(make([]byte, 0, BufferSize))
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

func getSysTechnics() int32 {
	return int32(C.get_sys_technics())
}

func getCurCoordinate() int32{
	return int32(C.get_cur_coordinate())
}

func getMannualSpeedRate() int32{
	return int32(C.get_mannual_speed_rate())
}

func getCycleMode() int32{
	return int32(C.get_cycle_mode())
}

func getToolNumber() int32{
	return int32(C.get_tool_number())
}

func getUserNumber() int32{
	return int32(C.get_user_number())
}

func getZeroEncode() []int32 {
	num := int(getAxisCount())
	r := make([]int32, 0)
	for i:=0; i<num; i++ {
		r = append(r, int32(C.get_zero_encode(C.int(i))))
	}
	return r
}

func getSystemPeriod() float64{
	return float64(C.get_system_period())
}

func getSysCtrlMode() int32{
	return int32(C.get_sys_ctrl_mode())
}

func getOrigin() []float64 {
	num := int(getAxisCount())
	r := make([]float64, 0)
	for i:=0; i<num; i++ {
		r = append(r, float64(C.get_origin(C.int(i))))
	}
	return r
}

func getNVAndCompare() (res []byte){
	mutex.Lock()
	defer mutex.Unlock()

	nvram := NvRam{
		ProjectName: 		getMainFile(),
		Sys_technics:		getSysTechnics(),
		CurCoordinate:		getCurCoordinate(),
		ManualSpeedRate:	getMannualSpeedRate(),
		CycleMode:			getCycleMode(),
		ToolNum:			getToolNumber(),
		UserNum:			getUserNumber(),
		Zero_encode:		getZeroEncode(),
		System_period:		getSystemPeriod(),
		System_ctrl_mode:	getSysCtrlMode(),
		Origin:				getOrigin(),
	}

	var crc int
	now, err := json.Marshal(nvram)
	if err!=nil {
		crc = 0
	} else {
		crc = int(crc32.ChecksumIEEE(now))
	}

	var cache []byte
	buf := NVRamPool.Get().(*bytes.Buffer)
	if buf == nil {
		buf = bytes.NewBuffer(make([]byte, 0, BufferSize))
		crc_nv = 0
	} else {
		cache = buf.Bytes()
	}
 	
 	if crc == 0 {
 		crc_nv = 0 
 		res = []byte("")
 	} else {
 		if crc == crc_nv {
 			res = nil
 			return
 		} else {
 			crc_nv = crc
 			res = now
 		}
 	}
 
 	buf.Reset()
 	buf.Write(res)
 	defer func(c []byte) {
		if e := recover(); e!=nil {
			Log.Error("buf write error: ", e)

			buf := make([]byte, 0, BufferSize*2)
			buf = c[:]
			NVRamPool = sync.Pool{
				New: func() interface{} {
					return bytes.NewBuffer(buf)
				},
			}
		}
	}(cache)
	
	NVRamPool.Put(buf)
	return
}

func watchNV(modified chan []byte) {
	if res := getNVAndCompare(); res != nil {
		modified <- res
	}
}

func InitNVRam() int {
	return int(C.init_nv_ram())
}