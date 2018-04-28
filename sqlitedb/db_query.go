package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db/db_context.h>

import "C"

func db_query(q_id string) ([]byte, error){
    opt := C.db_query_req_option{C.DB_QUERY_MODE_STANDARD}
    var req C.db_query_req = C.db_query_req{
    							C.CString(q_id),
    							C.CString(C.CONN_STRINGS),
    							&opt,
    							0,
    							0
    						}

    res := C.db_query(&req)
    if res==0 {
    	fmt.Println("fail to query")
    	return nil, errors.Error("fail to query")
    }
    return res, nil
}