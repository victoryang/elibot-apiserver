package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db_query.h>
import "C"
import (
    "unsafe"
    "fmt"
    "errors"
)

type db_query_req_option struct {
    option       C.db_query_req_option
}

func Db_query(q_id string) ([]byte, error){
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString("/root/elibotDB.db")
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

    res := C.db_query(req)
    if res==nil {
    	fmt.Println("fail to query")
    	return nil, errors.New("fail to query")
    }
    fmt.Println(res)
    return nil, nil
}