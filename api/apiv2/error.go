package apiv2

const (
	NONE = iota

	ERRINVALIDBODY = 100
	ERRMCSEVERNOTAVAILABLE = 101
	ERRREQUESTTIMEOUT = 102
	ERRREQUESTFAIL = 103
)

var ErrStringMap = map[int]string {
	ERRINVALIDBODY: 				"Could not parse request body",
	ERRMCSEVERNOTAVAILABLE:			"Mcserver is not available right now",
	ERRREQUESTTIMEOUT:				"Request times out or cancelled",
	ERRREQUESTFAIL:					"Request fails",
}

func ErrMsg(errno int) string {
	if v, ok := ErrStringMap[errno]; ok {
		return v
	}
	return "unknown err"
}