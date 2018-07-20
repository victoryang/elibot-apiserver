package dbproxy

import (
    "os"
    "errors"
    "io/ioutil"
    
    sql "elibot-apiserver/sqlitedb"
    Log "elibot-apiserver/log"
)

var DBName string
var BackupPath string

func RegisterAndQuery(sm sql.SqlMapper, mode int, vars map[string]interface{}) (res string, err error) {
    err = sm.RegisterSqlMapper(mode)
    if err!=nil {
        res = ""
        return
    }

    if vars == nil {
        res, err = sql.Db_query(sm.GetID(), DBName)
    } else {
        var res1 []byte
        res1, err = sql.Db_query_with_params(sm.GetID(), DBName, vars)
        res = string(res1)
    }

    Log.Debug(res)
    Log.Print("get_all_metadatas OK")
    return
}

func DBSetup(dbname string, backuppath string) {
    DBName = dbname
    BackupPath = backuppath
}

/* For sqlitedb backup, restore and upgrade*/
func DBBackup() error {
    Log.Debug("in DBBackup")
 
    mgr, err := sql.NewDBManager(DBName, BackupPath, "", sql.DB_BACKUP, 0)
    if err != nil {
        return err
    }
    return mgr.Exec()
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

func DBRestore(BackupName string) error{
    Log.Debug("in DBRestore")
 
    mgr, err := sql.NewDBManager(DBName, BackupPath, BackupPath+BackupName, sql.DB_RESTORE, 1)
    if err != nil {
        return err
    }
    return mgr.Exec()
}