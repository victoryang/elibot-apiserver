package db

import (
    "fmt"
    sql "elibot-apiserver/sqlitedb"
)

func Get_All_Bookprograms() ([]byte, error){
    fmt.Printf("in Get_All_Bookprograms")
    
	id := sql.ELIBOT_BOOKPROGRAM_GET_ALL
    m, err := sql.Get_bookprogram_sql_mapper(id)
    if err!=nil {
        return nil, err
    }

    res, err := sql.db_query(id)
    if err!=nil {
        fmt.Printf("query fails")
        return nil, err
    }
    fmt.Println(res)

    fmt.Println("get_all_metadatas OK")
}