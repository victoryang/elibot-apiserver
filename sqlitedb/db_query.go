package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db_query.h>
// #include<cJSON.h>
import "C"
import (
    "unsafe"
    "fmt"
    "errors"
)

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