package apiv1

const (
	NONE = iota

	ERRINVALIDBODY = 100
	ERRMCSEVERNOTAVAILABLE = 101
	ERRREQUESTTIMEOUT = 102
	ERRREQUESTFAIL = 103
	ERRINCORRECTRANGE = 104
	ERRQUERY = 105

	ERRBACKUPDB = 200
	ERRLISTDBS = 201
	ERRDELETEDB = 202
	ERRRESTOREDB = 203

	ERRRUNCMD = 301
)

var ErrStringMap = map[int]string {
	ERRINVALIDBODY: 				"Could not parse request body",
	ERRMCSEVERNOTAVAILABLE:			"Mcserver is not available right now",
	ERRREQUESTTIMEOUT:				"Request times out or cancelled",
	ERRREQUESTFAIL:					"Request fails",
	ERRINCORRECTRANGE:				"Request range is incorrect",
	ERRQUERY:						"Fail to query",

	ERRBACKUPDB: 				"Fail to backup db",
	ERRLISTDBS: 				"Fail to list dbs",
	ERRDELETEDB:				"Fail to delete db",
	ERRRESTOREDB: 				"Fail to restore db",
	ERRRUNCMD:					"Fail to run cmd",
}

func ErrMsg(errno int) string {
	if v, ok := ErrStringMap[errno]; ok {
		return v
	}
	return "unknown err"
}