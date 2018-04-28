package db

import (
    "fmt"
    sql "sqlitedb"
)

func Get_All_Bookprograms() {
    fmt.Printf("in Get_All_Bookprograms")
    
	id := ELIBOT_BOOKPROGRAM_GET_ALL
    m, err := sql.Get_bookprogram_sql_mapper(id)
    if err!=nil {
        return nil, err
    }

    res, err := db_query(id)
    if err!=nil {
        fmt.Printf("query fails")
        return nil, err
    }
    fmt.Println(res)

    fmt.Println("get_all_metadatas OK")
}