/********************************************************************
* Authors@threewater
* All Right Reserved
*********************************************************************/

#ifndef __ROBOTRESOURCE_H__
#define __ROBOTRESOURCE_H__

#include <stdint.h>
#include <share/resource.h>
#include <robot/axisctrl.h>
#include <src/robot/robotram.h>

typedef struct stitchData{
	unsigned char isused;
	unsigned char type;
	double v1;
	double v2;
}STRU_STITCH;

extern STRU_STITCH stitch_data;

#define get_stitch_data(n)	(stitch_data.n)
#define set_stitch_data(n,v)	(stitch_data.n=v)

void robotresource_init(void);
int open_libencryption(void);

void GetMachinePos(double machinePos[]);
void GetMachinePose(double machinePose[]);

void SetMachinePos(double machinePos[]);
int MachinePosCmp(double machinePos[]);
void SetMachinePosChip(sendpos_t *sndpos_data);
void sync_abs_pulse(void);
void exec_synchr(int pulse[], int opt);

void GetRobotOverMeasure(double overMeasure[]);
void SetRobotOverMeasure(double overMeasure[]);
void ClearRobotOverMeasure(void);

double GetRobotRate(unsigned int index);
void SetRobotRate(unsigned int index,double value);
int str2pos(const char *p,double pos[AXIS_COUNT]);
void userdata_init(void);
int get_buser(double p[6],int num);
void tooldata_init(void);
int get_btool(double p[6],int num);
void posTostring(char *str,double p[],int n);
void pulseget_ouhe_process(double *in, double *out);
void free_prog_list(void);
typedef int (*calclate_user_func)(double org_joint[],double xx_joint[],double xy_joint[],double UserData[]);
extern calclate_user_func user_CoordinateBuild;
int movj2specifiedPoint(double point[],unsigned char inqueue);
int exec_arcweld(int ison);
void interference_init(void);
void is_in_interference(double *pose);
int check_robot_axis(void);
int get_nvram_file(void);
int plc_thread_init(void);

#define pi		3.1415926535898 					 //定义的π
#define radianPerAngle	(pi/180)
#define anglePerRadian (180/pi)

//!Jbi文件备注长度
#define JBI_FILE_COMMENT_LENGHT			64
#define PERIOD	GetSysPeriod()						   //插补周期

/*POSE	单位系数*/
#define em 1e-6
#define threedecimal_em 1e-3

typedef double RobotPos[AXIS_COUNT];

//for offline path
#define offline_data_len		15
#define OFFLINE_SUB_COUNT 	5120
typedef double OFFLINE_POS[offline_data_len];
OFFLINE_POS* get_offline_subpath(int subp,int *num);
/***************************************************************/

#include "common.h"
/****************************************************IO define******************************************************************************/
//M0-M399(vin0-vin399) is for var input
//M400-M799(vout0-vout399) is for var output
//M800-M863 for physical out put,user IO,option by DOUT code.
#define GET_VIN_BIT(n)		GET_ARRAY_BIT(SHARE_RES(plc.PLC_M),n)
#define SET_VIN_BIT(n,v)	SET_ARRAY_BIT(SHARE_RES(plc.PLC_M),n,v)
#define GET_IN_BIT(n)		GET_ARRAY_BIT(SHARE_RES(plc.PLC_IN),n)
#define SET_IN_BIT(n,v)		SET_ARRAY_BIT(SHARE_RES(plc.PLC_IN),n,v)

#define GET_VOUT_BIT(n)	GET_ARRAY_BIT(SHARE_RES(plc.PLC_M),n)
#define SET_VOUT_BIT(n,v)	SET_ARRAY_BIT(SHARE_RES(plc.PLC_M),n,v)
#define GET_OUT_BIT(n)		GET_ARRAY_BIT(SHARE_RES(plc.PLC_OUT),n)
#define SET_OUT_BIT(n,v)	SET_ARRAY_BIT(SHARE_RES(plc.PLC_OUT),n,v)

//physical out from 0-16
#define BREAK_IO_OUT		18 //servo break
#define USER_IO_VOUT		800 //800-863 for DOUT can set
/************************************************************/
//phyical in 0-31
#define EMERGENCY_STOP_IN	0
/************************************************************/

#define BREAK_VIO_OUT	400
#define USER_ORIGIN_VIO_OUT	401
#define DASH_CLEAR_VIO_OUT	402
#define BOOKPROG_ENABLE_VIO_OUT	406

#define TEACHMODE_VIO_OUT	416
#define AUTO_VIO_OUT	417
#define REMOTE_VIO_OUT	418


#define STOP_VIO_OUT		424
#define PAUSE_VIO_OUT		425
#define EMESTOP_VIO_OUT	426
#define RUNNING_VIO_OUT	427
#define ERROR_VIO_OUT	428

