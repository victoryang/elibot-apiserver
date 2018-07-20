package sqlitedb

// #cgo CFLAGS: -I/root/mcserver/include
// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<mcsql.h>
import "C"
import (
    "unsafe"
    "bytes"
    "math"
    "errors"
    "encoding/binary"

    Log "elibot-apiserver/log"
)

func setsSqlparams(key string, value interface{}, param *C.sql_parameter) error {
    param.name = C.CString(key)

    switch t := value.(type) {
    case int32:
        C.add_int32_to_param(value.(int32), param)
    case uint32:
        param._type = C.DB_TYPE_INT32
        binary.LittleEndian.PutUint32(param.value[:], value.(uint32))
    case int64:
        param._type = C.DB_TYPE_INT64
        buf := new(bytes.Buffer)
        binary.Write(buf, binary.LittleEndian, value.(int64))
        copy(param.value[:], buf.Bytes())
    case uint64:
        param._type = C.DB_TYPE_INT64
        binary.LittleEndian.PutUint64(param.value[:], value.(uint64))
    case float64:
        param._type = C.DB_TYPE_DOUBLE
        binary.LittleEndian.PutUint64(param.value[:], math.Float64bits(value.(float64)))
    case string:
        param._type = C.DB_TYPE_TEXT
        p := C.CString(value.(string))
        binary.LittleEndian.PutUint64(param.value[:], uint64(uintptr(unsafe.Pointer(p))))

    default:
        Log.Print("not support for ", t)
        return errors.New("not support this type\n")
    }

    return nil
}

type Parameter struct {
    parameter       *C.db_query_req_parameter
    param_size      int
}

func (p *Parameter) newReqParameter(vars map[string]interface{}) *C.db_query_req_parameter {
    p.param_size = len(vars)
    p.parameter = C.new_db_query_req_parameter(C.int16_t(p.param_size))
    if p.parameter == nil {
        return nil
    }

    var i C.int16_t = 0
    for k,v := range vars {
        err := setsSqlparams(k, v, C.get_sqlparam_index(p.parameter, i))
        i++
        if i > C.int16_t(p.param_size) || err != nil{
            return nil
        }
    }
    return p.parameter
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

func Db_query_with_params(q_id, db_name string, vars map[string]interface{}) ([]byte, error) {
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString(db_name)
    defer C.free(unsafe.Pointer(conn))

    p := new(Parameter)
    req_parameter := p.newReqParameter(vars)
    defer p.freeReqParameter()

    buf := C.mcsql_query_with_param(id, conn, C.DB_QUERY_MODE_CUSTOM, req_parameter, nil)
    if buf == nil {
        return nil, errors.New("failed to query")
    }
    res := C.GoString(buf)
    defer C.free(unsafe.Pointer(buf))

    return []byte(res), nil
}

func Db_query(q_id string, db_name string) (string, error){
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString(db_name)
    defer C.free(unsafe.Pointer(conn))

    option := C.new_db_query_req_option(C.DB_QUERY_MODE_STANDARD)
    defer C.free(unsafe.Pointer(option))

    req := &C.db_query_req{
					id,
					conn,
					option,
					nil,
					nil,
				}

    q_res := C.db_query(req)
    if q_res==nil {
    	Log.Error("fail to query")
    	return "", errors.New("fail to query\n")
    }

    str := C.cJSON_Print(q_res)
    defer C.cJSON_Delete(q_res)

    res := C.GoString(str)
    defer C.free(unsafe.Pointer(str))

    return res, nil
}