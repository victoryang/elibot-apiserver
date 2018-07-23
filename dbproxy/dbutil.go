package dbproxy

import (
    "os"
    "errors"
    "io/ioutil"
    "path"
    
    sql "elibot-apiserver/sqlitedb"
    Log "elibot-apiserver/log"
)

var DBName string
var DBPath string
var BackupPath string

func RegisterAndQuery(sm sql.SqlMapper, mode int, vars map[string]interface{}) (res []byte, err error) {
    err = sm.RegisterSqlMapper(mode)
    if err!=nil {
        res = []byte("")
        return
    }

    if vars == nil {
        res, err = sql.Db_query(sm.GetID(), DBName)
    } else {
        res, err = sql.Db_query_with_params(sm.GetID(), DBName, vars)
    }

    Log.Debug(string(res))
    Log.Print("get_all_metadatas OK")
    return
}

func DBSetup(dbname string, backuppath string) {
    DBName = dbname
    DBPath = path.Dir(backuppath)
    BackupPath = backuppath
}

/* For sqlitedb backup, restore and upgrade*/
func DBBackup() int {
    Log.Debug("in DBBackup")
 
    return sql.SqlitedbBackup(DBName, DBPath)
}

func DBList() ([]string, error) {
    Log.Debug("in DBList")

    files, err := ioutil.ReadDir(BackupPath)
    if err != nil {
        return nil, err
    }

    var list []string
    for _, f := range files {
        list = append(list, f.Name())
    }
    return list, nil
}

func DBDel(Name string) error {
    Log.Debug("in DBDel")
    if Name == "" {
        return errors.New("file name is empty")
    }

    filename := BackupPath + Name
    if _, err := os.Stat(filename); os.IsNotExist(err) {
       return errors.New("file does not exist")
    }

    return os.Remove(filename)
}

func DBRestore(BackupName string) int {
    Log.Debug("in DBRestore")
 
    return sql.SqlitedbRestore(DBName, BackupPath, BackupPath+BackupName, 1)
}