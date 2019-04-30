#ifndef __SEND_POS_H__
#define __SEND_POS_H__

#include "../../config.h"

enum	OperationMode {DynOpMode_Normal, DynOpMode_DirTch,DynOpMode_FeedForward};//Normal-for normal operation;DirTch-for Direct Teach
enum    SetDynParas{SetLoad,SetFriction};
#if 0
typedef struct
{
	double *Pos;/*joint position /velocity /acceleration */
	double *Vel;
	double *Acc;
	double *Trq;//joint torque
} MotionData_PVAT;
#else
typedef struct
{
	double Pos[AXIS_COUNT];/*joint position /velocity /acceleration */
	double Vel[AXIS_COUNT];
	double Acc[AXIS_COUNT];
	double Trq[AXIS_COUNT];//joint torque
} MotionData_PVAT;
#endif

typedef struct
{
	MotionData_PVAT Ref;// command/Reference motion data: Pos Vel Acc Trq
	MotionData_PVAT Fb;// feedback motion data: Pos Vel Acc
} InvDyn_MotionData;

typedef struct sendpos_struct{
	int32_t *pulse;	//add positoin pulse
	int32_t *abs_pulse;	//for absolutely pulse
	double *chip;	//add positoin
	double *pos;	//absolute position
	double *user;
	double *speed;
	double periodratio;
	InvDyn_MotionData motionData;
}sendpos_t;

typedef struct dynamicterms_struct{
	double Mv[AXIS_COUNT];
	double Cv[AXIS_COUNT];
	double gv[AXIS_COUNT];
	double dv[AXIS_COUNT];
}dynamicterms_t;

#endif
