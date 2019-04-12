#ifndef _TRACKCRAFT_H_
#define _TRACKCRAFT_H_

#include <rt_thread.h>

#ifdef __cplusplus
extern "C" {
#endif

#define TRACK__TECHNICS_COUNT    8
#define Track_Buffer_Size      128
#define TRACK_GPNum1      50
#define TRACK_GPNum2      51
#define ObjectNumber      2

//直线型加减速标志
typedef enum _TVPsign
{
	Uni_Acc = 0,
	CONST_SPEED,
	Uni_Dec,
	End_TVP
}TVPsign;

//传送装置类型
typedef enum _ConveyorType
{
	StraightLine = 0,
	Disc
}ConveyorType;

//标定方式
typedef enum _CalibrationMode
{
	NoVision = 0,
	MonocularVision	
}CalibrationMode;

typedef struct _TrackCraftData{

//**********************************************************************//
//																		//
//					   跟踪第二级界面：跟踪工艺设置界面									//
//																		//
//**********************************************************************//

	int Mode;             //功能选择
	int Encodertype;      //编码器号
	int IOtype;           //低速IO：0 高速IO：1
	int EncoderAcount;    //编码器A数值显示
	int EncoderBcount;    //编码器B数值显示
	int EncoderCcount;    //编码器C数值显示
	int Checkcount;       //I\O切入点编码器值
	ConveyorType ConType; //0：直线型传送带 1：圆盘型传送带
	//int CalibraMode;      //0：IO接触开关（不带视觉） 1：单目视觉
	double DiscRadius;    //圆盘型传送带半径(mm)
	double AccTime;       //跟踪加速度(s)
	double EffStaPoint;   //跟踪有效范围开始点（从A点开始）
	double EffEndPoint;   //跟踪有效范围结束点（从A点开始）
	double directionX;    //跟踪位置点补偿X方向（传送带方向）
	double directionY;    //跟踪位置点补偿Y方向
	double directionZ;    //跟踪位置点补偿Z方向
	double CheckRange;    //检测范围
	double SpeedAdj;      //跟踪中速度调节
	double StaSprDis;     //开始喷距离
	double WorkSpa;       //工件间距
	double StopDis;       //停止距离
	double speed;         //功能4时外部给定传送带速度
	double Readcount;     //模拟量读取间隔（功能2-3时使用）
	double TimeComp;      //跟踪比例修调（功能2-3时使用）
	double ScaleFactor;   //传送带与编码器比例因子：位移/编码器增量（mm/个数）
	char note[32];        //工艺描述
	unsigned char useable;//当前工艺号参数是否可用
	unsigned char if_teached;//判断标定点是否示教

//**********************************************************************//
//																		//
//						  标定需要的变量以及标定结果									//
//																		//
//**********************************************************************//

	double JointA[AXIS_COUNT];   //标定点A对应的机器人关节角
	double referPos[AXIS_COUNT]; //参考点对应的机器人关节角
	double JointB[AXIS_COUNT];   //标定点B对应的机器人关节角
	double JointC[AXIS_COUNT];   //标定点C对应的机器人关节角（针对圆盘跟踪）
	double CircleCenter[3];      //圆盘型传送带中心点坐标
	double ConverMatrix[16];     //存储传送带固定坐标系相对于机器人坐标系之间转换矩阵的数据
	double ConveToRobotMatrix[4][4];	  //传送带固定坐标系相对于机器人坐标系之间的转换矩阵
}TrackCraftData;

typedef struct _TrackData{

//**********************************************************************//
//																		//
//					   跟踪第一级界面：跟踪文件号									//
//																		//
//**********************************************************************//

	int TrackCraftFileNum;	 //PLC中辅助继电器对应的工艺号

	TrackCraftData TrackCraft_Data[TRACK__TECHNICS_COUNT];
	double CurEncoderValue;  //IO检测所获得的当前编码器值
}TrackData;

typedef struct _ObjectInfo{
	int InitEncoderVal;
	int CurEncoderVal;
	double DertTheta;       //单个插补周期圆盘旋转的角度
	double Pos[6];          //当前时刻目标物体的姿态（相对于机器人基坐标系）
	double Rmatrix[4][4];   //当前时刻目标物体的位姿矩阵（相对于机器人基坐标系）
	double ConveVec[4];     //当前时刻目标物体的位置坐标（相对于传送带固定坐标系）
	double ConveDertVec[4];      //XYZ方向上的位移增量（相对于传送带固定坐标系）
	double RobotDertVec[4];      //XYZ方向上的位移增量（相对于机器人基坐标系）
	double ConveSpeedVec[4];     //目标物体的速度矢量（相对于传送带固定坐标系）
	double RobotSpeedVec[4];     //目标物体的速度矢量（相对于机器人基坐标系）
}ObjectInfo;

typedef struct _TVPAccDecPara{
	double length;            //当前长度（mm）
	double CurAcc;            //当前加速度（mm/s2）
	double CurSpeed;          //当前速度（mm/s）
	double MaxSpeed;          //最大速度（mm/s）
	double MaxAcc;            //最大加速度（mm/s2）
	double deltaDist;         //当前位移增量（mm）
	double period;            //单位采样周期时间（s）
	TVPsign StatueFlag;       //当前运动状态标志
}TVPAccDecPara;

typedef struct _EncoderValue{
	int initValue;            //初始时刻对应的编码器值
	int endValue;             //终止时刻对应的编码器值
	int EncoderCount;         //记录编码器读取次数
}EncoderValue;

typedef struct _TrackObjectData
{
	float X;              //目标物体在传送带固定坐标系X轴的坐标值（mm）
	float Y;              //目标物体在传送带固定坐标系Y轴的坐标值（mm）
	float Z;			  //目标物体在传送带固定坐标系Z轴的坐标值（mm）
	float AngleX;		  //目标物体相对于传送带固定坐标系x轴旋转的角度（rad）
	float AngleY;		  //目标物体相对于传送带固定坐标系y轴旋转的角度（rad）
	float AngleZ;         //目标物体相对于传送带固定坐标系z轴旋转的角度（rad）
	int InitEncoderVal;   //初始时刻的编码器值
	int CurEncoderVal;    //当前时刻的编码器值
}TrackObjectData;

typedef struct _TrackPara{
	TrackCraftData *trackcraft_data;
	ObjectInfo TargetInfo[ObjectNumber];
	TVPAccDecPara movpara;
	EncoderValue EncoRead;
	double ObjectToRobotMatrix[4][4];	  //目标物体坐标系相对于机器人坐标系之间的转换矩阵（跟踪过程中实时变化）
	double ConveSpeed;//传送带运行速度（mm/s）
	double WaitLength;        //从传送带固定坐标系到机器人起始抓取区域的固定距离（mm）
	double LimitLength;       //判断目标物体是否越界的长度（mm）
	int Isbeyond;     //判断目标物体是否超出机器人抓取空间 0：未超出  -1：超出
	int StartTrack;   //是否开始跟踪  0：未启用跟踪功能   1：开始跟踪
	int IsTracked;    //是否跟踪上       0 未跟踪上 1：跟踪上
	int IsGrasped;    //是否已经抓取到物体 0：未抓取到 1：已抓取到
	int RealTimeEncoderVal;//实时编码器值
}TrackPara;

void trackdata_init(void);

int Track_Calibration(TrackCraftData *track_data);
int Read_Conveyor_encoder(int *enccnt);
int wait_cutin(int ison);

int Is_Start_Track(void);
int Is_Tracked(void);
int Is_Grasped(void);
void Track_Position_Correct(TrackObjectData *ObjectData);
void Updata_ObjectToRobot_Matrix(void);
int DeterIsEnterTrackArea(TrackObjectData *ObjectData);
void Updata_track_info(TrackObjectData *ObjectData, ObjectInfo *GraspObject);

int RUN_TRACK_IO_CUTIN(int number);
int GET_TRACK_DATA(int number);
int CLEAR_TRACK_DATA(int number);
int TRACK_START(int number);
int TRACK_END(int number);

#ifdef __cplusplus
}
#endif

#endif
