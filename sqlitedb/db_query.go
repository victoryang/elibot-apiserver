package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db_query.h>
// #include<cJSON.h>
import "C"
import (
    "unsafe"
    "bytes"
    "fmt"
    "math"
    "errors"
    "encoding/binary"
)

func newsqlparams(key string, value interface{}, param *C.sql_parameter) error {
    param.name = C.CString(key)
   
    switch t := value.(type) {
    case int32:
        param._type = C.DB_TYPE_INT32
        buf := new(bytes.Buffer)
        binary.Write(buf, binary.LittleEndian, value.(int32))
        copy(param.value[:], buf.Bytes())
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
        fmt.Println("not support for ", t)
        return errors.New("not support this type\n")
    }

    return nil
}

func Db_query_with_params(q_id, db_name string, queries map[string]interface{}) (string, error) {
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString(db_name)
    defer C.free(unsafe.Pointer(conn))

    req_params := C.new_db_query_req_parameter()
    defer C.free(unsafe.Pointer(req_params))

    option := C.new_db_query_req_option(C.DB_QUERY_MODE_CUSTOM)
    defer C.free(unsafe.Pointer(option))

    req_params = C.new_db_query_req_parameter()
    defer C.free(unsafe.Pointer(req_params))
    req_params.param_size = C.int16_t(len(queries))

    req_params.params = C.new_sql_parameter(req_params.param_size)
    defer C.free(unsafe.Pointer(req_params.params))
 
    var i C.int16_t = 0
    for k,v := range queries {
        err := newsqlparams(k, v, C.getindexedsqlparam(req_params, i++))
        if i > req_params.param_size || err != nil{
            return "", errors.New("fail to start a query\n")
        }
    }

    req := &C.db_query_req{
                    id,
                    conn,
                    option,
                    req_params,
                    nil,
                }

    q_res := C.db_query(req)

    for i=0;i<req_params.param_size;i++ {
        param := C.getindexedsqlparam(req_params, i)
        if param!=nil {
            C.free(unsafe.Pointer(param.name))
            if param._type == C.DATA_STRING {
                p := binary.LittleEndian.Uint64(param.value[:])
                C.free(unsafe.Pointer(uintptr(p)))
            }
        }
    }

    if q_res==nil {
        fmt.Println("fail to query")
        return "", errors.New("fail to query\n")
    }

    str := C.cJSON_Print(q_res)
    defer C.cJSON_Delete(q_res)

    res := C.GoString(str)
    defer C.free(unsafe.Pointer(str))

    return res, nil
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
    	fmt.Println("fail to query")
    	return "", errors.New("fail to query\n")
    }

    str := C.cJSON_Print(q_res)
    defer C.cJSON_Delete(q_res)

    res := C.GoString(str)
    defer C.free(unsafe.Pointer(str))

    return res, nil
}