package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db_query.h>
// #include<cJSON.h>
import "C"
import (
    "unsafe"
    "fmt"
    "math"
    "errors"
    "encoding/binary"
)

const (
    SQL_PARAMETER_VALUE_SIZE = 8
)

func newsqlparams(key string, value interface{}, param *C.sql_parameter) error{
    param.name = C.Cstring(key)
    //defer C.free(unsafe.Pointer(param.name))
   
    switch t := value.(type) {
    case int32:
        param.type = C.DATA_UINT32
        binary.LittleEndian.PutUint32(param.value[:], value)
    case int64:
        param.type = C.DATA_UINT64
        binary.LittleEndian.PutUint64(param.value[:], value)
    case float64:
        param.type = C.DATA_DOUBLE
        binary.LittleEndian.PutUint32(param.value[:], math.Float64bits(value))
    case string:
        param.type = C.DATA_STRING
        copy(param.value, C.Cstring(value))

    default:
        fmt.Println("not support")
        return errors.New("not support this type")
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
    req_params.param_size = len(queries)

    req_params.params = C.new_sql_parameter(req_params.param_size)
    defer C.free(unsafe.Pointer(params))
 
    int i=0
    for k,v := range queries {
        err := newsqlparams(k, v, &req_params.params[i])
        if i > req_params.param_size || err != nil{
            return nil, errors.New("fail to start a query")
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

    for _,param := range req_params.params {
        C.free(unsafe.Pointer(param.name))
        if param.type = C.DATA_STRING {
            C.free(unsafe.Pointer(param.value))
        }
    }

    if q_res==nil {
        fmt.Println("fail to query")
        return "", errors.New("fail to query")
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
    	return "", errors.New("fail to query")
    }

    str := C.cJSON_Print(q_res)
    defer C.cJSON_Delete(q_res)

    res := C.GoString(str)
    defer C.free(unsafe.Pointer(str))

    return res, nil
}