//for ext machine option
#define EMERGENCY_STOP_VIO_IN	0
#define EXT_EMERGENCY_STOP_VIO_IN	1
#define DASH_ALARM_VIO_IN	2
#define EXT_START_VIO_IN	3
#define EXT_PAUSE_VIO_IN	4
#define EXT_SERVO_ON_VIO_IN	5
#define EXT_CLEAR_ALARM_IN	6
#define EXT_SAFE_DOOR_IN		7
#define AXIS_BREAKE_BASE	24
#define get_axis_break_IO(n)	GET_IN_BIT(n+AXIS_BREAKE_BASE)

//for work stack  book prog
#define STACK0_VIO_IN	8
#define STACK1_VIO_IN	9
#define STACK2_VIO_IN	10
#define STACK3_VIO_IN	11
#define STACK4_VIO_IN	12
#define STACK5_VIO_IN	13
#define STACK6_VIO_IN	14
#define STACK7_VIO_IN	15
#define HOMEPOS_AXIS_VIO_IN	16 //16+8 axis

#define STACK0_VIO_OUT	408
#define STACK1_VIO_OUT	409
#define STACK2_VIO_OUT	410
#define STACK3_VIO_OUT	411
#define STACK4_VIO_OUT	412
#define STACK5_VIO_OUT	413
#define STACK6_VIO_OUT	414
#define STACK7_VIO_OUT	415

//for arc weld
#define ARC_ON_VIO_OUT	403
#define SONGSI_VIO_OUT	436
#define DIANDONG_SONGSI_VIO_OUT	404
#define TUISI_VIO_OUT	405
#define ASPIRATION_VIO_OUT	432
#define START_ARC_CHECK_VIO_IN	32
#define ARC_MACHINE_ALARM_VOI_IN	33
#define IFARC_VIO_OUT_ST	433 //if arc technics
#define IFTRANSFER_VIO_OUT_ST	434 //if transfer technics
#define CAN_MOTOR_RUN_VOUT	465
//for hand
#define HSEN_VIO_IN	34 //34-35
#define HAND1_VIO_OUT	ARC_ON_VIO_OUT
#define HAND2_VIO_OUT	SONGSI_VIO_OUT
//end
#define EXT_SYNCHRONOUS	36

#define TRACK_CUTIN	37
#define GET_TRACK_CUTIN()	GET_VIN_BIT(TRACK_CUTIN)

#define INTERFERENCE_VIO_OUT_ST	440 //433-448
#define SET_INTERFERENCE_IO(n,v)	SET_VOUT_BIT((INTERFERENCE_VIO_OUT_ST+n),v)
#define GET_EXT_SYNCHRONOUS()	GET_VIN_BIT(EXT_SYNCHRONOUS)
//arc weld
#define GET_ARC_ON_IO()	GET_VOUT_BIT(ARC_ON_VIO_OUT)
#define SET_ARC_ON_IO(v)	SET_VOUT_BIT(ARC_ON_VIO_OUT,v)

#define SET_SONGSI(v)	 		SET_VOUT_BIT(SONGSI_VIO_OUT,v)
#define SET_DIANDONG_SONGSI(v)	 SET_VOUT_BIT(DIANDONG_SONGSI_VIO_OUT,v)
#define SET_DIANDONG_TUISI(v)	 SET_VOUT_BIT(TUISI_VIO_OUT,v)
#define SET_ARC_ASPIRATION(v)	 SET_VOUT_BIT(ASPIRATION_VIO_OUT,v)
#define SET_ARC_ENABLE(v)	 	SET_VOUT_BIT(IFARC_VIO_OUT_ST,v)
#define SET_TRANSFER_ENABLE(v)	SET_VOUT_BIT(IFTRANSFER_VIO_OUT_ST,v)

#define HAND1_ONOFF_VOUT()	GET_VOUT_BIT(HAND1_VIO_OUT)
#define HAND2_ONOFF_VOUT()	GET_VOUT_BIT(HAND2_VIO_OUT)
#define SET_HAND1_ONOFF(v)	SET_VOUT_BIT((HAND1_VIO_OUT),v)
#define SET_HAND2_ONOFF(v)	SET_VOUT_BIT((HAND2_VIO_OUT),v)
#define GET_HSEN_STATE(n)	GET_VIN_BIT((HSEN_VIO_IN+n))

#define SET_HAND_ONOFF(n,v) do{if(n==1) SET_HAND1_ONOFF(v);else if(n==2) SET_HAND2_ONOFF(v);}while(0)

#define ARC_START_CHECK_IO()	GET_VIN_BIT(START_ARC_CHECK_VIO_IN)
#define ARC_MACHINE_ALARM_IO()	GET_VIN_BIT(ARC_MACHINE_ALARM_VOI_IN)
#define ARC_DASH_CHECK_IO()	GET_VIN_BIT(DASH_ALARM_VIO_IN)

