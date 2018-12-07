package apiv1

const (
	NONE = iota

	ERRINVALIDBODY = 100
	ERRINCORRECTRANGE = 101
	ERRINVALIDREQUEST = 102

	ERRMCSEVERNOTAVAILABLE = 200
	ERRREQUESTTIMEOUT = 201
	ERRREQUESTFAIL = 202
	ERRQUERY = 203
	ERR_REQ_INVALID_PARAMETER = 204

	ERRBACKUPDB = 300
	ERRLISTDBS = 301
	ERRDELETEDB = 302
	ERRRESTOREDB = 303

	ERRRUNCMD = 400

	ERRFAILTOLISTUSER = 501
	ERRFAILTOADDUSER = 502
	ERRFAILTOREMOVEUSER = 503
	ERRFAILTOMODIFYUSER = 504
	ERRFAILTOCHANGEPWD = 505
)

var ErrStringMap = map[int]string {
	// 400
	ERRINVALIDBODY: 				"Could not parse request body",
	ERRINCORRECTRANGE:				"Request range is incorrect",
	ERRINVALIDREQUEST:				"Request is invalid",

	// 500
	ERRMCSEVERNOTAVAILABLE:			"Mcserver is not available right now",
	ERRREQUESTTIMEOUT:				"Request times out or cancelled",
	ERRREQUESTFAIL:					"Request fails",
	ERRQUERY:						"Fail to query",
	ERR_REQ_INVALID_PARAMETER:		"Invalid parameter",

	ERRBACKUPDB: 				"Fail to backup db",
	ERRLISTDBS: 				"Fail to list dbs",
	ERRDELETEDB:				"Fail to delete db",
	ERRRESTOREDB: 				"Fail to restore db",

	ERRRUNCMD:					"Fail to run cmd",

	ERRFAILTOLISTUSER: 			"Fail to list user"
	ERRFAILTOADDUSER:			"Fail to add user"
	ERRFAILTOREMOVEUSER: 		"Fail to remove user"
	ERRFAILTOMODIFYUSER:		"Fail to modify user"
	ERRFAILTOCHANGEPWD:			"Fail to change password"
}

func ErrMsg(errno int) string {
	if v, ok := ErrStringMap[errno]; ok {
		return v
	}
	return "unknown err"
}