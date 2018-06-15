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
func DBsetup(dbname string, backuppath string) {
    DBName = dbname
    BackupPath = backuppath
}

func RegisterAndQuery(sm sql.SqlMapper, mode int, vars map[string]interface{}) (res string, err error) {
    err = sm.RegisterSqlMapper(mode)
    if err!=nil {
        res = ""
        return
    }

    if vars == nil {
        res, err = sql.Db_query(sm.GetID(), DBName)
    } else {
        res, err = sql.Db_query_with_params(sm.GetID(), DBName, vars)
    }
    if err!=nil {
        Log.Error("query fails: ", err)
        return
    }

    Log.Debug(res)
    Log.Print("get_all_metadatas OK")
    return
}

func Get_ALL_Arc() (string, error) {
    Log.Debug("in Get_ALL_Arc")
    
    sm := new(sql.ArcSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_Arc_Params(vars map[string]interface{}) (string, error) {
    Log.Debug("in Get_Arc_Params")
    for k,v := range vars {
        Log.Debug("key: ", k, "v", v)
    }
    sm := new(sql.ArcSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_ARC_GET_PARAMS, vars)
}

func Get_All_Bookprograms() (string, error){
    Log.Debug("in Get_All_Bookprograms")
 
    sm := new(sql.BookProgramSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Enum() (string, error) {
    Log.Debug("in Get_ALL_Enum")
    
    sm := new(sql.EnumSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Extaxis() (string, error) {
    Log.Debug("in Get_ALL_Extaxis")
    
    sm := new(sql.ExtaxisSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Interference() (string, error) {
    Log.Debug("in Get_All_Interference")
    
    sm := new(sql.InterferenceSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_IO() (string, error) {
    Log.Debug("in Get_All_IO")
    
    sm := new(sql.IoSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Metadata(vars map[string]interface{}) (string, error) {
    Log.Debug("in Get_All_Metadata")
    
    sm := new(sql.MetadataSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, vars)
}

func Get_Params() (string, error) {
    Log.Debug("in Get_Parameter_By_Id")
    
    sm := new(sql.ParameterSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_PARAMETER_GET_PARAMS, nil)
}

func Get_Parameter_By_Id(vars map[string]interface{}) (string, error) {
    Log.Debug("in Get_Parameter_By_Id")
    
    sm := new(sql.ParameterSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_PARAMETER_GET_BY_ID, vars)
}

func Get_Parameter_By_Group(vars map[string]interface{}) (string, error) {
    Log.Debug("in Get_Parameter_By_Group")
    
    sm := new(sql.ParameterSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_PARAMETER_GET_BY_GROUP, vars)
}

func Get_All_Ref() (string, error) {
    Log.Debug("in Get_All_Ref")
    
    sm := new(sql.RefSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_REF_GET_ALL, nil)
}

func Get_ALL_Toolframe() (string, error) {
    Log.Debug("in Get_ALL_Toolframe")
    
    sm := new(sql.ToolframeSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Userframe() (string, error) {
    Log.Debug("in Get_ALL_Userframe")

    sm := new(sql.UserframeSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Zeropoints() (string, error) {
    Log.Debug("in Get_All_Zeropoints")
    
    sm := new(sql.ZeroPointSqlMapper)
    return RegisterAndQuery(sm, sql.ELIBOT_GET_ALL_PARAMS, nil)
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