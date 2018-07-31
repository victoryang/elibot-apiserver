#ifndef _EXTAXIS_H_
#define _EXTAXIS_H_

struct _AutoExtAxisInterpDataS;
typedef int (*autoext_speed_func_S)(struct _AutoExtAxisInterpDataS *interdata);

typedef struct _AutoExtAxisInterpDataS{

	autoext_speed_func_S change_speed;
	double startPos;			//起始点位置
	double targetPos;			  //目标位置
//	double TotalPeriod;
	int moveDirection;				//每个轴行走的方向
	double gearSpeed;				//匀速段重新计算的速度
	double initialSpeed; //每个轴运动的初始速度
	double endSpeed; //每个轴运动的末速度
	double tempSpeed;
	double tempAccJcc;						   //加速段重新计算的加加速度
	double tempAccJcc1;
	double tempAcc; 							 //加速段重新计算的加速度值
	double tempDecJcc;						   //减速段重新计算的加加速度
	double tempDecJcc1; 					  //重新计算的减减速区的加加速度
	double tempDec; 							 //减速段重新计算的加速度值 	
	double curSpeed;				//当前速度
	double AccAccPeriod;	   //加加速段
	double UniAccPeriod;	   //匀加速段
	double DecAccPeriod;	  //减加速段
	double uniformPeriod;    //匀速段
	double AccDecPeriod;	  //加减速段
	double UniDecPeriod;	  //匀减速段
	double DecDecPeriod;	 //减减速段
	int tick;				//周期Tick
	double PeriodRatio;//周期时间长度的百分比
	double residuelength;		//剩余长度	
	int combineFlag;

}AutoExtAxisInterpDataS;

int ExtAxisIsEnableMove(double strJoint[],double endJoint[], int coop);
void GetExt2Base(double currentJoint[],double matrix[][4], int coop);
//void ExtPreDataProcess_S(MoveData_t *md,int combineFlag);
//void AutoExtInterpolation_S(int isStop,AutoExtAxisInterpDataS axisdata[],double poschip[]);
int Ext_Interplation_S(AutoExtAxisInterpDataS axisdata[],double poschip[],double joint[], int coop);
void Ext_Interplation_Dual_S(int isStop,AutoExtAxisInterpDataS axisdata[],AutoExtAxisInterpDataS axisdata1[],double poschip[],double joint[]);
void initExtAxis(void );

int Ext_AxisAccDecSpeed(AutoExtAxisInterpDataS *interdata);
#endif // EXTAXIS_H


