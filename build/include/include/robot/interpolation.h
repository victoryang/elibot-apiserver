#ifndef _INTERPOLATION_H_
#define _INTERPOLATION_H_

#include <robot/compile.h>
#include <robot/robotresource.h>
#include <robot/parameters.h>
#include <robot/ExtAxis.h>
#include <robot/math_func.h>


#define SPEED2POS_CHIP(s)	((s)*PERIOD)

#define USE_MOVL_NEW_PL
//#undef USE_MOVL_NEW_PL

//for queue
typedef struct
{
	double startPos;			//起始点位置
	double targetPos;			  //目标位置
	double gearSpeed;			 //每个轴的最大速度
	double acc;					 //加速度
	double dec;					 //减速度 
}AutoMovjData;

typedef AutoMovjData ManualMovjData;

typedef struct
{
	double startPos;			//起始点位置
	double targetPos;			  //目标位置
	double gearSpeed;			 //每个轴的最大速度
	double initialSpeed;			//每个轴运动的初始速度
	double compileSpeed;		//set by user teach ,don't change it
	double endSpeed;				//每个轴运动的末速度
	double acc;					 //每个轴允许最大加速度
	double jcc; 								//每个轴加加速度
	double accperiod;                    //每个轴的加速时间周期个数，减速时间一样，与acc/jcc二选一
	int combineFlag;               //是否一起读入下一段标志，0:不用1:用
	int wait_tick;

	double motionsmooth;
	double acctime;
	double overlaptime;
	double k;
	unsigned int pl_level;
}AutoMovjcomm;

typedef AutoMovjcomm ManualMovjcomm;

struct _AutoMovjDataS;
typedef int (*automovj_speed_func_S)(struct _AutoMovjDataS *interdata);

typedef struct _AutoMovjAxisInterpDataS{

	automovj_speed_func_S change_speed;
	int moveDirection;				//每个轴行走的方向
	int currentFinish;              //0:need slide and no stop; 1:normal stop; 2:stop exactly endpoint
	double tempSpeed;				//匀速段重新计算的速度
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
//	int currentFinish;// 2:在当前段刚好结束0:在当前段没有结束1:在当前段停了但还有剩余位移

}AutoMovjAxisInterpDataS;

typedef AutoMovjAxisInterpDataS ManualMovjAxisInterpDataS;

typedef struct _AutoMovjDataS{
	AutoMovjcomm automovj;
	AutoMovjAxisInterpDataS movjdata;
}AutoMovjDataS;

typedef AutoMovjDataS ManualMovjDataS;

typedef struct {
	double startPos[AXIS_COUNT];   //zxc axis data
	double midPos[AXIS_COUNT];	 //use for movc
	double targetPos[AXIS_COUNT]; //zxc axis data 
	double gearSpeed;			//直线运动的最大速度
	double vrSpeed; 					 //旋转速度
	double vracc;						 //旋转加速度
	double vrdec;						 //旋转减速度
	double acc;					//加速度
	double dec;					//减速度

}AutoMovlData;

typedef struct{
	double startPos[AXIS_COUNT];  
	double targetPos[AXIS_COUNT];
	//velocity related
	double gearSpeed; //直线运动的最大速度
	double compileSpeed;
	double initialSpeed; //每个轴运动的初始速度
	double endSpeed; //每个轴运动的末速度
	double vrSpeed; //旋转速度
	double vracc; //旋转加速度
	double vrjcc; //旋转加加速度
	double acc; //加速度
	double jcc;//每个轴加加速度
	int wait_tick;
	int coop;//协同情况
	int movetype;//base on: 0-line move ,1-rotate move ,2-no move
	int moveSelector; // 1: line 2:clockwise circle 3:counter clockwise circle 4:parabola 
	double accperiod;//acceleration period
	int combineFlag;               //是否一起读入下一段标志，0:不用1:用
	int currentFinish;              //0:need slide and no stop; 1:normal stop; 2:stop exactly endpoint
	AutoExtAxisInterpDataS extAxisData[EXTAXIS_COUNT];
}AutoMovCommDataS;

