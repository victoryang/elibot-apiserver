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
	resourceapi.HandleFunc("/sysvar/crobb", getcrobb).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/irobi", getiRobI).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/drobd", getdRobD).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/drobp", getdRobP).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/sysvar/drobv", getdRobV).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/locvar/croblb/{num}", getcrobLb).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/locvar/irobli/{num}", getiRobLI).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/locvar/drobld/{num}", getdRobLD).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/locvar/droblp/{num}", getdRobLP).Methods("GET").Queries("start", "{start}", "end", "{end}")
	resourceapi.HandleFunc("/locvar/droblv/{num}", getdRobLV).Methods("GET").Queries("start", "{start}", "end", "{end}")

	dbapi := r.PathPrefix("/v1/db").Subrouter()
	dbapi.HandleFunc("/backup", DBBackup).Methods("POST")
	dbapi.HandleFunc("/backup", DBList).Methods("GET")
	dbapi.HandleFunc("/backup/{name}", DBDel).Methods("DELETE")
	dbapi.HandleFunc("/backup/{name}/restore", DBRestore).Methods("POST")

	robotapi := r.PathPrefix("/v1/robot").Subrouter()
	repositoryapi := robotapi.PathPrefix("/repository").Subrouter()
	repositoryapi.HandleFunc("/arc", getAllArc).Methods("GET")
	repositoryapi.HandleFunc("/arcparams", getArcParams).Methods("GET").Queries("file_no", "{file_no}", "group", "{group}")
	repositoryapi.HandleFunc("/arcparam/{file_no}/{md_id}", setArcParam).Methods("PUT")
	repositoryapi.HandleFunc("/bookprograms", getAllBookprograms).Methods("GET")
	repositoryapi.HandleFunc("/enum", getAllEnum).Methods("GET")
	repositoryapi.HandleFunc("/extaxis", getAllExtaxis).Methods("GET")
	repositoryapi.HandleFunc("/interference", getAllInterference).Methods("GET")
	repositoryapi.HandleFunc("/interference/{no}/{md_id}", setInterference).Methods("PUT")
	repositoryapi.HandleFunc("/metadata", getAllMetadata).Methods("GET").Queries("lang", "{lang}")
	repositoryapi.HandleFunc("/params", getParams).Methods("GET")
	repositoryapi.HandleFunc("/parameter/id/{md_id}", getParameterById).Methods("GET")
	repositoryapi.HandleFunc("/parameter/group/{group}", parameterbygroup).Methods("GET")
	repositoryapi.HandleFunc("/params/{md_id}", setParam).Methods("PUT")
	repositoryapi.HandleFunc("/ref", getAllRef).Methods("GET")
	repositoryapi.HandleFunc("/toolframe", getAllToolframe).Methods("GET")
	repositoryapi.HandleFunc("/toolframe/{tool_no}/{md_id}/pos/{pos_no}", setToolFrame).Methods("PUT")
	repositoryapi.HandleFunc("/userframe", getAllUserframe).Methods("GET")
	repositoryapi.HandleFunc("/userframe/{userno}/{md_id}", setUserFrame).Methods("PUT")
	repositoryapi.HandleFunc("/zeropoints", getAllZeroPoints).Methods("GET")
	repositoryapi.HandleFunc("/zeropoint/{md_id}", setZeroPoint).Methods("PUT")

	executeapi := robotapi.PathPrefix("/execute").Subrouter()
	executeapi.HandleFunc("/run/{args}", execCmd).Methods("GET")
	executeapi.HandleFunc("/pause/{args}", execPause).Methods("GET")
	executeapi.HandleFunc("/mode/{mode}", setRobotMode).Methods("PUT")
	executeapi.HandleFunc("/alarm/{args}", clearAlarm).Methods("POST")
	executeapi.HandleFunc("/speed/{data}", setSpeed).Methods("PUT")
	executeapi.HandleFunc("/mainfile/{filename}", setMainfile).Methods("PUT")
	executeapi.HandleFunc("/cyclemode/{cyclemode}", setCycleMode).Methods("PUT")

	manualinterpolationapi := robotapi.PathPrefix("/manualinterpolation").Subrouter()
	manualinterpolationapi.HandleFunc("/coord/{mode}", setCoordinateMode).Methods("PUT")
	manualinterpolationapi.HandleFunc("/stop", robotStop).Methods("GET")
	manualinterpolationapi.HandleFunc("/manual/{key}/{arg}", setManual).Methods("PUT")
	manualinterpolationapi.HandleFunc("/runforward", runForward).Methods("POST")
	manualinterpolationapi.HandleFunc("/runtozero/{zero}", runToZero).Methods("GET")

	axisctrlapi := robotapi.PathPrefix("/axisctrl").Subrouter()
	axisctrlapi.HandleFunc("/dragteach/{enabled}", setDragteach).Methods("PUT")

	/*r.HandleFunc("/v1/aio", getAIO).Methods("GET")
	r.HandleFunc("/v1/aio", setAIO).Methods("PUT")
	r.HandleFunc("/v1/aiodouble", getAIOdouble).Methods("GET")
	r.HandleFunc("/v1/aiodouble", setAIOdouble).Methods("PUT")

	r.HandleFunc("/v1/nvram", saveNvRam).Methods("POST")

	r.HandleFunc("/v1/lisence", inputLisence).Methods("PUT")
	r.HandleFunc("/v1/serialnum", getSerialnum).Methods("GET")
	r.HandleFunc("/v1/version", getVersion).Methods("GET")

	remoteapi := r.PathPrefix("/v1/remote").Subrouter()
	remoteapi.HandleFunc("/stamppara", setStampPara).Methods("PUT")
	remoteapi.HandleFunc("/stampip", modifyStampIp).Methods("PUT")
	remoteapi.HandleFunc("/stampsetting", saveStampSetting).Methods("GET")
	remoteapi.HandleFunc("/stacklevel", setStackLevel).Methods("PUT")
	remoteapi.HandleFunc("/stackrange", setStackRange).Methods("PUT")
	remoteapi.HandleFunc("/feedpoint", setFeedPoint).Methods("PUT")
	remoteapi.HandleFunc("/feedtoppoint", setFeedTopPoint).Methods("PUT")

	extdeviceapi := r.PathPrefix("/v1/extdevice").Subrouter()
	extdeviceapi.HandleFunc("/smtrkdev", setSmtrkdev).Methods("PUT")
	extdeviceapi.HandleFunc("/smtrkdev", saveSmtrkdev).Methods("POST")
	extdeviceapi.HandleFunc("/distsenserdata", setDistSenserdata).Methods("PUT")
	extdeviceapi.HandleFunc("/distsenserdata", saveDistSenserdata).Methods("POST")
	extdeviceapi.HandleFunc("/vision/enable", enableVision).Methods("GET")
	extdeviceapi.HandleFunc("/vision/disable", disableVision).Methods("GET")

	bookprogrameapi := robotapi.PathPrefix("/bookprograme").Subrouter()
	bookprogrameapi.HandleFunc("/", deleteRunFile).Methods("DELETE")
	bookprogrameapi.HandleFunc("/", deleteRunFile).Methods("PUT")
	bookprogrameapi.HandleFunc("/enable", setBookProgEnable).Methods("PUT")

	multipleSinapi := robotapi.PathPrefix("/multipleSin").Subrouter()
	multipleSinapi.HandleFunc("/mulsinx", mulsinx).Methods("PUT")
	multipleSinapi.HandleFunc("/sinpos", setMovxJoint).Methods("PUT")
	multipleSinapi.HandleFunc("/sinpos", gotoSinpos).Methods("GET")

	robotapi.HandleFunc("/parameters", setParameters).Methods("PUT")

	robcalibrateapi := robotapi.PathPrefix("/robcalibrate").Subrouter()
	robcalibrateapi.HandleFunc("/autocalczero", setAutoCalcZero).Methods("PUT").Queries("filename", "{filename}")
	robcalibrateapi.HandleFunc("/autorunzero", setAutoRunZero).Methods("PUT")

	trackdataapi := robotapi.PathPrefix("/trackdata").Subrouter()
	trackdataapi.HandleFunc("/craftnum", setTrackCraftNum).Methods("PUT").Queries("num", "{num}")
	trackdataapi.HandleFunc("/note", setTrackdataNote).Methods("PUT")
	trackdataapi.HandleFunc("/", setTrackdata).Methods("PUT").Queries("property", "{property}", "context", "{context}")
	trackdataapi.HandleFunc("/", saveTrackdata).Methods("POST")
	trackdataapi.HandleFunc("/calibration", startCalibration).Methods("POST").Queries("trigger", "{trigger}")
	trackdataapi.HandleFunc("/pointa", recordPointA).Methods("POST")
	trackdataapi.HandleFunc("/pointb", recordPointB).Methods("POST")
	trackdataapi.HandleFunc("/pointc", recordPointC).Methods("POST")
	trackdataapi.HandleFunc("/referencepoint", recordReferencePoint).Methods("POST")
	trackdataapi.HandleFunc("/gotoReferPos", gotoReferPos).Methods("POST")
	trackdataapi.HandleFunc("/calibrationrsult", getCalibrationRsult).Methods("GET")

	visiondataapi := robotapi.PathPrefix("/visiondata").Subrouter()
	visiondataapi.HandleFunc("/craftnum", setVisionCraftNum).Methods("PUT")
	visiondataapi.HandleFunc("/note", setVisionDataNote).Methods("PUT")
	visiondataapi.HandleFunc("/", setVisionData).Methods("PUT").Queries("property", "{property}", "context", "{context}")
	visiondataapi.HandleFunc("/", saveVisionData).Methods("POST")
	visiondataapi.HandleFunc("/camera", triggerCamera).Methods("POST")
	visiondataapi.HandleFunc("/visionencoval", recordVisionEncoVal).Methods("POST")
	visiondataapi.HandleFunc("/visualcalibration", setVisualCalibration).Methods("GET")
	visiondataapi.HandleFunc("/robencoval", recordRobEncoVal).Methods("POST")
	visiondataapi.HandleFunc("/robotpoint", recordRobotPoint).Methods("POST")
	visiondataapi.HandleFunc("/visualpoint", recordVisualPoint).Methods("POST")
	visiondataapi.HandleFunc("/photo", takePhoto).Methods("POST")
	visiondataapi.HandleFunc("/datarefresh", dataReresh).Methods("POST")
	visiondataapi.HandleFunc("/gotopoint", goToPoint).Methods("POST")*/

	return r
}