//ext machine option
#define EXT_EMEGENCY_STOP()	GET_VIN_BIT(EXT_EMERGENCY_STOP_VIO_IN)
#define EXT_START()	GET_VIN_BIT(EXT_START_VIO_IN)
#define EXT_PAUSE()	GET_VIN_BIT(EXT_PAUSE_VIO_IN)
#define EXT_SERVO_ON()	GET_VIN_BIT(EXT_SERVO_ON_VIO_IN)
#define EXT_CLEAR_ALARM()	GET_VIN_BIT(EXT_CLEAR_ALARM_IN)
#define EXT_SAFE_DOOR()	GET_VIN_BIT(EXT_SAFE_DOOR_IN)
//work stack book prog
#define set_book_enable_vio_out(v)	SET_VOUT_BIT(BOOKPROG_ENABLE_VIO_OUT,v)
#define get_stack_vio_in(n)	GET_VIN_BIT(STACK0_VIO_IN+n)
#define set_stack_vio_out(n,v)	SET_VOUT_BIT((STACK0_VIO_OUT+n),v)
#define get_stack_vio_out(n)	GET_VOUT_BIT (STACK0_VIO_OUT+n)
////////////////////////////////

#define break_vio_ctrl(v) SET_VOUT_BIT(BREAK_VIO_OUT,v)
#define breakio_ctrl(v) break_vio_ctrl(v)//do{SET_OUT_BIT(BREAK_IO_OUT,v);break_vio_ctrl(v);}while(0)
#define get_break_vio_ctrl() GET_VOUT_BIT(BREAK_VIO_OUT)
#define emergencystop() GET_IN_BIT(EMERGENCY_STOP_IN)
#define set_user_origin(v) SET_VOUT_BIT(USER_ORIGIN_VIO_OUT,v)
#define HOMEPOS_AXIS_VIO(n)	GET_VIN_BIT(HOMEPOS_AXIS_VIO_IN+n)

#define SetForceClearAlarm(v)	SET_VOUT_BIT(DASH_CLEAR_VIO_OUT,v)
#define GetForceClearAlarm()	GET_VOUT_BIT(DASH_CLEAR_VIO_OUT)

#define GET_BYTE(bit)	(bit>>3)
#define plc_setstate_stop()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(STOP_VIO_OUT),0x01)
#define plc_setstate_pause()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(STOP_VIO_OUT),0x02)
#define plc_setstate_emestop()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(STOP_VIO_OUT),0x04)
#define plc_setstate_running()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(STOP_VIO_OUT),0x08)
#define plc_setstate_error()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(STOP_VIO_OUT),0x10)

#define plc_getstate_stop()	GET_VOUT_BIT(STOP_VIO_OUT)
#define plc_getstate_pause()	GET_VOUT_BIT(PAUSE_VIO_OUT)
#define plc_getstate_emestop()	GET_VOUT_BIT(EMESTOP_VIO_OUT)
#define plc_getstate_running()	GET_VOUT_BIT(RUNNING_VIO_OUT)
#define plc_getstate_error()	GET_VOUT_BIT(ERROR_VIO_OUT)

#define plc_setmode_teach()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(TEACHMODE_VIO_OUT),0x01)
#define plc_setmode_auto()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(TEACHMODE_VIO_OUT),0x02)
#define plc_setmode_remote()	SET_ARRAY_BYTE(SHARE_RES(plc.PLC_M),GET_BYTE(TEACHMODE_VIO_OUT),0x04)

#define plc_getmode_teach()	GET_VOUT_BIT(TEACHMODE_VIO_OUT)
#define plc_getmode_auto()	GET_VOUT_BIT(AUTO_VIO_OUT)
#define plc_getmode_remote()	GET_VOUT_BIT(REMOTE_VIO_OUT)

//used IO in manual run robot
#define MANUAL_VIO_IN_SERVO_ON			72
#define MANUAL_VIO_IN_SERVO_OFF			73
#define MANUAL_VIO_IN_SPEEDUP		74
#define MANUAL_VIO_IN_SPEEDDOWN		75
#define MANUAL_VIO_IN_PRESS_START	40
#define MANUAL_VIO_IN_RELEASE_START	56

#define plc_manual_keyPress(n)	GET_VIN_BIT(MANUAL_VIO_IN_PRESS_START+n)
#define plc_manual_keyRelease(n)	GET_VIN_BIT(MANUAL_VIO_IN_RELEASE_START+n)
#define plc_manual_servo_on()		GET_VIN_BIT(MANUAL_VIO_IN_SERVO_ON)
#define plc_manual_servo_off()	GET_VIN_BIT(MANUAL_VIO_IN_SERVO_OFF)
#define plc_manual_speedUp()	GET_VIN_BIT(MANUAL_VIO_IN_SPEEDUP)
#define plc_manual_speedDown()	GET_VIN_BIT(MANUAL_VIO_IN_SPEEDDOWN)

#endif 