typedef AutoMovlData ManualMovlData;
typedef AutoMovCommDataS ManualMovlDataS;

typedef AutoMovlData	AutoMovcData;

//typedef AutoMovCommDataS AutoMovlDataS;
struct _AutoMovCommonS;
typedef int (*automov_speed_func_S)(struct _AutoMovCommonS *interdata);

typedef struct _AutoMovCommonS{
	automov_speed_func_S change_speed;//zxc指示加速，匀速，减速
	int tick;					//当前周期数(当前在哪个周期)
	double starMax[4][4];			//起点位姿矩阵
	double targetMax[4][4];		//目标位姿矩阵
	double matrix[4][4];			//当前位姿矩阵
	double ff[3];					//位姿旋转矢量
	double sAngle;				//位姿旋转角度单位:弧度
	double displacement;			//圆弧/直线的长度	 抛物线弦长
	double linedisplacement;	//only pure line move 
	double gearSpeed;			//运动的最大速度
	double initialSpeed;			//运动的初始速度
	double endSpeed;				//运动的末速度
	double tempSpeed;			//匀速行走时的速度，临时数据。
	double AccAccPeriod;	   //加加速段
	double UniAccPeriod;	   //匀加速段
	double DecAccPeriod;	  //减加速段
	double uniformPeriod;    //匀速段
	double AccDecPeriod;	  //加减速段
	double UniDecPeriod;	  //匀减速段
	double DecDecPeriod;	 //减减速段
	double AccJcc;						   //加速段重新计算的加加速度
	double AccJcc1;                                        //减加速段的J
	double Acc; 							 //加速段重新计算的加速度值
	double DecJcc;						   //减速段重新计算的加加速度
	double DecJcc1;                                         //减减速段的J
	double Dec; 							 //减速段重新计算的加速度值 	
	double stopDec;				//Stop 减速度 
	int outside;				//六轴越界的变量
	int over;						//轴运动是否停止
	int pause;					//暂停信号
	int lable;					//用于临时记录暂停前的速度标示
	int Mark;					//标示运行完毕但减速未结束
	double curSpeed;				//当前速度
	double residuelength;		//剩余长度	
	//int currentFinish;              //0:need slide and no stop; 1:normal stop; 2:stop exactly endpoint
	double AccPeriodInitAcc;
	double AccPeriodEndAcc;
	double DecPeriodInitAcc;
	double DecPeriodEndAcc;
	double DecPeriodInitSpeed;
	double AccDisplacement;
	double DecDisplacement;
	double UniDisplacement;
	int StartPointZoneFlag;   //startpoint zone, 0:accacc/accdec, 1:uniacc/unidec, 2:decacc/decdec
	int EndPointZoneFlag;    //endpoint zone, 0:accacc/accdec, 1:uniacc/unidec, 2:decacc/decdec
	double InitialAcc;
	double EndAcc;
	int DecPeriodFirstSegFlag;//减速区第一个程序段标志
	double PeriodRatio;//周期时间长度的百分比
	double FinalAxisAngle;
	double accperiod;                    //加速时间周期个数，减速时间一样，与acc/jcc二选一
//	int combineFlag;               //是否一起读入下一段标志，0:不用1:用

}AutoMovCommonS;

typedef struct _AutoMovcInterpDataS{
	AutoMovCommonS comm;
	double star[3];				//变换到平面的起始点
	double mid[3];				//变换到平面的参考点
	double target[3];				//变换到平面的终点
	double afa;					//变换到平面的起始点弧度。（0~2*pi）
	double sit;					//累加的旋转角度
	double sita;					//圆弧的圆心角（0~2*pi）
	int symbol;					//是否可以组成圆的标志
	int direction;					//圆弧行走的方向（从上往下看，顺时针或逆时针）1顺时针0逆时针
	double radius;				//圆的半径
	double center[3];				//圆心
	double matrix[3][3];			//XOY平面圆到空间圆的转换矩阵
	double tempAngle;	 //记录走过的角度
	double vrSpeed;
	double finalAxisSita;
}AutoMovcInterpDataS;

