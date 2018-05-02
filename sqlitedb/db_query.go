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

/*type query_req_option struct {
    option C.db_query_req_option
}

type query_req_parameter struct {
    parameter       *C.db_query_req_parameter
}

type query_req_page struct {
    page            *C.db_query_req_page
}*/

type query_req struct {
    query_id        string
    conn_str        string
    option          *C.db_query_req_option
    parameter       *C.db_query_req_parameter
    page            *C.db_query_req_page
}

func db_query(q_id string) ([]byte, error){
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString("/root/elibotDB.db")
    defer C.free(unsafe.Pointer(conn))

    opt := &C.db_query_req_option{C.DB_QUERY_MODE_STANDARD}
    req := &query_req{
					id,
					conn,
					opt,
					0,
					0,
				}

    res := C.db_query(req)
    if res==0 {
    	fmt.Println("fail to query")
    	return nil, errors.Error("fail to query")
    }
    return res, nil
}