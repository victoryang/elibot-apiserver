package sqlitedb

import (
    Log "elibot-apiserver/log"
)

func GetAllArc() ([]byte, error) {
    Log.Debug("in GetAllArc")
    
    sm := new(ArcSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetArcParams(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in GetArcParams")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    sm := new(ArcSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_ARC_GET_PARAMS, vars)
}

func GetAllBookprograms() ([]byte, error){
    Log.Debug("in GetAllBookprograms")
 
    sm := new(BookProgramSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllEnum() ([]byte, error) {
    Log.Debug("in GetAllEnum")
    
    sm := new(EnumSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllExtaxis() ([]byte, error) {
    Log.Debug("in GetAllExtaxis")
    
    sm := new(ExtaxisSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllInterference() ([]byte, error) {
    Log.Debug("in GetAllInterference")
    
    sm := new(InterferenceSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllIO() ([]byte, error) {
    Log.Debug("in GetAllIO")
    
    sm := new(IoSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllMetadata(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in GetAllMetadata")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(MetadataSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, vars)
}

func GetParams() ([]byte, error) {
    Log.Debug("in GetParams")
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_PARAMS, nil)
}

func GetParameterById(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in GetParameterById")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_BY_ID, vars)
}

func GetParameterByGroup(vars map[string]interface{}) ([]byte, error) {
    Log.Debug("in GetParameterByGroup")
    for k,v := range vars {
        Log.Debug("key: ", k, " v ", v)
    }
    
    sm := new(ParameterSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_PARAMETER_GET_BY_GROUP, vars)
}

func GetAllRef() ([]byte, error) {
    Log.Debug("in GetAllRef")
    
    sm := new(RefSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_REF_GET_ALL, nil)
}

func GetAllToolframe() ([]byte, error) {
    Log.Debug("in GetAllToolframe")
    
    sm := new(ToolframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllUserframe() ([]byte, error) {
    Log.Debug("in GetAllUserframe")

    sm := new(UserframeSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}

func GetAllZeropoints() ([]byte, error) {
    Log.Debug("in GetAllZeropoint")
    
    sm := new(ZeroPointSqlMapper)
    return RegisterAndQuery(sm, ELIBOT_GET_ALL_PARAMS, nil)
}