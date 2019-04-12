#ifndef _EXTAXIS_H_
#define _EXTAXIS_H_

#ifdef __cplusplus
extern "C" {
#endif

#define MAX_AXIS_GROUP    2
struct _AutoExtAxisInterpDataS;
typedef int (*autoext_speed_func_S)(struct _AutoExtAxisInterpDataS *interdata);

typedef struct _AutoExtAxisInterpDataS{
	autoext_speed_func_S change_speed;
	double startPos;    //起始点位置
	double targetPos;	//目标位置
	int moveDirection;	//每个轴行走的方向
	double gearSpeed;	//匀速段重新计算的速度
	double initialSpeed; //每个轴运动的初始速度
	double endSpeed;    //每个轴运动的末速度
	double tempSpeed;
	double Acc;
	double AccJcc;
	double Dec;
	double DecJcc;
	double AccDisplacement;
	double DecDisplacement;
	double UniDisplacement;
	double tempAccJcc;	//加速段重新计算的加加速度
	double tempAccJcc1;
	double tempAcc; 	//加速段重新计算的加速度值
	double tempDecJcc;	//减速段重新计算的加加速度
	double tempDecJcc1; //重新计算的减减速区的加加速度
	double tempDec; 	//减速段重新计算的加速度值
	double curSpeed;	//当前速度
	double AccAccPeriod;    //加加速段
	double UniAccPeriod;	//匀加速段
	double DecAccPeriod;	//减加速段
	double uniformPeriod;   //匀速段
	double AccDecPeriod;	//加减速段
	double UniDecPeriod;	//匀减速段
	double DecDecPeriod;    //减减速段
	int tick;				//周期Tick
	double PeriodRatio;     //周期时间长度的百分比
	double residuelength;   //剩余长度
	int combineFlag;
}AutoExtAxisInterpDataS;

typedef struct _SingleExtAxisData{
	char note[32];		  //工艺描述
	unsigned char useable;//当前外部轴组是否可用
	unsigned char coordAxis;
	unsigned char if_teached[EXTAXIS_COUNT];
	unsigned char dirct[EXTAXIS_COUNT];
	unsigned char caliSuccess[EXTAXIS_COUNT];
	unsigned char coordEnable[EXTAXIS_COUNT];
	unsigned char extaxis_type[EXTAXIS_COUNT];
	int recordsafe_joint;
	int extaxis_number;
	double ext_joint[EXTAXIS_COUNT][3][AXIS_COUNT];
	double saft_joint[AXIS_COUNT];
	double ExtAxisMax[EXTAXIS_COUNT][4][4];
}SingleExtAxisData;

typedef struct{
	SingleExtAxisData CoextaxisDataS[MAX_AXIS_GROUP];
	int CurExtAxisGroupNum;    //当前协同外部轴轴组代号（参数设置时使用）
	int AxisCoopState;         //自动模式下外部轴使用情况
}ExtAxisData;

int ExtAxis_CoordinateBuild(double FirstJoint[], double SecondJoint[], double ThirdJoint[], double matrix[][4], int direct);
int ExtAxisIsEnableMove(double strJoint[], double endJoint[], int coop);
void PositiveBase2Ext(double matrix[][4], double joint[], int coop);
void InverExt2Base(double matrix[][4], double target[][4], double joint[], int coop);
void GetExt2Base(double currentJoint[], double matrix[][4], int coop);
int Ext_Interplation_S(AutoExtAxisInterpDataS axisdata[], double poschip[], double joint[], int coop);
void Ext_Interplation_Dual_S(int isStop, AutoExtAxisInterpDataS axisdata[], AutoExtAxisInterpDataS axisdata1[], double poschip[], double joint[]);
void initExtAxis(void);
int Ext_AxisAccDecSpeed(AutoExtAxisInterpDataS *interdata);
int CalculateStartNumber(int ExtGroupNumber);
void init_coextaxis(ExtAxisData *ext_axis_data);
void ExtPrePeriodProcess_S(AutoExtAxisInterpDataS *extmovedata);

#ifdef __cplusplus
}
#endif

#endif
