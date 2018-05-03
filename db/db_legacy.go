package db

import (
    "fmt"
    "errors"
    sql "elibot-apiserver/sqlitedb"
)

const (
    DBName="/root/elibotDB.db"
)

func RegisterAndQueryAll(sm interface{}) (string, error) {
    switch sm.(type) {
    case *sql.InterferenceSqlMapper:

    default:
        fmt.Println("type not correct")
        return "", errors.New("type not correct")
    }

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

func Get_All_Interference() (string, error) {
    fmt.Println("in Get_All_Interference")
    
    sm := new(sql.InterferenceSqlMapper)
    return RegisterAndQueryAll(sm)
}

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