typedef struct{
	AutoMovCommDataS autoMovS;
	AutoMovcInterpDataS movcdata;

	double midPos[AXIS_COUNT];	
	double includedAngle[2]; //空间曲线(圆弧/抛物线)至平面曲线的旋转角度
	double matrix[3][3];	//XOY平面抛物线到空间抛物线的转换矩阵
	//circle related
	double circleR; //circle radius 
	double circleCenter[3]; //circle center 
}AutoMovcDataS;

typedef struct _AutoMovlInterpDataS{
	AutoMovCommonS comm;
	double sinAngle[3];			//沿直线行走时直线和XYZ轴夹角的正弦值	 
	double sita;						//累加的旋转角度
	double vrSpeed;
	double finalAxisSita;
}AutoMovlInterpDataS;

typedef struct{
	AutoMovCommDataS autoMovS;
	AutoMovlInterpDataS movldata;
}AutoMovlDataS;

typedef AutoMovlDataS AutoMovnDatas;

typedef struct _AutoMovaInterpDataS{
	AutoMovCommonS comm;
	double joint[AXIS_COUNT];
}AutoMovaInterpDataS;

typedef struct{
	AutoMovCommDataS autoMovS;
	AutoMovaInterpDataS movadata;
}AutoMovaDataS;

typedef struct _AutoMMovInterpData{
	int subp;
	double tooldist;
	double measuredist;
}AutoMMovInterpData;

typedef struct{
	AutoMovCommDataS autoMovS;
	AutoMMovInterpData mmovdata;
}AutoMMovData;

typedef struct _AutoMovsInterpDataS{
//	automovs_speed_func_S change_speed;//zxc指示加速，匀速，减速
	AutoMovCommonS comm;
	double star[3];				//变换到平面的起始点
	double mid[3];				//变换到平面的参考点
	double target[3];				//变换到平面的终点
	double ratio;                	 //x/l 的比例
	double aa;                     //抛物线的方程参数
	double parabolaVertex[3];	//抛物线的顶点
//	double afa;					//变换到平面的起始点弧度。（0~2*pi）
	double sit;					//累加的旋转角度
	double matrix[3][3];			//XOY平面抛物线到空间抛物线的转换矩阵
//	double tempAngle;	 //记录走过的角度
	double vrSpeed;
	double finalAxisSita;
}AutoMovsInterpDataS;


typedef AutoMovlData AutoMovsData;
typedef struct{
	AutoMovCommDataS autoMovS;
	AutoMovsInterpDataS movsdata;
	
	double midPos[AXIS_COUNT];	
	double includedAngle[2]; //空间曲线(圆弧/抛物线)至平面曲线的旋转角度
	double matrix[3][3];	//XOY平面抛物线到空间抛物线的转换矩阵
	//parabola related
	double parabolaVertex[3];	//抛物线的顶点
	double angleS;				//到xoy平面后需转动的角度，特指抛物线
	double aa;					   //抛物线的方程参数
	double ratio;			 //x/l 的比例
	double displacement;  //抛物线弦长
	double star[3]; 	 //变换到平面的起始点
	double target[3];		 //变换到平面的终点
}AutoMovsDataS;

//sinx
typedef struct _AutoMovxInterpDataS{
	AutoMovCommonS comm;

	double ReferenceJoint[AXIS_COUNT];//参考点
	double InitPhase;//初始相位
	double WaveLength;//波长
	double Amplitude;//振幅
	
	double includedAngle[2]; //空间曲线至平面曲线的旋转角度
	double angleS;				//到xoy平面后需转动的角度
	double star[3];				//变换到平面的起始点
	double target[3];				//变换到平面的终点
	double omega;   //w
	double matrix[3][3];	//XOY平面正弦线到空间正弦线的转换矩阵
	
	double vrSpeed;
	double finalAxisSita;
	double sit;					//累加的旋转角度
	
}AutoMovxInterpDataS;

