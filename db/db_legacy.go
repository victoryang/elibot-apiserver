package db

import (
    "fmt"
    sql "elibot-apiserver/sqlitedb"
)

func Get_All_Bookprograms() (string, error){
    fmt.Println("in Get_All_Bookprograms")
    
	sm := new(sql.BookProgramSqlMapper)
    err := sm.RegisterSqlMapper()
    if err!=nil {
        return "", err
    }

    res, err := sql.Db_query(sm.Id, "/root/elibotDB.db")
    if err!=nil {
        fmt.Printf("query fails")
        return "", err
    }

    fmt.Println(res)
    fmt.Println("get_all_metadatas OK")
    return res, nil
}