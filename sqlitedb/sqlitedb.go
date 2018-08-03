package sqlitedb

import (
    "elibot-apiserver/sqlitedb/mapper"
    Log "elibot-apiserver/log"
)

func Get_ALL_Arc() ([]byte, error) {
    Log.Debug("in Get_ALL_Arc")
    
    sm := new(ArcSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_Arc_Params(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in Get_Arc_Params")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    sm := new(ArcSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_ARC_GET_PARAMS, vars)
}

func Get_All_Bookprograms() ([]byte, error){
    Log.Debug("in Get_All_Bookprograms")
 
    sm := new(BookProgramSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Enum() ([]byte, error) {
    Log.Debug("in Get_ALL_Enum")
    
    sm := new(EnumSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Extaxis() ([]byte, error) {
    Log.Debug("in Get_ALL_Extaxis")
    
    sm := new(ExtaxisSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Interference() ([]byte, error) {
    Log.Debug("in Get_All_Interference")
    
    sm := new(InterferenceSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_IO() ([]byte, error) {
    Log.Debug("in Get_All_IO")
    
    sm := new(IoSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Metadata(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in Get_All_Metadata")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(MetadataSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, vars)
}

func Get_Params() ([]byte, error) {
    Log.Debug("in Get_Params")
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_PARAMS, nil)
}

func Get_Parameter_By_Id(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in Get_Parameter_By_Id")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_BY_ID, vars)
}

func Get_Parameter_By_Group(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in Get_Parameter_By_Group")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_BY_GROUP, vars)
}

func Get_All_Ref() ([]byte, error) {
    Log.Debug("in Get_All_Ref")
    
    sm := new(RefSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_REF_GET_ALL, nil)
}

func Get_ALL_Toolframe() ([]byte, error) {
    Log.Debug("in Get_ALL_Toolframe")
    
    sm := new(ToolframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_ALL_Userframe() ([]byte, error) {
    Log.Debug("in Get_ALL_Userframe")

    sm := new(UserframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func Get_All_Zeropoints() ([]byte, error) {
    Log.Debug("in Get_All_Zeropoints")
    
    sm := new(ZeroPointSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}