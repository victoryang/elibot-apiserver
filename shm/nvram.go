package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include<stdlib.h>
// #include<sharedmemory.h>
// #include<nv.h>
import "C"
import (
	"fmt"
)

type NvRam struct {
	ProjectName			string
	Sys_technics		int32
	CurCoordinate		int32
	ManualSpeedRate		int32
	CycleMode			int32
	ToolNum				int32
	UserNum				int32
	Zero_encode			[]int32
	System_period		float64
	System_ctrl_mode	int32
	Origin				[]float64
}

func getAxisCount() int32 {
	return int32(C.get_axis_count());
}

func checkNV() []byte {
	return []byte("checkNV")
}

func getNV() *NvRam{
	projectname := C.CString(C.get_main_file())
	defer C.free(unsafe.Pointer(projectname))

	return &NvRam{
		ProjectName: 		projectname,
		CurCoordinate:		int32(C.get_cur_coordinate()),
	}
}

func watchNV() string {
	value := C.watch_nv()
	return fmt.Sprint(uint64(value))
}