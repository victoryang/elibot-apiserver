package shm

// #cgo CFLAGS: -I/root/mcserver/include/
// #cgo LDFLAGS: -lshare
// #include <stdlib.h>
// #include <shared.h>
import "C"
import (
	"bytes"
	"sync"
	"encoding/json"
	"hash/crc32"

	Log "elibot-apiserver/log"
)

type DisplayInfo struct {
	DisplayState		uint8 			`json:"displaystate,omitempty"`
	DisplayMsg			string			`json:"displaymsg,omitempty"`
	SerialNum			[]uint8 		`json:"serialnum,omitempty"`
	RunTime				uint32			`json:"runtime,omitempty"`
	TotalTime			uint32			`json:"totaltime,omitempty"`
}

type RobPLC struct {
	PLCIn			[]uint8
	PLCOut			[]uint8
	PLCVIn			[]uint8
	PLCVOut			[]uint8
	PLCM			[]uint8
}

type RobotSYSVAR struct {
	CRobB			[]int32
	IRobI			[]int16
	DRobD			[]float64
	DRobP			[][]float64
	DRobV			[][]float64
}

type RobotLOCVAR struct {
	CRobLB			[]uint32
	IRobLI			int32
	DRobLD			[]float64
	DRobLP			[][]float64
	DRobLV			[][]float64
}

type UserFrameData struct {
	IfTeached			uint8
	Useable				uint8
	Org					[]float64
	Xx					[]float64
	Xy					[]float64
	Buser				[]float64
	Note				string
}

type ToolData struct {
	IfTeached 			uint8
	Useable 			uint8
	ToolPos				[][]float64
	Btool				[]float64
	Note				string
}

type INTERFINFOR struct {
	IsUsed 			uint8
	UsedType		int32
	AxisNum			uint8
	OrigJoint		[]float64
	VertexJoint		[][]float64
	OrigPose		[]float64
	Dist 			[]float64
}

type WeldOffset struct {
	ManufacID		int
	FuncType		byte
	IPAddress 		[]byte
	ServerPort 		int
	IsUsed			int
	Mode 			int
	Accuracy		float64
	Distance		float64
	Freq			float64
	Schcnt			int
	Sta 			int
	Ex				float64
	Ez				float64
}

type ArcPara struct {
	Index 				int
	Note				[]byte
	ArcCurrent			float64
	ArcVoltage			float64
	StartCurrent		float64
	StartVoltage		float64
	StartTime			float64
	EndCurrent			float64
	EndVoltage			float64
	EndTime				float64
	AntiStickCurrent	float64
	AntistickVoltage	float64
	AntistickTime		float64
	AntiRoolback		int
}

type WeldPara struct {
	RestartDist					float64
	RestartSpeed				float64
	ArcCheckTime				float64
	ArcConfirmTime				float64
	ArcExhaustedCheckTime		float64
	ScrapintPaintDist			float64
	ScrapintPaintSpeed			float64
	PrepareAspirationTime		float64
	DelayAspirationTime			float64
	CheckCurrentInterruption	int
	CheckAlarmInterruption		int
	CheckWaterInterruption		int
	RestartUpAction				int
	ScrapintPaintAction			int
	AntiCollisionDetection		int
	DiandongTime				int
	OutputCurrent				[]float64
	CorrespondCurrent			[]float64
	OutputVoltage				[]float64
	CorrespondVoltage			[]float64
	RealCur2Volt				float64
	RealVolt2Volt				float64
	Arcpara						[]ArcPara
}

type DISTSENSERINFOR struct {
	IsUsed 				uint32
	ToolMeasure			uint32
	ToolWork 			uint32
	AnalogPort			uint32
	ValidDist			float64
	MinAnalog			float64
	MaxAnalog			float64
	MinDist				float64
	MaxDist				float64
}

type TrackParameter struct {
	
}

type SharedResource struct {
	DspInfor			DisplayInfo
	AutoRunCycleMode	uint8
	DisableAxis			uint32
	NopLine				uint32
	ProgEnd				uint8
	TeachMode			uint8
	MachinePos			[]float64
	MachinePose			[]float64
	AbsPulse			[]int32
	AbzPulse			[]int32
	CurEncode 			[]int32
	RobotState			int32
	RobotPause			int32
	RobotStop			int32
	ServoReady			int32
	CanMotorRun			int32
	CurrentLine			int32
	AutoRunLineNum		int32
	ZeroOffset			[]float64
	MotorSpeed			int32
	Plc					RobPLC
	Sysvar 				RobotSYSVAR
	RobotMode			int32
	AutoRunToolNum		int32
	Locvar				[]RobotLOCVAR
	LocvarNum			int32
	CommState			int32
	ArcEnable			int32
	UserFrame			[]UserFrameData
	ToolFrame			[]ToolData
	InterfInfor			[]INTERFINFOR
	Weldoffset			WeldOffset
	StruArcWeld			WeldPara
	ForceClAlarm		int
	AxisIoInput			[]int
	AnalogIoInput		[]float64
	PositiveVar			[]float64
	DistSenser			DISTSENSERINFOR
	Trackpara			TrackParameter
}

var SharedResourcePool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, BufferSize))
	},
}
var crc_shared_resource int = 0

func getResourceAndCompare() (res []byte){
	mutex.Lock()
	defer mutex.Unlock()

	resource := SharedResource{
		Test:		1,
	}

	var crc int
	now, err := json.Marshal(resource)
	if err!=nil {
		crc = 0
	} else {
		crc = int(crc32.ChecksumIEEE(now))
	}

	var cache []byte
	buf := SharedResourcePool.Get().(*bytes.Buffer)
	if buf == nil {
		buf = bytes.NewBuffer(make([]byte, 0, BufferSize))
		crc_shared_resource = 0
	} else {
		cache = buf.Bytes()
	}

	if crc == 0 {
		crc_shared_resource = 0
 		res = []byte("")
 	} else {
 		if crc == crc_shared_resource {
 			res = nil
 			return
 		} else {
 			crc_shared_resource = crc
 			res = now
 		}
 	}

	buf.Reset()
	buf.Write(res)
	defer func(c []byte) {
		if e := recover(); e!=nil {
			Log.Error("buf write error: ", e)

			buf := make([]byte, 0, BufferSize*2)
			buf = c[:]
			SharedResourcePool = sync.Pool{
				New: func() interface{} {
					return bytes.NewBuffer(buf)
				},
			}
		}
	}(cache)

	SharedResourcePool.Put(buf)
	return
}

func watchSharedResource(modified chan []byte) {
	if res := getResourceAndCompare(); res != nil {
		modified <- res
	}
}

func InitSharedResource() int {
	return int(C.init_shared_resource())
}