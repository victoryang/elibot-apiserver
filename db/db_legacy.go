package db

import (
    "fmt"
    sql "elibot-apiserver/sqlitedb"
)

const (
    DBName="/root/elibotDB.db"
)
func Get_All_Zeropoints() (string, error) {
    fmt.Println("in Get_All_Zeropoints")
    
    sm := new(sql.ZeroPointSqlMapper)
    err := sm.RegisterSqlMapper()
    if err!=nil {
        return "", err
    }

    res, err := sql.Db_query(sm.Id, DBName)
    if err!=nil {
        fmt.Printf("query fails")
        return "", err
    }

    fmt.Println(res)
    fmt.Println("get_all_metadatas OK")
    return res, nil
}

func Get_All_Bookprograms() (string, error){
    fmt.Println("in Get_All_Bookprograms")
    
	sm := new(sql.BookProgramSqlMapper)
    err := sm.RegisterSqlMapper()
    if err!=nil {
        return "", err
    }

    res, err := sql.Db_query(sm.Id, DBName)
    if err!=nil {
        fmt.Printf("query fails")
        return "", err
    }

    fmt.Println(res)
    fmt.Println("get_all_metadatas OK")
    return res, nil
}