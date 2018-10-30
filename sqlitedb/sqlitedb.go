package sqlitedb

import (
    Log "elibot-apiserver/log"
)

func GetAllArc() ([]byte, error) {  
    sm := new(ArcSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetArcParams(vars map[string]interface{}) ([]byte, error) {
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    sm := new(ArcSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_ARC_GET_PARAMS, vars)
}

func GetAllBookprograms() ([]byte, error){ 
    sm := new(BookProgramSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllEnum() ([]byte, error) {  
    sm := new(EnumSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllExtaxis() ([]byte, error) {    
    sm := new(ExtaxisSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllInterference() ([]byte, error) {   
    sm := new(InterferenceSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllIO(vars map[string]interface{}) ([]byte, error) {
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(IoSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, vars)
}

func GetAllMetadata(vars map[string]interface{}) ([]byte, error) {
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(MetadataSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, vars)
}

func GetParams() ([]byte, error) {   
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_PARAMS, nil)
}

func GetParameterById(vars map[string]interface{}) ([]byte, error) {
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_BY_ID, vars)
}

func GetParameterByGroup(vars map[string]interface{}) ([]byte, error) {
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_BY_GROUP, vars)
}

func GetAllRef() ([]byte, error) {  
    sm := new(RefSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_REF_GET_ALL, nil)
}

func GetAllToolframe() ([]byte, error) {  
    sm := new(ToolframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetToolframeByToolNo(vars map[string]interface{}) ([]byte, error) {
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    sm := new(ToolframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_COMMON_GET_TOOLFRAMES, vars)
}

func GetAllUserframe() ([]byte, error) {
    sm := new(UserframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetUserframeByUserNo(vars map[string]interface{}) ([]byte, error) {
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    sm := new(UserframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_USER_FRAME_GET_BY_USER_NO, vars)
}

func GetAllZeropoint() ([]byte, error) {  
    sm := new(ZeroPointSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}