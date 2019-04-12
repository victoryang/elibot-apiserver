#ifndef _VISIONCRAFT_H_
#define _VISIONCRAFT_H_

#include <ext_device/vision.h>

#ifdef __cplusplus
extern "C" {
#endif

#define VISION__TECHNICS_COUNT    8
#define VISION_GPNum1      52
#define VISION_GPNum2      53


//视觉触发方式
typedef enum _VisualTriggerMode
{
	TIME_TRIGGER = 0,
	DIST_TRIGGER,
	CODE_TRIGGER
}VisualTriggerMode;

//采用的视觉模式
typedef enum _VisionMode
{
	TwoD_VisionMode = 0,
	ThreeD_VisionMode
}VisionMode;

//手抓抓取方向
typedef enum _GraspDirection
{
	SameDirect = 0,
	Reserve
}GraspDirection;

typedef struct _VisionTrackData{

//**********************************************************************//
//																		//
//					   视觉第二级界面：视觉配置参数									//
//							  用户输入参数									//
//																		//
//**********************************************************************//

	int VisionSysmId;		  //视觉系统厂家代码
	int TrackMode;			  //0: 无 视觉系统不使用跟踪功能，定点抓取 				   1：有 使用跟踪功能，传送带运动中抓取物体
	int TrackFileNum;		  //跟踪文件号
	VisualTriggerMode Trigger;//0：时间触发 1：距离触发 2：指令触发
	int CommuContentSign;	  //本参数设置系统与相机通讯时，系统需要接收的数据个数标志
	int WorkAreaCoinSing;	  //0：重合 机器人的工作区域可以完全覆盖视觉的视场区域							1：不重合 视觉的视场区域无法与机器人的工作区域重合
	VisionMode VisionType;	  //采用的视觉模式 0：2D 1：3D

	double CommuTimeout;	  //系统执行触发指令到得到视觉数据之间的间隔
	double TriggerInterval;   //用于设置触发条件中选择的时间（ms）或距离的具体间隔（mm）
	double OverlapDist; 	  //视觉数据重叠判断距离（mm）

	char note[32];		  //工艺描述
	unsigned char useable;//当前工艺号参数是否可用

//**********************************************************************//
//																		//
//						 视觉第三级界面：调试参数									//
//							  示教器显示参数 								    //
//																		//
//**********************************************************************//

	char IP[16];			  //IP地址
	int port;				  //端口号
	int VisionEncoVal;		  //视觉侧编码器值
	int RobEncoVal; 		  //机器侧编码器值
	double PixelRatioX;       //像素与实际距离的比例（一个像素对应毫米），X方向
	double PixelRatioY;       //像素与实际距离的比例（一个像素对应毫米），Y方向
	int isConnected;//判断与相机的通讯连接状况
	int isCalibrated;//是否已经获得标定结果

//**********************************************************************//
//																		//
//						 视觉第四级界面：测试参数									//
//					记录参考点、旋转方向及校验相关设置 									//
//																		//
//**********************************************************************//

	GraspDirection GraspDirect; 		//手抓抓取方向

	double ReferRobJointVal[AXIS_COUNT];//参考点机器人侧关节角
	double RecVisionData[6];			//参考点相机侧物体位姿信息
	double RobotPos[6]; 				//试运行下，机器人实际位姿
	double RobotJointVal[AXIS_COUNT];	//试运行下，机器人实际关节角
	double HeightZ; 					//实际物体在相机侧坐标系下的高度（mm）

	int ReferVisionEncoVal;  //参考点相机侧编码器值
	int ReferRobotEncoVal;	 //参考点机器人侧编码器值
	int IsRobotRecoded;//判断机器人侧点是否记录
	int IsVisionRecoded;//判断机器人侧点是否记录
	int if_moved;//判断是否可以使用<试运行该点>按钮

//**********************************************************************//
//																		//
//								标定结果									//
//																		//
//**********************************************************************//

	double ConveToRobotMatrix[4][4];	  //传送带固定坐标系相对于机器人坐标系之间的转换矩阵

}VisionCraftData;

typedef struct _VisionData{

//*****************视觉第一级界面：视觉文件号***************************

	int VisionCraftFileNum;   //视觉工艺文件号，表示一款相机所设定相关参数及工艺设定的一个组合编号

	VisionCraftData VisionCraft_Data[VISION__TECHNICS_COUNT];
}VisionData;

void visiondata_init(void);

int Read_Conveyor_Encoder(int *enccnt, int Conveyor_Encoder_Num);
int Calc_Robot_Pose(int sign);

int VisionTrack_Getdata(int number);

#ifdef __cplusplus
}
#endif

#endif
