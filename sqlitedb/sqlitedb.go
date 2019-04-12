package sqlitedb

// #cgo CFLAGS: -I../build/include/include
// #cgo LDFLAGS: -L../build/lib -lsqlitedb
// #include <sqlitedb.h>
import "C"
import (
    "unsafe"

    //Log "api-server/log"
)

func InitSqliteDB(db_name string) {
    c_name := C.CString(db_name)
    defer C.free(unsafe.Pointer(c_name))

    C.mcsql_set_db_file(c_name)

    C.register_all_sql_mappers()
}

func decode(cstr *C.char) []byte {
    gostr := C.GoString(cstr)
    defer C.free(unsafe.Pointer(cstr))

    return []byte(gostr)
}

func GetAllArc() []byte {
    return decode(C.mcsql_arc_get_all())
}

func GetArcParams(file_no int32, group string) []byte {
    c_group := C.CString(group)
    defer C.free(unsafe.Pointer(c_group))

    return decode(C.mcsql_arc_get_params(C.int32_t(file_no), c_group))
}

func GetAllBookprograms() []byte {
    return decode(C.mcsql_bookprogram_get_all())
}

func GetAllDynamics() []byte {
    return decode(C.mcsql_dynamics_get_all())
}

func GetDynamicsById(id string) []byte {
    c_id := C.CString(id)
    defer C.free(unsafe.Pointer(c_id))

    return decode(C.mcsql_dynamics_get_by_id(c_id))
}

func GetAllEnum() []byte {
    return decode(C.mcsql_enum_get_all())
}

func GetAllExtaxis() []byte {
    return decode(C.mcsql_extaxis_get_all())
}

func GetAllInterference() []byte {
    return decode(C.mcsql_interference_get_all())
}

func GetAllIO(group string, lang string, auth int32, tech int32) []byte {
    c_group := C.CString(group)
    defer C.free(unsafe.Pointer(c_group))

    c_lang := C.CString(lang)
    defer C.free(unsafe.Pointer(c_lang))

    return decode(C.mcsql_ios_get_all(c_group, c_lang, C.int32_t(auth), C.int32_t(tech)))
}

func GetAllMetadata(lang string) []byte {
    c_lang := C.CString(lang)
    defer C.free(unsafe.Pointer(c_lang))

    return decode(C.mcsql_metadata_get_all(c_lang))
}

func GetAllOperationRecord(created_time int32, start int32, page_size int32) []byte {
    return decode(C.mcsql_operation_record_get_all(C.int32_t(created_time), C.int32_t(start), C.int32_t(page_size)))
}

func GetParams() []byte {
    return decode(C.mcsql_params_get_params())
}

func GetParameterById(id string) []byte {
    c_id := C.CString(id)
    defer C.free(unsafe.Pointer(c_id))

    return decode(C.mcsql_params_get_valid_param_by_id(c_id))
}

func GetParameterByGroup(group string) []byte {
    c_group := C.CString(group)
    defer C.free(unsafe.Pointer(c_group))

    return decode(C.mcsql_params_get_valid_param_by_group(c_group))
}

func GetAllRef() []byte {
    return decode(C.mcsql_ref_get_all())
}

func GetAllToolframe() []byte {
    return decode(C.mcsql_toolframe_get_all())
}

func GetToolframeByToolNo(tool_no int32) []byte {
    return decode(C.mcsql_toolframe_get_by_toolno(C.int32_t(tool_no)))
}

func GetAllUserframe() []byte {
    return decode(C.mcsql_userframe_get_all())
}

func GetUserframeByUserNo(user_no int32) []byte {
    return decode(C.mcsql_userframe_get_by_userno(C.int32_t(user_no)))
}

func GetAllZeropoint() []byte {  
    return decode(C.mcsql_zeropoint_get_all())
}