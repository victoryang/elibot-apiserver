package sqlitedb

// #cgo CFLAGS: -I/root/mcserver/include -I../thirdparty/mcsql
// #cgo LDFLAGS: -lsqlitedb -L../thirdparty/mcsql -lmcsql
// #include<mcsqlmapper.h>
import "C"
import (
	Log "elibot-apiserver/log"
)
// SqlMapper defines methods of a sql mapper
type SqlMapper interface {
	RegisterSqlMapper(mode int)		error
	GetID()		                    string
}

const (
	/* Base action that sqlmappers should take*/
	ELIBOT_GET_ALL_PARAMS = iota
	ELIBOT_UPDATE_PARAMS

	/*For arc sqlmapper*/
	ELIBOT_ARC_GET_PARAMS

	/*For extaxis sqlmapper*/
	ELIBOT_EXTAXIS_UPDATE
	ELIBOT_EXTAXIS_UPDATE_POS

	/*For io sqlmapper*/
	ELIBOT_IO_GET_VALID_IOS_BY_GROUP

	/*For parameter sqlmapper*/
	ELIBOT_PARAMETER_GET_PARAMS
	ELIBOT_PARAMETER_GET_BY_ID
	ELIBOT_PARAMETER_GET_BY_GROUP

	/*For ref sqlmapper*/
	ELIBOT_REF_GET_ALL
)

func InitSqlitedb() {
	Log.Debug("init sqlitedb sqlmappers")
	C.register_all_sql_mappers()
}