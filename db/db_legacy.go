package db

import (
    "fmt"
    sql "elibot-apiserver/sqlitedb"
)

func Get_All_Bookprograms() (string, error){
    fmt.Println("in Get_All_Bookprograms")
    
	id := sql.ELIBOT_BOOKPROGRAM_GET_ALL
    err := sql.Get_bookprogram_sql_mapper(id)
    if err!=nil {
        return nil, err
    }

    res, err := sql.Db_query(id)
    if err!=nil {
        fmt.Printf("query fails")
        return "", err
    }

    fmt.Println(res)
    fmt.Println("get_all_metadatas OK")
    return res, nil
}