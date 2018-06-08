package sharedmemory

import "C"
import (
	"encoding/binary"
	Log "elibot-apiserver/log"
)

func ConvertToByteArray(value interface{}) []byte{
	var buf []byte
	switch t := value.(type) {
	case uint32:
		binary.LittleEndian.PutUint32(buf, value.(uint32))
		return buf
	default:
		Log.Error("Type not support")
		return nil
	}
}