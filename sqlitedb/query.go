package sqlitedb

// #cgo CFLAGS: -I/root/mcserver/include -I../thirdparty/mcsql
// #cgo LDFLAGS: -lsqlitedb -L../thirdparty/mcsql -lmcsql
// #include<stdlib.h>
// #include<mcsql.h>
import "C"
import (
    "unsafe"
    "errors"

    Log "elibot-apiserver/log"
)

func Db_query_with_params(q_id, db_name string, vars map[string]interface{}) ([]byte, error) {
    Log.Debug("Db_query_with_params")
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString(db_name)
    defer C.free(unsafe.Pointer(conn))

    p := new(Parameter)
    p.newReqParameter(vars)
    defer p.freeReqParameter()

    buf := C.mcsql_query_with_param(id, conn, C.DB_QUERY_MODE_CUSTOM, p.parameter, nil)
    if buf == nil {
        return nil, errors.New("query empty")
    }
    res := C.GoString(buf)
    defer C.free(unsafe.Pointer(buf))

    return []byte(res), nil
}

func Db_query(q_id string, db_name string) ([]byte, error){
    id := C.CString(q_id)
    defer C.free(unsafe.Pointer(id))

    conn := C.CString(db_name)
    defer C.free(unsafe.Pointer(conn))

    buf := C.mcsql_query(id, conn, C.DB_QUERY_MODE_STANDARD)
    if buf == nil {
        return nil, errors.New("query empty")
    }
    res := C.GoString(buf)
    defer C.free(unsafe.Pointer(buf))
    return []byte(res), nil
}

func RegisterAndQuery(sm SqlMapper, mode int, vars map[string]interface{}) (res []byte, err error) {
    err = sm.RegisterSqlMapper(mode)
    if err!=nil {
        res = []byte("")
        return
    }

    if vars == nil {
        res, err = Db_query(sm.GetID(), DBName)
    } else {
        res, err = Db_query_with_params(sm.GetID(), DBName, vars)
    }

    Log.Debug(string(res))
    Log.Print("get_all_metadatas OK")
    return
}