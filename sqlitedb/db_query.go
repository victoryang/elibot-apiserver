package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db_query.h>
import "C"

type query_req_option struct {
    option C.query_req_option
}

type query_req_parameter struct {
    parameter       *C.db_query_req_parameter
}

type query_req_page struct {
    page            *C.db_query_req_page
}

type query_req struct {
    query_id        C.CString
    conn_str        C.CString
    option          *query_req_option.option
    parameter       *query_req_parameter.parameter
    page            *query_req_page.page
}

func db_query(q_id string) ([]byte, error){
    id := CString(q_id)
    defer C.free(unsafe.Pointer(id))

    opt := &query_req_option{C.DB_QUERY_MODE_STANDARD}
    req := &query_req{
					id,
					C.CString(C.CONN_STRINGS),
					opt,
					0,
					0,
				}

    res := C.db_query(&req)
    if res==0 {
    	fmt.Println("fail to query")
    	return nil, errors.Error("fail to query")
    }
    return res, nil
}