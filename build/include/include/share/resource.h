#ifndef __SHARE_RESOURCE_H__
#define __SHARE_RESOURCE_H__

#include <stdint.h>
#include "../../config.h"
#include "para_define.h"
#include "arcwelding.h"
#include "ext_device_def.h"
#include <robot/trackcraft.h>
#include <robot/visioncraft.h>
#include <robot/ExtAxis.h>
#include "ext_device/chuangxiangzhikong.h"
#include <robot/dynamic_param.h>
#include <sys/mman.h>
#include <share/weaving.h>
#include <math.h>
#include <iobus.h>

//#include <share/fishweld.h>
#ifdef __cplusplus
extern "C" {
#endif 

/* 坐标系*/
enum Robot_Coordinate{
	ROBOT_COORDINAT_JOINT=0,					  /* 关节坐标系*/
	ROBOT_COORDINAT_CART=1,					   /* 直角坐标	   */
	ROBOT_COORDINAT_TOOL=2,					   /*  工具坐标 	*/
	ROBOT_COORDINAT_USER=3,					   /*  用户坐标系*/
	ROBOT_COORDINAT_CYLINDER=4,					 /* 圆柱坐标	 */
	ROBOT_COORDINAT_MAXV
};

enum Robot_State{
	ROBOT_STATE_STOP=0,
	ROBOT_STATE_PAUSE=1,
	ROBOT_STATE_EMESTOP=2,
	ROBOT_STATE_RUNNING=3,
	ROBOT_STATE_ERROR=4
};

enum Robot_Mode{
	ROBOT_MODE_TEACH=0,	//示教模式
	ROBOT_MODE_PLAY=1,	//再现模式
	ROBOT_MODE_REMOTION=2	//远程模式
};

enum Robot_Cycle_Mode{
	CYCLE_MODE_STEP=0,   //单步
	CYCLE_MODE_ONE=1,   //one 循环
	CYCLE_MODE_SERIES=2   //连续循环
};

enum Robot_Motor_type{
	Motor_NoUsed,
	Motor_Step,
	Motor_servo
};
//option by bit
enum Sys_Technics{arc_sys=0,transfer_sys=1,common_sys=2,stamp_sys=3 };
#define getTechnics()	(NVRAM_DATA(sys_technics)>>28&0xF)

typedef enum {
	modual_extAxis=0,//外部轴
	modual_bookProg=1,//预约
	modual_vision=2,//视觉
	modual_trace=3,//跟踪
	modual_laserTrace=4,//激光跟踪
	modual_laserRanging=5,//激光测距
	modual_remote=6,
	modual_robotiq=7,//robotiq抓手
	modual_6DSensor=8,	//６维传感
}MODUAL_FUNCTION;
#define getModualFunc()	(NVRAM_DATA(sys_technics)&0x0FFFFFFF)
#define ModualFuncEnable(m)	(getModualFunc() & (0x01 << m))

enum Sys_Ctrl_Mode{c_ascii,c_raw,c_ethercat,c_rbtfpga};

//enum OperationMode {DynOpMode_Normal, DynOpMode_DirTch,DynOpMode_FeedForward};

//!文件名长度
#define FILE_NAME_LENGTH			32
#define USER_FRAME_COUNT	8
#define TOOL_FRAME_COUNT	8

#define GLOBAL_VAR_COUNT	256
#define LOCAL_VAR_COUNT		128
#define GLOBAL_PVAR_COUNT	256

#define B_COUNT		GLOBAL_VAR_COUNT
#define LB_COUNT	LOCAL_VAR_COUNT

#define I_COUNT		8256
#define LI_COUNT	LOCAL_VAR_COUNT

#define D_COUNT		GLOBAL_VAR_COUNT
#define LD_COUNT	LOCAL_VAR_COUNT

#define V_COUNT		GLOBAL_VAR_COUNT
#define LV_COUNT	LOCAL_VAR_COUNT

#define P_COUNT		GLOBAL_PVAR_COUNT
#define LP_COUNT	LOCAL_VAR_COUNT


#define VSub_COUNT	AXIS_COUNT
#define BOOK_FILE_NUM		8

#define C_COUNT 	10240
#define AXIS_IN_NUM 4
#define ANALOG_IN_NUM 2
//by zxc///////////////关于IO的定义//////////////////////////
#define IO_IN_NUM	16
#define IO_OUT_NUM	16
#define IO_VIN_NUM	16
#define IO_VOUT_NUM	16
#define IO_M_TMR	259	//260	//定时器用到的号

#define IO_M_NUM	192
#define IO_S_NUM	125

#define IO_M_NUM_BIT	(IO_M_NUM*8)
#define IO_M_NUM_BIT4	(IO_M_NUM*2)
#define IO_M_NUM_BYTE	(IO_M_NUM)

#define IO_IN_NUM_BIT	(IO_IN_NUM*8)
#define IO_IN_NUM_BIT4	(IO_IN_NUM*2)
#define IO_IN_NUM_BYTE	(IO_IN_NUM)
#define IO_OUT_NUM_BIT	(IO_OUT_NUM*8)
#define IO_OUT_NUM_BIT4	(IO_OUT_NUM*2)
#define IO_OUT_NUM_BYTE	(IO_OUT_NUM)

#define SCAN_PERIOD	5

extern struct weav_para *weavData;
extern unsigned char needdynamic;
//#define NEED_DYNAMIC	(needdynamic)

typedef enum {
	st_init=0,st_error=1,st_encode_success=2,st_cal_tool_success=3,st_cal_user_success=4,st_set_extaxis_success=5,
	st_bookprog_success=6,st_calc_zero_offset_success=7,st_save_lasertrack_success=8,st_test_lasertrack_success=9,
	st_save_distSenser_success=10,st_save_trackdata_success=11,st_save_weavingdata_success=12,st_display_meg=13,
	st_save_arcwelddata_success = 14,st_save_visiondata_success=15,
}SYSCOMM_STATE;

typedef struct{
	unsigned char enable;
	unsigned char num;
	char book_prog_name[BOOK_FILE_NUM][FILE_NAME_LENGTH];
}StructBookFile;

typedef struct user_frame_data{
	unsigned char  if_teached;
	unsigned char useable;
	double org[AXIS_COUNT];
	double xx[AXIS_COUNT];
	double xy[AXIS_COUNT];
	double buser[6];
	char note[32];
}USER_FRAME_DATA;

typedef struct TOOL_DATA{
	unsigned char if_teached;
	unsigned char useable;
	double toolPos[7][AXIS_COUNT];
	double btool[6];
	char note[32];

    // 工具质量相关
    double quality;                 // 质量, 单位 kg
    double centre_of_mass[3];       // 质心,[x,y,z]
}TOOLDATA;

#define INTERF_NUM	16
enum interf_type{error,vertex_2,vertex_axis,vertex_1};
typedef struct {
	unsigned char is_used;
	enum interf_type used_type;//error:0, vertex and teach 2 point:1, axis:2,vertex and teach 1 point:3
	unsigned char axis_num;
	double orig_joint[AXIS_COUNT];
	double vertex_joint[2][AXIS_COUNT];
	double orig_pose[6];
	double dist[3]; //1/2 distense for interference (vertex to orig)
}INTERF_INFOR;

/* 冲压工艺　*/
#define STAMP_D_COUNT 100
#define STAMP_P_COUNT 50

typedef struct {
	unsigned char isenable;
	unsigned char islast;
}UDP_CONNECT_INFIOR;

typedef struct tech_stamp_paras_s {
	/*　存储整数　*/
	int iRobI[STAMP_D_COUNT];
	/*　存储关节角度　*/
	double dRobP[STAMP_P_COUNT][AXIS_COUNT];
} tech_stamp_paras_t;


typedef struct
{
	volatile uint8_t  PLC_IN[IO_IN_NUM];
	volatile uint8_t  PLC_OUT[IO_OUT_NUM];
	volatile uint8_t  PLC_VIN[IO_VIN_NUM];
	volatile uint8_t  PLC_VOUT[IO_VOUT_NUM];
	volatile uint8_t PLC_M[IO_M_NUM];
	volatile uint8_t PLC_S[IO_S_NUM];
}Rob_PLC;

//局部变量结构定义
#define CALL_NEST_NUM 10//子程序的层数,即局部变量的组数,

typedef struct localvar
{
	unsigned int cRobLB[LB_COUNT];		//字节型变量
	int iRobLI[LI_COUNT];		//整数型变量
	double dRobLD[LD_COUNT];	//双精度型变量
	double dRobLP[LP_COUNT][AXIS_COUNT];	//机器人轴位置
	double dRobLV[LV_COUNT][VSub_COUNT];
}Robot_LOCVAR;

typedef struct systemvar
{
	unsigned int cRobB[B_COUNT]; //字节型变量
	short iRobI[I_COUNT]; 				//整数型变量
	double dRobD[D_COUNT];	   //双精度型变量
	double dRobP[P_COUNT][AXIS_COUNT];
	double dRobV[V_COUNT][VSub_COUNT];
}Robot_SYSVAR;

struct DISPLAY_INFOR{
	unsigned char display_state;
	char display_msg[1024];
};

struct ENCRYPT_INFOR {
#define ENCRYPTSTATE_INVALID   -1 // 无效版本（试用期截至或序列号无效）此时不允许机器人自动运行
#define ENCRYPTSTATE_UNENCRYPT  0 // 未加密版本
#define ENCRYPTSTATE_LEGAL      1 // 解密版本
#define ENCRYPTSTATE_TRIAL      2 // 试用版本（试用期内版本）
	int           encrypt_state;

	unsigned char mach_code[33]; // 机器码
	int           remain_time;   // 剩余试用期
};

#define ROB_SHARE_RES_MAGIC		0x53485253	//SHRS
struct para_info{
    uint32_t magic;
    uint32_t crc32;
    struct rbctrl_para data;
};

typedef struct{
	uint32_t magic;
	struct DISPLAY_INFOR dspInfor;
	struct ENCRYPT_INFOR encryptInfo;
	unsigned char autorun_cycleMode;
	unsigned int disable_axis;
	unsigned int nop_line;
	unsigned char prog_end;
	unsigned char teach_mode;
	double machinePos[AXIS_COUNT];	//机器人位置 ,关节角度
	double machinePose[6];//机器人(x,y,z,Rx,Ry,Rz)
	volatile int32_t abs_pulse[AXIS_COUNT];
	volatile int32_t abz_pulse[AXIS_COUNT];
	volatile int32_t cur_encode[AXIS_COUNT];
	volatile int32_t robotState;	//当前运行状态
	volatile int32_t robotPause;	//外部暂停状态 
	volatile int32_t robotStop;	//外部停止状态
	volatile int32_t servoReady;
	volatile int32_t can_motor_run;
	volatile int32_t currentLine; //current run line No.
	volatile int32_t autorun_lineNum; //
	double zero_offset[9];
	int motor_speed[AXIS_COUNT];
	Rob_PLC plc;	//机器人的PLC
	Robot_SYSVAR sysvar;
	volatile int32_t robotMode; //机器人控制模式
	volatile int32_t autorun_toolNum; //
	Robot_LOCVAR locvar[CALL_NEST_NUM];	//局部变量数组定义 
	volatile int locvar_num;//当前程序段所使用的局部变量组中的序号
	SYSCOMM_STATE comm_state;
	volatile int32_t arc_enable;
	USER_FRAME_DATA user_frame[USER_FRAME_COUNT];
	TOOLDATA tool_frame[TOOL_FRAME_COUNT];
	INTERF_INFOR interf_infor[INTERF_NUM];
	struct weld_para stru_arcWeld;
	int forceClAlarm;
	int axis_ioInput[AXIS_IN_NUM];//set by bit
	double analog_ioInput[ANALOG_IN_NUM];
	double positive_var[AXIS_COUNT];
	DIST_SENSER_INFOR dist_senser;
	VisionData vision_data;
	TrackData track_data;
	UDP_CONNECT_INFIOR udp_infor;
	StructBookFile bookfile_infor;
	ExtAxisData extAxis_Data;
	WeavingPara weave_data;
	LaserTrackData lasertrack_data;
	//TIGdata tig_data;
	tech_stamp_paras_t stamp_paras;
	volatile int32_t servo_dirve_mode;	// 私服驱动模式
    struct para_info sys_parameters;
    inverse_dynamic_param idy_param;  // 动力学参数
    io_bus iobus;
	uint8_t remoteAccessEnabled; 		// 远程访问启用状态
}robot_share_resouce;

extern robot_share_resouce *share_resource;
/**
 * 获取机器人共享资源句柄
 * @return 共享资源句柄
 */
extern robot_share_resouce* get_rob_share_resource(void);

#define SHARE_RES(x)	(get_rob_share_resource())->x

extern int resource_create(const char *key_file);
extern int resource_init(const char *key_file);
extern void resource_release(void);


#include "para_define.h"
#include "hash_define.h"

#define PARA_DATA_MAGIC	HASH4('P', 'A', 'R', 'A')

struct nv_data_s{
	uint32_t magic;
	char projectName[FILE_NAME_LENGTH];
	volatile int32_t sys_technics;
	volatile int32_t curCoordinate; //当前坐标系
	volatile int32_t manualSpeedRate;
	volatile int32_t cycleMode; //机器人循环状态, CYCLE_MODE_...
	volatile int32_t toolNum; //current of tool number
	volatile int32_t userNum; //current of tool number
	volatile int32_t trackNum;  //当前跟踪工艺号
	volatile int32_t zero_encode[AXIS_COUNT]; //robot's zero pos ,encode data
	volatile double system_period;
	volatile enum Sys_Ctrl_Mode system_ctrl_mode;
	volatile double origin[AXIS_COUNT];
};

#define NV_DATA_MAGIC	HASH4_S('N', 'V', 'R', 'A', sizeof(struct nv_data_s))

struct nvram_data_s{
//	struct para_info para;
	struct nv_data_s data;
};


/**
 * 获取机器人NVRAM共享资源句柄
 * @return NVRAM共享资源句柄
 */
extern struct nvram_data_s* get_rob_nvaram_data(void);

#define NVRAM_SIZE		(sizeof(struct nvram_data_s))
#define NVRAM_MAX_SIZE		(16*1024)
int init_nvram(const char *nvram);
void release_nvram(void);
extern struct nvram_data_s *nvram_data;

#define NVRAM_DATA(x)	(get_rob_nvaram_data())->data.x
//#define NVRAM_PARA(x)nvram_data->para.data.x
#define NVRAM_PARA(x)	(SHARE_RES(sys_parameters.data)).x

#define NVRAM_PARA_DATA_POINTER() &((SHARE_RES(sys_parameters)).data)
#define NVRAM_PARA_POINTER() &(SHARE_RES(sys_parameters))

#define msync_NVRAM()	msync(nvram_data, NVRAM_SIZE, MS_SYNC)
#define RobotRes_MachPos_base	SHARE_RES(machinePos)
#define RobotRes_MachinePos(n)	RobotRes_MachPos_base[n]

#define RobotRes_MachPose_base	SHARE_RES(machinePose)
#define RobotRes_MachinePose(n)	RobotRes_MachPose_base[n]
#define Get_RobotRes_Origin(n)		NVRAM_DATA(origin)[n]
#define Set_RobotRes_Origin(n,v)	do{NVRAM_DATA(origin)[n]=v;msync_NVRAM();}while(0)

#define PositiveVar_base	SHARE_RES(positive_var)
#define PositiveVar(n)		PositiveVar_base[n]

#define UDP_infor_data(n) 	SHARE_RES(udp_infor.n)

#define SetCurLine(line)		SHARE_RES(currentLine) = line
#define GetCurLine()			SHARE_RES(currentLine)

#define SetAutoRunLine(line)		SHARE_RES(autorun_lineNum)=line
#define GetAutoRunLine()			SHARE_RES(autorun_lineNum)

#define SetRobotPause(isPause)	SHARE_RES(robotPause)=isPause
#define GetRobotPause()			SHARE_RES(robotPause)

#define SetRobotStop(isStop) 		SHARE_RES(robotStop)=isStop
#define GetRobotStop()			SHARE_RES(robotStop)

#define GetRemoteAccessState() SHARE_RES(remoteAccessEnabled)

#define GETMOTOR_RUN_STATE()	SHARE_RES(can_motor_run)
#define SETMOTOR_RUN_STATE(v)	do{if(SHARE_RES(can_motor_run) != v) SHARE_RES(can_motor_run) = v; SET_VOUT_BIT(CAN_MOTOR_RUN_VOUT, v);}while(0)

#define SetRobotState(state) 		SHARE_RES(robotState)=state
#define GetRobotState()			SHARE_RES(robotState)
#define SetRobotMode(mode)		SHARE_RES(robotMode)=mode
#define GetRobotMode()			SHARE_RES(robotMode)
#define SetCycleMode(mode)		do{NVRAM_DATA(cycleMode)=mode;msync_NVRAM();}while(0)
#define GetCycleMode()			NVRAM_DATA(cycleMode)

#define SetCurCoordinate(coord)	do{NVRAM_DATA(curCoordinate)=coord;msync_NVRAM();}while(0)
#define GetCurCoordinate()		NVRAM_DATA(curCoordinate)

#define SetDisable_Axis(n,v) 		SHARE_RES(disable_axis)=((SHARE_RES(disable_axis)&(~(0x01<<n)))|(v<<n))
#define GetDisableAxis(n) 		(SHARE_RES(disable_axis) & (0x01<<n))
#define ResetDisableAxis() 		(SHARE_RES(disable_axis)=0)

#define SetServoReady(flg)		SHARE_RES(servoReady)=flg
#define GetServoReady()			SHARE_RES(servoReady)

#define GetMaxLimt()			   NVRAM_PARA(limit_max_pos)
#define GetMinLimt()			   NVRAM_PARA(limit_min_pos)

#define GetMainfile()				NVRAM_DATA(projectName)
#define SetMainfile(fname)		do{sprintf( NVRAM_DATA(projectName),"%s",fname);SetAutoRunLine(GetCurLine());}while(0)

#define RBTEACH_MAX_BACKLIGHT		100
#define GetLCDBackLight()			(NVRAM_PARA(lcd_backlight[0])<RBTEACH_MAX_BACKLIGHT ? NVRAM_PARA(lcd_backlight[0]):RBTEACH_MAX_BACKLIGHT)
#define SetLCDBackLight(light)		(NVRAM_PARA(lcd_backlight[0])=(light<RBTEACH_MAX_BACKLIGHT?light:RBTEACH_MAX_BACKLIGHT))

#define GetHalfLCDBackLight()			(NVRAM_PARA(lcd_backlight[1])<RBTEACH_MAX_BACKLIGHT ? NVRAM_PARA(lcd_backlight[1]) :RBTEACH_MAX_BACKLIGHT)
#define SetHalfLCDBackLight(light)		(NVRAM_PARA(lcd_backlight[1])=(light<RBTEACH_MAX_BACKLIGHT?light:RBTEACH_MAX_BACKLIGHT))

#define GetLCDHalfCloseTime()				NVRAM_PARA(lcd_time[0])
#define SetLCDHalfCloseTime(time)			NVRAM_PARA(lcd_time[0])=time

#define GetLCDCloseTime()				NVRAM_PARA(lcd_time[1])
#define SetLCDCloseTime(time)			NVRAM_PARA(lcd_time[1])=time

#define GetRobotType()	NVRAM_PARA(robot_type) 
#define SetRobotType(type)	NVRAM_PARA(robot_type)=type
#define getUserLevelPara()	NVRAM_PARA(user_level)
#define GetToolNumber()	NVRAM_DATA(toolNum)
#define SetToolNumber(v)	do{NVRAM_DATA(toolNum)=v;msync_NVRAM();}while(0)
#define ToolData(n)		SHARE_RES(tool_frame[n])

#define Get_AutoRun_ToolNumber()	SHARE_RES(autorun_toolNum)
#define Set_AutoRun_ToolNumber(v)	(SHARE_RES(autorun_toolNum)=v)

#define GetUserNumber()	NVRAM_DATA(userNum)
#define SetUserNumber(v)	do{NVRAM_DATA(userNum)=v;msync_NVRAM();}while(0)
#define userFrameData(n)		SHARE_RES(user_frame[n])
#define axisIOInput(n)	SHARE_RES(axis_ioInput[n])
#define analogIOInput(n)	SHARE_RES(analog_ioInput[n])
#define interf_struct		SHARE_RES(interf_infor);
#define interf_isused(n)		SHARE_RES(interf_infor[n].is_used)

#define getSysVar_BID(var,num)	(SHARE_RES(var)[num])
#define setSysVar_BID(var,num,v)	(SHARE_RES(var)[num]=v)

#define getSysVar_P(num,sub)	SHARE_RES(sysvar.dRobP)[num][sub]
#define setSysVar_P(num,sub,v)	SHARE_RES(sysvar.dRobP)[num][sub]=v

#define getSysVar_V(num,sub)	SHARE_RES(sysvar.dRobV)[num][sub]
#define setSysVar_V(num,sub,v)	SHARE_RES(sysvar.dRobV)[num][sub]=v

#define getLocalVar_P(l,num,sub)	SHARE_RES(locvar[l]).dRobLP[num][sub]
#define setLocalVar_P(l,num,sub,v)	SHARE_RES(locvar[l]).dRobLP[num][sub]=v

#define GetWaitTime()    NVRAM_PARA(catch_wait_time)

#define getMotorSpeed(num)		SHARE_RES(motor_speed)[num]
#define setMotorSpeed(num,v)	getMotorSpeed(num)=v

/* 冲压工艺参数宏方法 */
#define getStampPara_I(num) 			(SHARE_RES(stamp_paras).iRobI[num])
#define setStampPara_I(num,v)			(SHARE_RES(stamp_paras).iRobI[num]=v)
#define getStampPara_P(num,sub)			(SHARE_RES(stamp_paras).dRobP[num][sub])
#define setStampPara_P(num,sub,v)		(SHARE_RES(stamp_paras).dRobP[num][sub]=v)
#define getStampPara_Ptr()				(&SHARE_RES(stamp_paras))

#define PRESS_RESET	422
#define plc_press_reset(v)	SET_VOUT_BIT(PRESS_RESET,v)

//return 1:single line ,0:double line 
#define GetLineNumer()	NVRAM_PARA(line_number)

#define BOOK_PROG_INFOR		SHARE_RES(bookfile_infor)
#define book_prog_num	(BOOK_PROG_INFOR.num)
#define book_prog_file(n)		(BOOK_PROG_INFOR.book_prog_name[n])
#define book_prog_enable	(BOOK_PROG_INFOR.enable)

#define ZERO_OFFSET(n)		SHARE_RES(zero_offset)[n]

#define BOUND(i, v, x)	((v)>(x)?(x):(v)<(i)?(i):(v))
#define GetManualSpeedRate()	BOUND(1,NVRAM_DATA(manualSpeedRate),10000)
#define SetManualSpeedRate(v)	do{NVRAM_DATA(manualSpeedRate)=(v>10000?10000:v);msync_NVRAM();}while(0)

#define EXTAXIS_DATA()		SHARE_RES(extAxis_Data)

#define WEAVING_DATA()		SHARE_RES(weave_data)
#define WEAVING_INFOR(x) (WEAVING_DATA().weave_infor[x])
#define SET_WEAVING_ENABLE(b)	(WEAVING_DATA().enable=b)
#define GET_WEAVING_ENABLE()	(WEAVING_DATA().enable)
#define SET_WEAVING_NUM(n)	(WEAVING_DATA().tech_num=n)
#define SET_WEAVING_REFP(p,n)	memcpy(WEAVING_DATA().refp[n],p,sizeof(double)*AXIS_COUNT)
#define RESET_WEAVING_REFP(n)	memset(WEAVING_DATA().refp[n],0,sizeof(double)*AXIS_COUNT)

#define GetTrackNumber()	NVRAM_DATA(trackNum)
#define SetTrackNumber(v)	do{NVRAM_DATA(trackNum)=v;msync_NVRAM();}while(0)
#define TRACK_DATA()     SHARE_RES(track_data)
#define TRACKCRAFT_DATA(x)  (TRACK_DATA().TrackCraft_Data[x])

#define VISION_DATA()     SHARE_RES(vision_data)
#define VISIONCRAFT_DATA(x)  (VISION_DATA().VisionCraft_Data[x])

#define GET_TEACH_MODE()	SHARE_RES(teach_mode)
#define SET_TEACH_MODE(v)	GET_TEACH_MODE()=v
#define SET_AUTORUN_CYCLEMODE(v)	SHARE_RES(autorun_cycleMode) = v
#define GET_AUTORUN_CYCLEMODE()	SHARE_RES(autorun_cycleMode)

#define INVERSE_DYNAMIC_PARAM(x) SHARE_RES(idy_param).x

const char* errlog_file_name(void);

//for robot_dh_infor
extern void CalcUrInfor(void);
extern void CalcMaduoInfor(void);
extern void CalcGeneralInfor(void);

#define GetGeneralDHInfor(x)	NVRAM_PARA(dh_stru[x])
#define SetGeneralDHInfor(x,v)	(NVRAM_PARA(dh_stru[x])=v)

#define Get_ZeroEncode(a)	(NVRAM_DATA(zero_encode)[a])
#define Set_ZeroEncode(a,v)	do{(NVRAM_DATA(zero_encode)[a])=v;msync_NVRAM();}while(0)

#define Get_CurEncode(a)	(SHARE_RES(cur_encode)[a])
#define Set_CurEncode(a,v)	(SHARE_RES(cur_encode)[a]=v)

#define setCommState(d)	SHARE_RES(comm_state)=d
#define getCommState()	SHARE_RES(comm_state)

#define STRU_ARCWELD_POINT()	&(SHARE_RES(stru_arcWeld))
#define STRU_ARCWELD_DATA(d)	(SHARE_RES(stru_arcWeld).d)
#define setArcData(x,d)	SHARE_RES(stru_arcWeld.x)=d
#define getArcData(x)		SHARE_RES(stru_arcWeld.x)

#define setArcPara(i,x,d)	SHARE_RES(stru_arcWeld.arcPara[i].x)=d
#define getArcPara(i,x)	SHARE_RES(stru_arcWeld.arcPara[i].x)

#define setArcEnable(d)	SHARE_RES(arc_enable)=d
#define isArcEnable()	SHARE_RES(arc_enable)
#define setToolEnable(d) setArcEnable(d)
#define isToolEnable()	isArcEnable()

#define PROG_RUN_END	SHARE_RES(prog_end)
#define PROG_NOP_LINE()	SHARE_RES(nop_line)

#define GetSysPeriod()	NVRAM_DATA(system_period)
#define SetSysPeriod(p)	do{if(fabs(GetSysPeriod()-p)>em){NVRAM_DATA(system_period)=p;msync_NVRAM();}}while(0)

#define GetSysCtrlMode()	NVRAM_DATA(system_ctrl_mode)
#define SetSysCtrlMode(x)	do{if(x!=NVRAM_DATA(system_ctrl_mode)){NVRAM_DATA(system_ctrl_mode)=x;msync_NVRAM();}}while(0)

#define GetDisplayMsg()	SHARE_RES(dspInfor).display_msg
#define SetDisplayMsg(msg)	snprintf(SHARE_RES(dspInfor).display_msg,1024,"%s",msg)

#define GetDisplayMsg_State()	SHARE_RES(dspInfor).display_state
#define SetDisplayMsg_State()	GetDisplayMsg_State()=1
#define ReSetDisplayMsg_State()	GetDisplayMsg_State()=0

// 加密信息共享内存设置
//////////////////////////////////////////////////////////////////////////
// 获取与设置机器码
#define GetEncryptMachCode()	SHARE_RES(encryptInfo).mach_code
#define SetEncryptMachCode(msg) memcpy(SHARE_RES(encryptInfo).mach_code, msg, 33)

// 获取与设置剩余试用时间
#define GetEncryptRemainTime()	  SHARE_RES(encryptInfo).remain_time
#define SetEncryptRemainTime(rtm) SHARE_RES(encryptInfo).remain_time = rtm

// 获取与设置系统加密状态
#define GetEncryptState()	    SHARE_RES(encryptInfo).encrypt_state
#define SetEncryptState(estate) SHARE_RES(encryptInfo).encrypt_state = estate
//////////////////////////////////////////////////////////////////////////

// 伺服驱控模式
typedef enum {
	SDM_ONE_AXISX = 0, 		// 1 轴
	SDM_TWO_AXISX = 1, 		// 2 轴
	SDM_THREE_AXISX = 2, 	// 3 轴
	SDM_FOUR_AXISX = 3, 	// 4 轴
	SDM_FIVE_AXISX = 4, 	// 5 轴
	SDM_SIX_AXISX = 5, 		// 6 轴
	SDM_SEVEN_AXISX = 6, 	// 7 轴
	SDM_EIGHT_AXISX = 7, 	// 8 轴
} ServoDriveModeOption;

typedef enum {
	SDM_Ethercat = 1,			// 总线式
	SDM_Pulse = 2,				// 脉冲式
	SDM_Collaboration = 3,		// 协作
}ServoDriveMode;

// 获取伺服驱控方式
#define GetServoDriveMode() ((SHARE_RES(servo_dirve_mode) >> 28) & 0xF)
// 获取伺服驱动选项
#define GetServoDirveOption() (SHARE_RES(servo_dirve_mode) & 0x0FFFFFFF)
// 指定的选项是否启用
#define ServoDriveOptionEnable(option) (GetServoDirveOption() & (0x00000001 << option))
// 清空伺服驱控选项
#define ClearServoDirveOption()	do { SHARE_RES(servo_dirve_mode) = SHARE_RES(servo_dirve_mode)  & 0xF0000000 } while(0)
// 设置伺服驱控模式
#define SetServoDriveMode(mode) do { SHARE_RES(servo_dirve_mode) = mode; } while(0)

#include <remote.h>

#ifdef __cplusplus
}
#endif

#endif

