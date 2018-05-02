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

/*type query_req struct {
    query_id        *C.char
    conn_str        *C.char
    option          *C.db_query_req_option
    parameter       *C.db_query_req_parameter
    page            *C.db_query_req_page
}*/

func Db_query(q_id string) ([]byte, error){
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString("/root/elibotDB.db")
    defer C.free(unsafe.Pointer(conn))

    opt := &C.db_query_req_option{C.DB_QUERY_MODE_STANDARD}
    req := &C.db_query_req{
					id,
					conn,
					opt,
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