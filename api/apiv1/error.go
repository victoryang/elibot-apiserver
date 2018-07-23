package apiv1

const (
	NONE = iota

	ERRBACKUPDB = 100
	ERRLISTDBS = 101
	ERRDELETEDB = 102
	ERRRESTOREDB = 103

	ERRQUERY = 104
)

var ErrStringMap = map[int]string {
	ERRBACKUPDB: 				"Fail to backup db",
	ERRLISTDBS: 				"Fail to list dbs",
	ERRDELETEDB:				"Fail to delete db",
	ERRRESTOREDB: 				"Fail to restore db",

	ERRQUERY:					"Fail to query",
}

func ErrMsg(errno int) string {
	if v, ok := ErrStringMap[errno]; ok {
		return v
	}
	return "unknown err"
}