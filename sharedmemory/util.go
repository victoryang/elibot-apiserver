package sharedmemory

import "C"
import (
	"reflect"
	"encoding/binary"
	Log "elibot-apiserver/log"
)

func ConvertType(value interface{}) []byte{
	var buf []byte
	t := reflect.TypeOf(C.size_t(value))
	switch t {
	case _Ctype_uint:
		binary.LittleEndian.PutUint32(buf, uint32(value))
		return buf
	default:
		Log.Error("Type not support")
		return nil
	}
}