typedef struct{
	AutoMovCommDataS autoMovS;
	AutoMovxInterpDataS movxdata;

//	double ReferenceJoint[AXIS_COUNT];//参考点
//	double InitPhase;//初始相位
//	double WaveLength;//波长
//	double Amplitude;//振幅
	
//	double includedAngle[2]; //空间曲线至平面曲线的旋转角度
//	double angleS;				//到xoy平面后需转动的角度
//	double star[3];				//变换到平面的起始点
//	double target[3];				//变换到平面的终点
//	double displacement;  //正弦波长度
//	double omega;   //w
//	double matrix[3][3];	//XOY平面正弦线到空间正弦线的转换矩阵
}AutoMovxDataS;
//sinx end 
typedef struct _AutoMovMLInterpData{
	enum Robot_Coordinate rc;
	double cpose[6];
	char file[FILE_NAME_LENGTH];
}AutoMovMLInterpData;

typedef struct{
	AutoMovCommDataS autoMovS;
	AutoMovMLInterpData movmldata;
}AutoMovMLData;

struct _MoveData;
typedef int (*move_func)(struct _MoveData *md);
typedef STRU_INSTRUCT* (*notmove_exec_func)(STRU_INSTRUCT *instruct, FILE *pfout);

typedef struct _MoveData{
	move_func func;  //zxc 指向插补函数
	notmove_exec_func notmove_func;
	STRU_INSTRUCT *cur_instr;
	union{
		AutoMovjData autoMovj[AXIS_COUNT];
		AutoMovjDataS autoMovjS[AXIS_COUNT];
		AutoMovCommDataS autoMovS;
		AutoMovlData autoMovl;
		AutoMovlDataS autoMovlS;
		AutoMovcData autoMovc;
		AutoMovcDataS autoMovcS;
		AutoMovsData autoMovs;
		AutoMovsDataS autoMovsS;
		AutoMovxDataS autoMovxS;
		AutoMMovData autoMMov;
		AutoMovaDataS autoMova;
		AutoMovMLData autoMovML;
		AutoMovnDatas autoMovn;
	};
}MoveData_t;

//for inner
struct _AutoMovjAxisInterpData;
typedef int (*automovj_speed_func)(AutoMovjData *movjdata, struct _AutoMovjAxisInterpData *interdata);

typedef struct _AutoMovjAxisInterpData{
	automovj_speed_func change_speed;

	int moveDirection;				//每个轴行走的方向
	double tempSpeed;				//匀速行走时的速度，临时数据。
	double curSpeed;				//当前速度
	unsigned int accPeriod;			//加速周期
	unsigned int decPeriod;			//减速周期
	unsigned int uniformPeriod;		//匀速周期
	unsigned int tick;				//周期Tick
}AutoMovjAxisInterpData;
typedef AutoMovjAxisInterpData ManualMovjAxisInterpData;

typedef struct {
	double starMax[4][4];			//起点位姿矩阵
	double targetMax[4][4];		//目标位姿矩阵
	double matrix[4][4];			//当前位姿矩阵
	double ff[3];					//位姿旋转矢量
	double sAngle;				//位姿旋转角度
	double displacement;			//圆弧/直线的长度	 
	double gearSpeed;			//运动的最大速度
	double initialSpeed;			//运动的初始速度
	double endSpeed;				//运动的末速度
	double tempSpeed;			//匀速行走时的速度，临时数据。
	int accPeriod;				//加速周期
	int uniformPeriod;			//匀速周期
	int decPeriod;				//减速周期
	int accelerateN;			//加速周期数
	int decelerateN;			//减速周期数
	int uniformelerateN;		//匀速周期数
	double acc;					//加速度
	double dec;					//减速度 
	double stopDec;				//Stop 减速度 
	int outside;				//六轴越界的变量
	int over;						//轴运动是否停止
	int pause;					//暂停信号
	int lable;					//用于临时记录暂停前的速度标示
	int Mark;					//标示运行完毕但减速未结束
	double curSpeed;				//当前速度
	double residuelength;		//剩余长度	
}AutoMovCommon;

