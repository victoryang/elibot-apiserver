package db

import (
    "fmt"
    sql "elibot-apiserver/sqlitedb"
)

const (
    DBName="/root/elibotDB.db"
)

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

func Get_ALL_ArcSqlMapper() (string, error) {
    fmt.Println("in Get_ALL_ArcSqlMapper")
    
    sm := new(sql.ArcSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Bookprograms() (string, error){
    fmt.Println("in Get_All_Bookprograms")
 
    sm := new(sql.BookProgramSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_ALL_ExtaxisSqlMapper() (string, error) {
    fmt.Println("in Get_ALL_ExtaxisSqlMapper")
    
    sm := new(sql.ExtaxisSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Interference() (string, error) {
    fmt.Println("in Get_All_Interference")
    
    sm := new(sql.InterferenceSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_ALL_ToolframeSqlMapper() (string, error) {
    fmt.Println("in Get_ALL_ToolframeSqlMapper")
    
    sm := new(sql.ToolframeSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_ALL_UserframeSqlMapper() (string, error) {
    fmt.Println("in Get_ALL_UserframeSqlMapper")

    sm := new(sql.UserframeSqlMapper)
    return RegisterAndQueryAll(sm)
}

func Get_All_Zeropoints() (string, error) {
    fmt.Println("in Get_All_Zeropoints")
    
    sm := new(sql.ZeroPointSqlMapper)
    return RegisterAndQueryAll(sm)
}

