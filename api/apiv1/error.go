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

	// User management
	ERRFAILTOLISTUSER = 501
	ERRFAILTOOPERATEUSER = 502
	ERRFAILTOOPERATEPWD = 503

	// Err From Modules
	ERRADDEXISTUSER = 102067
	ERRCONVERTPWD = 200001
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

	ERRFAILTOLISTUSER: 			"Fail to list user",
	ERRFAILTOOPERATEUSER:		"Fail to operate user",
	ERRFAILTOOPERATEPWD:		"Fail to operate password",

	// User Management
	ERRADDEXISTUSER:			"User exists",
	ERRCONVERTPWD:				"Wrong format password",
}

func ErrMsg(errno int) string {
	if v, ok := ErrStringMap[errno]; ok {
		return v
	}
	return "unknown err"
}