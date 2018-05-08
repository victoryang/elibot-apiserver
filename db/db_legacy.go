package db

import (
    "fmt"
    sql "elibot-apiserver/sqlitedb"
)

const (
    DBName="/root/elibotDB.db"
)

func RegisterAndQueryWithParams(sm sql.SqlMapper, vars map[string]interface{}) (string, error) {
    err := sm.RegisterSqlMapper(sql.ELIBOT_GET_WITH_PARAMS)
    if err!=nil {
        return "", err
    }

    res, err := sql.Db_query_with_params(sm.GetID(), DBName, vars)
    if err!=nil {
        fmt.Printf("query fails")
        return "", err
    }

    fmt.Println(res)
    fmt.Println("get_all_metadatas OK")
    return res, nil
}

func RegisterAndQueryAll(sm sql.SqlMapper) (string, error) {
    err := sm.RegisterSqlMapper(sql.ELIBOT_GET_ALL_PARAMS)
    if err!=nil {
        return "", err
    }

    res, err := sql.Db_query(sm.GetID(), DBName)
    if err!=nil {
        fmt.Printf("query fails")
        return "", err
    }

    fmt.Println(res)
    fmt.Println("get_all_metadatas OK")
    return res, nil
}

func Get_ALL_Arc() (string, error) {
    fmt.Println("in Get_ALL_Arc")
    
    sm := new(sql.ArcSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_Arc_Params(vars map[string]interface{}) (string, error) {
    fmt.Println("in Get_Arc_Params")
    for k,v := range vars {
        fmt.Println("key: ", k, "v", v)
    }
    sm := new(sql.ArcSqlMapper)
    return RegisterAndQueryWithParams(sm, vars)
}

func Get_All_Backup() (string, error){
    fmt.Println("in Get_All_Backup")
 
    sm := new(sql.BackupSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Bookprograms() (string, error){
    fmt.Println("in Get_All_Bookprograms")
 
    sm := new(sql.BookProgramSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_ALL_Enum() (string, error) {
    fmt.Println("in Get_ALL_Enum")
    
    sm := new(sql.EnumSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_ALL_Extaxis() (string, error) {
    fmt.Println("in Get_ALL_Extaxis")
    
    sm := new(sql.ExtaxisSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Interference() (string, error) {
    fmt.Println("in Get_All_Interference")
    
    sm := new(sql.InterferenceSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_IO() (string, error) {
    fmt.Println("in Get_All_IO")
    
    sm := new(sql.IoSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Metadata() (string, error) {
    fmt.Println("in Get_All_Metadata")
    
    sm := new(sql.MetadataSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Parameter() (string, error) {
    fmt.Println("in Get_All_Parameter")
    
    sm := new(sql.ParameterSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Ref() (string, error) {
    fmt.Println("in Get_All_Ref")
    
    sm := new(sql.RefSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_ALL_Toolframe() (string, error) {
    fmt.Println("in Get_ALL_Toolframe")
    
    sm := new(sql.ToolframeSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_ALL_Userframe() (string, error) {
    fmt.Println("in Get_ALL_Userframe")

    sm := new(sql.UserframeSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Zeropoints() (string, error) {
    fmt.Println("in Get_All_Zeropoints")
    
    sm := new(sql.ZeroPointSqlMapper)
    return RegisterAndQueryAll(sm)
}

