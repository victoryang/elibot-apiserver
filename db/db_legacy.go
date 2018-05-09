package db

import (
    "fmt"
    sql "elibot-apiserver/sqlitedb"
)

const (
    DBName="/root/elibotDB.db"
)

func RegisterAndQuery(sm sql.SqlMapper, mode int, vars map[string]interface{}) (res string, err error) {
    err = sm.RegisterSqlMapper(mode)
    if err!=nil {
        res = ""
        return
    }

    if vars == nil {
        res, err = sql.Db_query(sm.GetID(), DBName)
    } else {
        res, err = sql.Db_query_with_params(sm.GetID(), DBName, vars)
    }
    if err!=nil {
        fmt.Println("query fails: ", err)
        return
    }

    fmt.Println(res)
    fmt.Println("get_all_metadatas OK")
    return
}

func Get_ALL_Arc() (string, error) {
    fmt.Println("in Get_ALL_Arc")
    
    sm := new(sql.ArcSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_Arc_Params(vars map[string]interface{}) (string, error) {
    fmt.Println("in Get_Arc_Params")
    for k,v := range vars {
        fmt.Println("key: ", k, "v", v)
    }
    sm := new(sql.ArcSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_ARC_GET_PARAMS, vars)
}

func Get_All_Backup() (string, error){
    fmt.Println("in Get_All_Backup")
 
    sm := new(sql.BackupSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Bookprograms() (string, error){
    fmt.Println("in Get_All_Bookprograms")
 
    sm := new(sql.BookProgramSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Enum() (string, error) {
    fmt.Println("in Get_ALL_Enum")
    
    sm := new(sql.EnumSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Extaxis() (string, error) {
    fmt.Println("in Get_ALL_Extaxis")
    
    sm := new(sql.ExtaxisSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Interference() (string, error) {
    fmt.Println("in Get_All_Interference")
    
    sm := new(sql.InterferenceSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_IO() (string, error) {
    fmt.Println("in Get_All_IO")
    
    sm := new(sql.IoSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Metadata(vars map[string]interface{}) (string, error) {
    fmt.Println("in Get_All_Metadata")
    
    sm := new(sql.MetadataSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, vars)
}

func Get_Params() (string, error) {
    fmt.Println("in Get_Parameter_By_Id")
    
    sm := new(sql.ParameterSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_PARAMETER_GET_PARAMS)
}

func Get_Parameter_By_Id(vars map[string]interface{}) (string, error) {
    fmt.Println("in Get_Parameter_By_Id")
    
    sm := new(sql.ParameterSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_PARAMETER_GET_BY_ID, vars)
}

func Get_Parameter_By_Group(vars map[string]interface{}) (string, error) {
    fmt.Println("in Get_Parameter_By_Group")
    
    sm := new(sql.ParameterSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_PARAMETER_GET_BY_GROUP, vars)
}

func Get_All_Ref() (string, error) {
    fmt.Println("in Get_All_Ref")
    
    sm := new(sql.RefSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_REF_GET_ALL, nil)
}

func Get_ALL_Toolframe() (string, error) {
    fmt.Println("in Get_ALL_Toolframe")
    
    sm := new(sql.ToolframeSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Userframe() (string, error) {
    fmt.Println("in Get_ALL_Userframe")

    sm := new(sql.UserframeSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Zeropoints() (string, error) {
    fmt.Println("in Get_All_Zeropoints")
    
    sm := new(sql.ZeroPointSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}