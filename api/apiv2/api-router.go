package apiv2

import (

	Log "elibot-apiserver/log"

	"github.com/gorilla/mux"
)

func RegisterAPIv2(r *mux.Router) {
	Log.Debug("Register V2 api...")

	r.HandleFunc("/", hello).Methods("GET")
	r.HandleFunc("/health", handleHealth).Methods("GET")

	/*r.HandleFunc("/v2/aio", getAIO).Methods("GET")
	r.HandleFunc("/v2/aio", setAIO).Methods("PUT")
	r.HandleFunc("/v2/aiodouble", getAIOdouble).Methods("GET")
	r.HandleFunc("/v2/aiodouble", setAIOdouble).Methods("PUT")

	r.HandleFunc("/v2/nvram", saveNvRam).Methods("POST")

	r.HandleFunc("/v2/lisence", inputLisence).Methods("PUT")
	r.HandleFunc("/v2/serialnum", getSerialnum).Methods("GET")
	r.HandleFunc("/v2/version", getVersion).Methods("GET")

	remoteapi := r.PathPrefix("/v2/remote").Subrouter()
	remoteapi.HandleFunc("/stamppara", setStampPara).Methods("PUT")
	remoteapi.HandleFunc("/stampip", modifyStampIp).Methods("PUT")
	remoteapi.HandleFunc("/stampsetting", saveStampSetting).Methods("GET")
	remoteapi.HandleFunc("/stacklevel", setStackLevel).Methods("PUT")
	remoteapi.HandleFunc("/stackrange", setStackRange).Methods("PUT")
	remoteapi.HandleFunc("/feedpoint", setFeedPoint).Methods("PUT")
	remoteapi.HandleFunc("/feedtoppoint", setFeedTopPoint).Methods("PUT")

	extdeviceapi := r.PathPrefix("/v2/extdevice").Subrouter()
	extdeviceapi.HandleFunc("/smtrkdev", setSmtrkdev).Methods("PUT")
	extdeviceapi.HandleFunc("/smtrkdev", saveSmtrkdev).Methods("POST")
	extdeviceapi.HandleFunc("/distsenserdata", setDistSenserdata).Methods("PUT")
	extdeviceapi.HandleFunc("/distsenserdata", saveDistSenserdata).Methods("POST")
	extdeviceapi.HandleFunc("/vision/enable", enableVision).Methods("GET")
	extdeviceapi.HandleFunc("/vision/disable", disableVision).Methods("GET")*/

	robotapi := r.PathPrefix("/v2/robot").Subrouter()
	/*arcweldingapi := robotapi.PathPrefix("/arcwelding").Subrouter()
	arcweldingapi.HandleFunc("/arcen", setArcEn).Methods("PUT")
	arcweldingapi.HandleFunc("/diandong", setDiandong).Methods("PUT")
	arcweldingapi.HandleFunc("/songqi", setSongqi).Methods("PUT")
	arcweldingapi.HandleFunc("/arcweldpara", setArcWeldPara).Methods("PUT")
	arcweldingapi.HandleFunc("/arcweldpara", saveArcWeldPara).Methods("POST")
	arcweldingapi.HandleFunc("/weavepara", setWeavePara).Methods("PUT")
	arcweldingapi.HandleFunc("/weavepara", saveWeavePara).Methods("POST")

	autosettoolapi := robotapi.PathPrefix("/autosettool").Subrouter()
	autosettoolapi.HandleFunc("/toolframe", setToolFrame).Methods("PUT")
	autosettoolapi.HandleFunc("/autotoolframe", autoSetToolFrame).Methods("PUT")
	autosettoolapi.HandleFunc("/toolpos", setToolPos).Methods("PUT")
	autosettoolapi.HandleFunc("/toolnum", setToolNum).Methods("PUT")
	autosettoolapi.HandleFunc("/toolnote", setToolNote).Methods("PUT")

	bookprogrameapi := robotapi.PathPrefix("/bookprograme").Subrouter()
	bookprogrameapi.HandleFunc("/", deleteRunFile).Methods("DELETE")
	bookprogrameapi.HandleFunc("/", deleteRunFile).Methods("PUT")
	bookprogrameapi.HandleFunc("/enable", setBookProgEnable).Methods("PUT")

	executeapi := robotapi.PathPrefix("/execute").Subrouter()
	executeapi.HandleFunc("/run", ).Methods("GET")
	executeapi.HandleFunc("/pause", ).Methods("GET")
	executeapi.HandleFunc("/mode", setRobotMode).Methods("PUT")
	executeapi.HandleFunc("/origin", getOrigin).Methods("GET")
	executeapi.HandleFunc("/origin", setOrigin).Methods("POST")
	executeapi.HandleFunc("/curpos", resetCurPos).Methods("POST")
	executeapi.HandleFunc("/alarm", clearAlarm).Methods("POST")
	executeapi.HandleFunc("/prog", progReset).Methods("POST")
	executeapi.HandleFunc("/speed", setSpeed).Methods("PUT")
	executeapi.HandleFunc("/mainfile", setMainfile).Methods("PUT")
	executeapi.HandleFunc("/cyclemode", setCycleMode).Methods("PUT")
	executeapi.HandleFunc("/exiteditprog", exitEditProg).Methods("PUT")
	executeapi.HandleFunc("/hand", setHand).Methods("PUT")
	executeapi.HandleFunc("/load", loadFile).Methods("POST")
	executeapi.HandleFunc("/idump", instructionDump).Methods("POST")
	executeapi.HandleFunc("/jo2pose", joint2pose).Methods("POST")
	executeapi.HandleFunc("/servodisable", servoDisable).Methods("PUT")
	executeapi.HandleFunc("/technics", setTechnics).Methods("PUT")
	executeapi.HandleFunc("/help", gethelp).Methods("GET")

	extaxisapi := robotapi.PathPrefix("/extaxis").Subrouter()
	extaxisapi.HandleFunc("/", setExtAxis).Methods("PUT")
	extaxisapi.HandleFunc("/", setExtAxis).Methods("DELETE")
	extaxisapi.HandleFunc("/cooperate", setExtAxisCooperate).Methods("PUT")
	extaxisapi.HandleFunc("/gotopos", gotoExtAxisPos).Methods("POST")
	extaxisapi.HandleFunc("/pos", setExtAxisPos).Methods("PUT")

	interferenceapi := robotapi.PathPrefix("/interference").Subrouter()
	interferenceapi.HandleFunc("/", setInterference).Methods("PUT")
	interferenceapi.HandleFunc("/pos", setInterferencePos).Methods("PUT")
	interferenceapi.HandleFunc("/gotopos", gotoInterferencepos).Methods("POST")
	interferenceapi.HandleFunc("/", resetInterference).Methods("DELETE")

	manualinterpolationapi := robotapi.PathPrefix("/manualinterpolation").Subrouter()
	manualinterpolationapi.HandleFunc("/coord", setCoordinateMode).Methods("PUT")
	manualinterpolationapi.HandleFunc("/stop", robotStop).Methods("GET")
	manualinterpolationapi.HandleFunc("/manual", setManual).Methods("PUT")
	manualinterpolationapi.HandleFunc("/runforward", runForward).Methods("POST")
	manualinterpolationapi.HandleFunc("/runtozero", runToZero).Methods("GET")

	multipleSinapi := robotapi.PathPrefix("/multipleSin").Subrouter()
	multipleSinapi.HandleFunc("/mulsinx", mulsinx).Methods("PUT")
	multipleSinapi.HandleFunc("/sinpos", setMovxJoint).Methods("PUT")
	multipleSinapi.HandleFunc("/sinpos", gotoSinpos).Methods("GET")

	robotapi.HandleFunc("/parameters", setParameters).Methods("PUT")

	robcalibrateapi := robotapi.PathPrefix("/robcalibrate").Subrouter()
	robcalibrateapi.HandleFunc("/autocalczero", setAutoCalcZero).Methods("PUT")
	robcalibrateapi.HandleFunc("/autorunzero", setAutoRunZero).Methods("PUT")

	trackdataapi := robotapi.PathPrefix("/trackdata").Subrouter()
	trackdataapi.HandleFunc("/calipoint", setCaliPoint).Methods("PUT")
	trackdataapi.HandleFunc("/", setTrackdata).Methods("PUT")
	trackdataapi.HandleFunc("/", saveTrackdata).Methods("POST")
	trackdataapi.HandleFunc("/trackencode", getTrackEncode).Methods("GET")

	userframeapi := robotapi.PathPrefix("/userframe").Subrouter()
	userframeapi.HandleFunc("/", setUserFrame).Methods("PUT")
	userframeapi.HandleFunc("/userpos", setUserPos).Methods("PUT")
	userframeapi.HandleFunc("/userpos", gotoUserPos).Methods("POST")
	userframeapi.HandleFunc("/usernum", setUserNum).Methods("PUT")
	userframeapi.HandleFunc("/usernote", setUserNote).Methods("PUT")*/

	visiondataapi := robotapi.PathPrefix("/visiondata").Subrouter()
	visiondataapi.HandleFunc("/craftnum", setVisionCraftNum).Methods("PUT")
	visiondataapi.HandleFunc("/note", setVisionDataNote).Methods("PUT")
	/*visiondataapi.HandleFunc("/", getVisiondata).Methods("GET")
	visiondataapi.HandleFunc("/", setVisiondata).Methods("PUT")
	visiondataapi.HandleFunc("/currentencoval", getCurrentEncoVal).Methods("GET")
	visiondataapi.HandleFunc("/visionencoval", getVisionEncoVal).Methods("GET")
	visiondataapi.HandleFunc("/robencoval", getRobEncoVal).Methods("GET")
	visiondataapi.HandleFunc("/lowlimitval", getVisionLowLimitVal).Methods("GET")
	visiondataapi.HandleFunc("/uplimitval", getUpLimitVal).Methods("GET")*/

	axisctrlapi := robotapi.PathPrefix("/axisctrl").Subrouter()
	axisctrlapi.HandleFunc("/dynamicMode", setDynamicMode).Methods("PUT")
	axisctrlapi.HandleFunc("/slotio", testSlotIO).Methods("PUT").Queries("slotnum", "{slotnum}")
	axisctrlapi.HandleFunc("/encode", getEncode).Methods("GET")
	axisctrlapi.HandleFunc("/zeroencode", setZeroEncode).Methods("PUT").Queries("axis", "{axis}")
	axisctrlapi.HandleFunc("/servo", setServoStatus).Methods("PUT")
	axisctrlapi.HandleFunc("/axisinput", getAxisInput).Methods("GET").Queries("channel", "{channel}", "io_num", "{io_num}")
	axisctrlapi.HandleFunc("/axisonput", setAxisOnput).Methods("PUT").Queries("channel", "{channel}", "io_num", "{io_num}")
	axisctrlapi.HandleFunc("/dragteach", setDragteach).Methods("PUT")
	//axisctrlapi.HandleFunc("/encoder", getEncoder).Methods("GET")

	repositoryapi := robotapi.PathPrefix("/repository").Subrouter()
	repositoryapi.HandleFunc("/arcparam", setArcParam).Methods("PUT").Queries("md_id", "{md_id}", "file_no", "{file_no}", "index", "{index}")
	repositoryapi.HandleFunc("/interference", setInterference).Methods("PUT").Queries("md_id", "{md_id}", "no", "{no}", "index", "{index}")
	repositoryapi.HandleFunc("/param", setParam).Methods("PUT").Queries("md_id", "{md_id}", "index", "{index}")
	repositoryapi.HandleFunc("/toolframe", setToolFrame).Methods("PUT").Queries("md_id", "{md_id}", "index", "{index}", "index1", "{index1}", "index2", "{index2}")
	repositoryapi.HandleFunc("/userframe", setUserFrame).Methods("PUT").Queries("md_id", "{md_id}", "userNo", "{userNo}")
	repositoryapi.HandleFunc("/zeropoint", setZeroPoint).Methods("PUT").Queries("md_id", "{md_id}", "index", "{index}")

	return
}