typedef struct _AutoMovlInterpData{
	AutoMovCommon comm;
	double sinAngle[3];			//沿直线行走时直线和XYZ轴夹角的正弦值	 
	double sita;						//累加的旋转角度
	int tick;					//当前周期数(当前在哪个周期)
	int moveType;				//旋转或直线运动类型
	double vrSpeed;
}AutoMovlInterpData;


typedef struct _AutoMovcInterpData{
	AutoMovCommon comm;
	double star[3];				//变换到平面的起始点
	double mid[3];				//变换到平面的参考点
	double target[3];				//变换到平面的终点
	double afa;					//变换到平面的起始点弧度。（0~2*pi）
	double sit;					//累加的旋转角度
	double sita;					//圆弧的圆心角（0~2*pi）
	int symbol;					//是否可以组成圆的标志
	int direction;					//圆弧行走的方向（从上往下看，顺时针或逆时针）
	double radius;				//圆的半径
	double center[3];				//圆心
	double matrix[3][3];			//XOY平面圆到空间圆的转换矩阵
	double tempAngle;	 //记录走过的角度
}AutoMovcInterpData;


struct _ManualMovlInterpDataS;
typedef int (*manualmovl_speed_func_S)(struct _ManualMovlInterpDataS *interdata);

typedef struct _ManualMovlInterpDataS{

	manualmovl_speed_func_S change_speed;
	double poseMatrix[4][4];		// 在某坐标系下的位姿矩阵
	double toolMatrix[4][4];		//当前工具的数据矩阵
	double useMatrix[4][4]; 		//当前用户的数据矩阵
	double tempSpeed;			//匀速行走时的速度，临时数据。
	double gearSpeed;				//运动的最大速度
	double tempAccJcc;						   //加速段重新计算的加加速度
	double tempAccJcc1;
	double tempAcc; 							 //加速段重新计算的加速度值
	double tempDecJcc;						   //减速段重新计算的加加速度
	double tempDecJcc1; 					  //重新计算的减减速区的加加速度
	double tempDec; 							 //减速段重新计算的加速度值 	
	double curSpeed;					//当前速度
	double AccAccPeriod;	   //加加速段
	double UniAccPeriod;	   //匀加速段
	double DecAccPeriod;	  //减加速段
	double uniformPeriod;    //匀速段
	double AccDecPeriod;	  //加减速段
	double UniDecPeriod;	  //匀减速段
	double DecDecPeriod;	 //减减速段
	int over;								//轴运动是否停止
	int beyond; 						//关节角度越界
	int tick;				//周期Tick
	double finalAxisSita;

}ManualMovlInterpDataS;

//////////////////////PL相关数据结构/////////////////////////////

//PL MOVJ
typedef struct _MovjPlInterpData
{
	double movjSpeed[AXIS_COUNT];			 //每个轴的最大速度
	double movjacc[AXIS_COUNT];					//movj加速度
	double movjdec[AXIS_COUNT];					//movj减速度 
	double movcSpeed;				 //圆弧的最大速度。
	double movcacc; 				  //圆弧的加速度
	double movcdec; 				  //圆弧的减速度
	int pl; 					  //中间点的pl等级
	double plLength ;			   //pl的初始参量
	double midMatrix[4][4]; 	   //参考点位姿矩阵

}MovjPlInterpData;

//PL MOVJ
typedef struct _MovjPlInterpDataS
{
	double movjSpeed[AXIS_COUNT];			 //每个轴的最大速度
	double movjacc[AXIS_COUNT];					//movj加速度
	double movjjcc[AXIS_COUNT];					//movj减速度 
	double movcSpeed;				 //圆弧的最大速度。
	double movcacc; 				  //圆弧的加速度
	double movcjcc; 				  //圆弧的减速度
//	int pl; 					  //中间点的pl等级
	double plLength ;			   //pl的初始参量
//	double midMatrix[4][4]; 	   //参考点位姿矩阵
}MovjPlInterpDataS;

