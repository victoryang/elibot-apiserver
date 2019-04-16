package resource

// #cgo CFLAGS: -I../build/include/include
// #cgo LDFLAGS: -L../build/lib -lrobresource
// #include <resource.h>
import "C"
import (
    "sync"
    "hash/crc32"
    "unsafe"
)

func decode(cstr *C.char) []byte {
    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    return []byte(gostr)
}

func GetNVOnce() []byte {
    return decode(C.elt_get_nv())
}

func GetSysVar(dtype int, start int, end int) []byte {
    return decode(C.elt_get_sysvar(C.int(dtype), C.int(start), C.int(end)))
}

func GetLocVar(dtype int, num int, start int, end int) []byte {
    return decode(C.elt_get_locvar(C.int(dtype), C.int(num), C.int(start), C.int(end)))
}

func GetPLCOnce() []byte {
    return decode(C.elt_get_plc())
}

func GetResourceOnce() []byte {
    return decode(C.elt_get_resource())
}

/*v2 api*/
func GetIoInput(addr int) int {
    return int(C.elt_get_io_in(C.int(addr)))
}

func GetIoOutput(addr int) int {
    return int(C.elt_get_io_out(C.int(addr)))
}

func GetIoVirtualInput(addr int) int {
    return int(C.elt_get_io_vin(C.int(addr)))
}

func GetIoVirtualOutput(addr int) int {
    return int(C.elt_get_io_vout(C.int(addr)))
}

func GetRobotState() int {
    return int(C.elt_get_robot_state())
}

func GetRobotMode() int {
    return int(C.elt_get_robot_mode())
}

func GetCurRobotPos() []byte {
    return decode(C.elt_get_cur_robot_pos())
}

func GetCurRobotPose() []byte {
    return decode(C.elt_get_cur_robot_pose())
}

func GetCurProgramLine() int {
    return int(C.elt_get_cur_program_line())
}

func GetServoStatus() int {
    return int(C.elt_get_servo_enabled())
}

func GetMotorStatus() int {
    return int(C.elt_can_motor_run())
}

func GetPlaySpeed() int {
    return int(C.elt_get_speed_modify_play())
}

func GetMainProgramName() []byte {
    return decode(C.elt_get_main_program_name())
}

func GetManualSpeedRate() int {
    return int(C.elt_get_manual_speed_rate())
}

func GetToolNum() int {
    return int(C.elt_get_tool_num())
}

func GetUserNum() int {
    return int(C.elt_get_user_num())
}

func GetEncryptionStatus() int {
    return int(C.elt_get_encryption_status())
}

func GetEncryptionRemainTime() int {
    return int(C.elt_get_encryption_remain_time())
}

func GetRemoteModeFromResource() int {
    return int(C.elt_get_remote_mode_status())
}

/*--------------------------------------------------*/

var crc_nv int = 0
var crc_resource int = 0
var crc_plc int = 0

var nv_mutex sync.Mutex
var resource_mutex sync.Mutex
var plc_mutex sync.Mutex

func getNVAndCompare() []byte {
    nv_mutex.Lock()
    defer nv_mutex.Unlock()

    cstr := C.elt_get_nv()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    buf := []byte(gostr)

    crc := int(crc32.ChecksumIEEE(buf))
    if crc == crc_nv {
        return nil
    }
    crc_nv = crc
    return buf
}

func watchNV(modified chan []byte) {
    if res := getNVAndCompare(); res != nil {
        modified <- res
    }
}

func getPLCAndCompare() []byte {
    plc_mutex.Lock()
    defer plc_mutex.Unlock()

    cstr := C.elt_get_plc()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    buf := []byte(gostr)

    crc := int(crc32.ChecksumIEEE(buf))
    if crc == crc_plc {
        return nil
    }
    crc_plc = crc
    return buf
}

func watchPLC(modified chan []byte) {
    if res := getPLCAndCompare(); res != nil {
        modified <- res
    }
}

func getResourceAndCompare() []byte {
    resource_mutex.Lock()
    defer resource_mutex.Unlock()

    cstr := C.elt_get_resource()

    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    buf := []byte(gostr)

    crc := int(crc32.ChecksumIEEE(buf))
    if crc == crc_resource {
        return nil
    }
    crc_resource = crc
    return buf
}

func watchResource(modified chan []byte) {
    if res := getResourceAndCompare(); res != nil {
        modified <- res
    }
}

func InitResource() int {
    return int(C.elt_resource_init())
}