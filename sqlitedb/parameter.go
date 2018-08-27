package sqlitedb

// #include<stdlib.h>
// #include<mcsql.h>
import "C"
import (
    "unsafe"
    "errors"
    "encoding/binary"

    Log "elibot-apiserver/log"
)

type Parameter struct {
    parameter       *C.db_query_req_parameter
    param_size      int
}

func setsSqlparams(key string, value interface{}, param *C.sql_parameter) error {
    param.name = C.CString(key)

    switch t := value.(type) {
    case int32:
        C.add_int32_to_param(C.int(value.(int32)), param)
    
    case int64:
        C.add_int64_to_param(C.longlong(value.(int64)), param)

    case float64:
        C.add_double_to_param(C.double(value.(float64)), param)

    case string:
        param._type = C.DB_TYPE_TEXT
        C.add_string_to_param(C.CString(value.(string)), param)

    default:
        Log.Print("not support for ", t)
        return errors.New("not support this type\n")
    }

    return nil
}

func (p *Parameter) newReqParameter(vars map[string]interface{}) {
    p.param_size = len(vars)
    p.parameter = C.new_db_query_req_parameter(C.int16_t(p.param_size))
    if p.parameter == nil {
        return
    }

    var i C.int16_t = 0
    for k,v := range vars {
        setsSqlparams(k, v, C.get_sqlparam_index(p.parameter, i))
        i++
    }
    return
}

func (p *Parameter) freeReqParameter() {
    if p.parameter == nil {
        return
    }

    for i:=0; i<p.param_size; i++ {
        param := C.get_sqlparam_index(p.parameter, C.int16_t(i))
        if param!=nil {
            C.free(unsafe.Pointer(param.name))
            if param._type == C.DATA_STRING {
                p := binary.LittleEndian.Uint64(param.value[:])
                C.free(unsafe.Pointer(uintptr(p)))
            }
        }
    }

    C.free_db_query_req_parameter(p.parameter)
}