typedef struct _PlInterpData {
	double lpoint1[6];			  //直线起点
	double lStarSpeed;			  //直线的起始速度
	double lpoint2[6];			  //直线终点
	double cpoint1[6];			  //圆弧起点
	double cpoint2[6];			  //圆弧终点
	double Speed1;				  //运行的p1p2的最大速度
	double Speed2;				  //运行时p2p3的最大速度
	double acc; 				  //运行的加速度
	double dec; 				  //运行的减速度
	int pl; 					  //中间点的pl等级
	double plLength ;			  //pl的初始参量
	double crlenth; 			  //cr的半径
	double midMatrix[4][4]; 			 //参考点位姿矩阵
} PlInterpData;

typedef struct _PlInterpDataS {
	double speed;				 //运行的p1p2的最大速度
	double acc; 				  //运行的加速度
	double jcc;
	double vrspeed;
//	int pl; 					  //中间点的pl等级
	double plLength ;			  //pl的初始参量
} PlInterpDataS;


typedef struct _MovcPlInterpData
{
	double lpoint1[6];			  //直线起点
	double lStarSpeed;			  //直线的起始速度
	double lpoint2[6];			  //直线终点
	double cpoint1[6];			  //圆弧起点
	double cpoint2[6];			  //圆弧终点
	double Speed1;				  //运行的p1p2的最大速度
	double Speed2;				  //运行时p2p3的最大速度
	double acc; 				  //运行的加速度
	double dec; 				  //运行的减速度
	int pl; 					  //中间点的pl等级
	double plLength ;			   //pl的初始参量
	double midMatrix[4][4]; 			 //参考点位姿矩阵
	double sita;					//第一段圆弧所走的弧度
	double targetMatrix[4][4];			  //第一段圆弧的末矩阵

}MovcPlInterpData;

typedef struct _AccCombineData
{
	double initialSpeed;
	double endSpeed;	
	double displacement;
	double jcc;
	double acc;
	double AccAccTime;
	double UniAccTime;
	double AccAccdisplacement;
	double UniAccdisplacement;
	
}AccCombineData;

/***************for vision grasp and fold clothes*************************/
typedef struct _fold_data
{
	double pt[AXIS_COUNT];
	vec3 dir;
	double disp;
	double accumulator_disp;
}fold_data;
#define MOVML_POS_SIZE 50000
#define MOVML_SEGMENT_LEN 20
typedef struct auto_movml_interp_seg_pose_s {
	int num;
	int cur_num;
	double pose[MOVML_POS_SIZE][AXIS_COUNT];
} auto_movml_interp_seg_pose_t;

typedef struct auto_movml_interp_pose_s {
	int segment_num;
	int cur_seg_num;
	auto_movml_interp_seg_pose_t seg_poses[MOVML_SEGMENT_LEN];
} auto_movml_interp_pose_t;

//extern double extrinsic[4][4];
//extern double intrinsic[3][3];
int set_catchPos(double pt[]);
int fold_interpolation(MoveData_t *md,auto_movml_interp_seg_pose_t *input);
void reset_movml_pos(void);
void reset_movml_pos_fold(void);

int MOVML_interpolation_vision(MoveData_t *md );
int MOVML_interpolation_normal(MoveData_t *md );

/***********************************************************************/
int interpolation_init(void);
void interpolation_exit(void);

void wait_mov();
void wait_move_queue_empty(void);
void clean_move_queue(void);

MoveData_t* get_move_data_head(int *mov_done);
MoveData_t* get_last_move_data(int head_back);

void put_move_data_head(void);
MoveData_t* get_move_data_update(unsigned int offset);
void put_move_data_update(void);
void begin_ForeRead(void);
void update_move_speed_all(void);
void send_interpInfo_signal(void);	
void put_technic_mov_queue(void);

void AutoMovPreProcess_S(AutoMovCommonS *movldata);
void MOVFW_PreProcess(MoveData_t *md, int number);
void Weave_PreProcess_S(MoveData_t *md);
void ExtPreDataProcess_S0(MoveData_t *md, int combineFlag);

