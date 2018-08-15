package apiv1

import (
	"net/http"
	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

func RegisterAPIv1(r *mux.Router) http.Handler {
	Log.Debug("Register v1 api...")

	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/health", handleHealth).Methods("GET")
	resourceapi := r.PathPrefix("/v1/resource").Subrouter()
	resourceapi.HandleFunc("/sysvar/crobb", getcRobB).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/irobi", getiRobI).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/drobd", getdRobD).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/drobp", getdRobP).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/drobv", getdRobV).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/locvar/croblb", getcRobLB).Methods("GET").Queries("start", "{start}", "end", "{end}", "num", "{num}")
	resourceapi.HandleFunc("/locvar/irobli", getiRobLI).Methods("GET").Queries("start", "{start}", "end", "{end}", "num", "{num}")
	resourceapi.HandleFunc("/locvar/drobld", getdRobLD).Methods("GET").Queries("start", "{start}", "end", "{end}", "num", "{num}")
	resourceapi.HandleFunc("/locvar/droblp", getdRobLP).Methods("GET").Queries("start", "{start}", "end", "{end}", "num", "{num}")
	resourceapi.HandleFunc("/locvar/droblv", getdRobLV).Methods("GET").Queries("start", "{start}", "end", "{end}", "num", "{num}")
	resourceapi.HandleFunc("/", getResourceOnce).Methods("GET")
	resourceapi.HandleFunc("/nv", getNVOnce).Methods("GET")

	dbapi := r.PathPrefix("/v1/db").Subrouter()
	dbapi.HandleFunc("/backup", DBBackupDB).Methods("POST")
	dbapi.HandleFunc("/backup", DBListBackups).Methods("GET")
	dbapi.HandleFunc("/backup/{name}", DBDelBackup).Methods("DELETE")
	dbapi.HandleFunc("/backup/{name}/restore", DBRestoreBackup).Methods("POST")

	robotapi := r.PathPrefix("/v1/robot").Subrouter()
	repositoryapi := robotapi.PathPrefix("/repository").Subrouter()
	repositoryapi.HandleFunc("/arc", getAllArc).Methods("GET")
	repositoryapi.HandleFunc("/arcparams/{file_no}", getArcParams).Methods("GET").Queries("group", "{group}")
	repositoryapi.HandleFunc("/arcparam/{file_no}/{md_id}", setArcParam).Methods("PUT")
	repositoryapi.HandleFunc("/bookprograms", getAllBookprograms).Methods("GET")
	repositoryapi.HandleFunc("/enum", getAllEnum).Methods("GET")
	repositoryapi.HandleFunc("/extaxis", getAllExtaxis).Methods("GET")
	repositoryapi.HandleFunc("/interference", getAllInterference).Methods("GET")
	repositoryapi.HandleFunc("/interference/{no}/{md_id}", setInterference).Methods("PUT")
	repositoryapi.HandleFunc("/ios/{group}", getAllIos).Methods("GET").Queries("lang", "{lang}", "auth", "{auth}", "tech", "{tech}")
	repositoryapi.HandleFunc("/metadata", getAllMetadata).Methods("GET").Queries("lang", "{lang}")
	repositoryapi.HandleFunc("/params", getParams).Methods("GET")
	repositoryapi.HandleFunc("/parameter/id/{md_id}", getParameterById).Methods("GET")
	repositoryapi.HandleFunc("/parameter/group/{group}", getParameterByGroup).Methods("GET")
	repositoryapi.HandleFunc("/params/{md_id}", setParam).Methods("PUT")
	repositoryapi.HandleFunc("/ref", getAllRef).Methods("GET")
	repositoryapi.HandleFunc("/toolframe", getAllToolframe).Methods("GET")
	repositoryapi.HandleFunc("/toolframe/{tool_no}/{md_id}/pos/{pos_no}", setToolFrame).Methods("PUT")
	repositoryapi.HandleFunc("/userframe", getAllUserframe).Methods("GET")
	repositoryapi.HandleFunc("/userframe/{userno}/{md_id}", setUserFrame).Methods("PUT")
	repositoryapi.HandleFunc("/zeropoint", getAllZeroPoint).Methods("GET")
	repositoryapi.HandleFunc("/zeropoint/{md_id}", setZeroPoint).Methods("PUT")

	settingsapi := r.PathPrefix("/v1/settings").Subrouter()
	settingsapi.HandleFunc("/reboot", rebootSystem).Methods("GET")
	settingsapi.HandleFunc("/date", getSystemDate).Methods("GET")
	settingsapi.HandleFunc("/date/{date}", setSystemDate).Methods("PUT")
	settingsapi.HandleFunc("/ip", getSystemIP).Methods("GET")
	settingsapi.HandleFunc("/ip", setSystemIP).Methods("PUT")
	settingsapi.HandleFunc("/kv", getAllSettingsKV).Methods("GET")
	settingsapi.HandleFunc("/kv/{key}", getSettingsKV).Methods("GET")
	settingsapi.HandleFunc("/kv/{key}", setSettingsKV).Methods("POST")

	executeapi := robotapi.PathPrefix("/execute").Subrouter()
	executeapi.HandleFunc("/cmd_run", doRunCmd).Methods("POST")
	executeapi.HandleFunc("/cmd_pause/{mode}", doPause).Methods("POST")
	executeapi.HandleFunc("/robotmode/{mode}", setRobotMode).Methods("PUT")
	executeapi.HandleFunc("/cmd_clearalarm", doClearAlarm).Methods("POST")
	executeapi.HandleFunc("/cmd_progreset", doProgReset).Methods("POST")
	executeapi.HandleFunc("/speed/{data}", setSpeed).Methods("PUT")
	executeapi.HandleFunc("/mainprog/{progname}", setMainfile).Methods("PUT")
	executeapi.HandleFunc("/cyclemode/{cyclemode}", setCycleMode).Methods("PUT")

	manualinterpolationapi := robotapi.PathPrefix("/manual").Subrouter()
	manualinterpolationapi.HandleFunc("/coord/{mode}", setCoordinateMode).Methods("PUT")
	manualinterpolationapi.HandleFunc("/cmd_manual/{axis}", doManual).Methods("POST")
	manualinterpolationapi.HandleFunc("/cmd_runforward", doRunForward).Methods("POST")
	manualinterpolationapi.HandleFunc("/cmd_runtozero/{status}", doRunToZero).Methods("POST")
	manualinterpolationapi.HandleFunc("/cmd_stop", doRobotStop).Methods("POST")

	axisctrlapi := robotapi.PathPrefix("/axisctrl").Subrouter()
	axisctrlapi.HandleFunc("/servo/{status}", setServoStatus).Methods("PUT")
	axisctrlapi.HandleFunc("/dragteach/{status}", setDragteachStatus).Methods("PUT")

	return r
}