double AutoMovlInterpolation_S(double dertR[],double xp_joint[],MoveData_t *md);
double AutoMovcInterpolation_S(double dertR[],double starJoint[],MoveData_t *md);
double AutoMovsInterpolation_S(double dertR[],double starJoint[],MoveData_t *md);

//int MOVJ_interpolation_S(MoveData_t *md);
int MOVL_interpolation_S(MoveData_t *md);
int MOVC_interpolation_S(MoveData_t *md);
int MOVFW_interpolation_S(MoveData_t *md);
int MOVS_interpolation_S(MoveData_t *md);
int MOVX_interpolation_S(MoveData_t *md);
int MOVML_interpolation_S(MoveData_t *md);
int MMOVE_interpolation_S(MoveData_t *md);
int PREMMOV_interpolation_S(MoveData_t *md );
int MOVA_interpolation_S(MoveData_t *md );
int MOVN_interpolation_S(MoveData_t *md);

//摆焊工艺
int MovcWeavingInterpolationApi(MoveData_t *md);
int MovlWeavingInterpolationApi(MoveData_t *md);

//鱼鳞焊接工艺
int MovlStitchWeldInterpolationApi(MoveData_t *md);
int MovcStitchWeldInterpolationApi(MoveData_t *md);

//激光焊缝跟踪工艺
int WeldSeamSearchMovlInterpolationApi(MoveData_t *md);
int WeldSeamTrackMovlInterpolationApi(MoveData_t *md);

//针对跟踪工艺
int TrackInterpolationApi(MoveData_t *md);
int TrackGraspInterpolationApi(MoveData_t *md);
int TrackLiftupInterpolationApi(MoveData_t *md);

int get_weaving_file(void);
int get_arc_weld_file(void);

void execute_next_io_code();
//double CalNextX(double currentX, double aa, double stepLength, double ratio);
double SearchResult_1(double MinSpeed, double MaxSpeed, double InitialSpeed, double EndSpeed, double Acc, double Jcc, double distance);
double SearchResult_2(double MinSpeed, double MaxSpeed, double InitialSpeed, double EndSpeed, double Acc, double Jcc, double distance);
double SearchResult_3(double MinSpeed, double MaxSpeed, double InitialSpeed, double EndSpeed, double Acc, double Jcc, double distance);

int Auto_Movement_S(AutoMovCommonS *mov_comm, MoveData_t *md);
int AutoMov_AxisDecDecSpeed(AutoMovCommonS *interdata);
int AutoMov_AxisUniDecSpeed(AutoMovCommonS *interdata);
int AutoMov_AxisAccDecSpeed(AutoMovCommonS *interdata);
int AutoMov_AxisSUniformSpeed(AutoMovCommonS *interdata);
int AutoMov_AxisDecAccSpeed(AutoMovCommonS *interdata);
int AutoMov_AxisUniAccSpeed(AutoMovCommonS *interdata);
int AutoMov_AxisAccAccSpeed(AutoMovCommonS *interdata);
int update_automovl_dataS(MoveData_t *md, STRU_MOV *mov, int mov_done);
double MOVL_Stepinterplation_S(double *poschip,MoveData_t *md,double *machinePos,int *ext_over);
double MOVC_Stepinterplation_S(double *poschip,MoveData_t *md,double *machinePos,int *ext_over);
double MOVS_Stepinterplation_S(double *poschip,MoveData_t *md,double *machinePos,int *ext_over);
MoveData_t * reUpdata_movData_forStop(MoveData_t *md,double curPos[]);
double GET_EXT_CURSPEED(int n);
MoveData_t* get_move_data_tail(void);
void put_move_data_tail(void);
double getCurSpeed(void);
void setCurSpeed(double speed);
void set_cur_joint_speed(int axis, double v);
int get_cur_joint_speed(int axis);

void set_start_arcPoint(double *deltaPos);
void set_JUMP_CALL_INSTRUCT(STRU_INSTRUCT *instru);
STRU_INSTRUCT * get_JUMP_CALL_INSTRUCT();

#endif
