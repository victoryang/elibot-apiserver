#ifndef CXZK_H
#define CXZK_H

#ifdef __cplusplus
extern "C" {
#endif

#define LASER_ON_OFF_FAILED    -1
#define SEARCH_FAILED          -2
#define SEAM_NOT_FOUND         -3
#define TRACE_FAILED           -4
#define TCP_SEND_ERR           -5
#define TCP_RECV_ERR           -6
#define SOCKET_FAILED          -7
#define CONNECT_ERR            -8

#define LASER_IS_ON             1
#define SEARCHING_SEAM          2
#define SEAM_FOCUSED            3
#define TRACE_KEEP_ON           4
#define TRACE_IS_ON             5
#define TRACE_IS_OFF            6
#define LASER_IS_OFF            7

#define LASER_TRACK_TECHNICS_COUNT    8
#define Calibration_Number    6

typedef struct _WeldTrackData{
	char note[32];
	unsigned char useable;
	int factory;    //激光跟踪器厂家
	int communication_status;    //激光通讯状态
	char IP_address[16];    //IP地址
	char Controller_IPaddress[16];    //控制器IP地址
	int port;    //端口号
	double connection_timeout;    //连接超时（ms）
	double readwrite_timeout;    //读写超时时间（ms）

	double front_distance;    //前置距离（mm）
	int sensitivity;    //灵敏度
	double Xaxis_ratio;    //X轴比例
	double Zaxis_ratio;    //Z轴比例
	unsigned char if_teached[Calibration_Number+1];
	double Xdeviation;    //X轴偏差（mm）
	double Zdeviation;    //Z轴偏差（mm）
	double SensorPoint[4][Calibration_Number];    //激光传感器坐标系下标定点的坐标值
	double RobotJoint[Calibration_Number+1][AXIS_COUNT];    //标定点对应的机器人关节角度值
	double SensorFrame[4][4];    //激光传感器坐标系到机器人末端法兰盘坐标系下的转换矩阵

	int search_type;    //0:NG->OK 1:OK->NG 2:STOP
	double OverlapDistance;    //重叠距离（mm）
	double WeldDetectionOverlap;    //焊缝检测重叠距离（mm）
	double NonWeldDetectionOverlap;    //非焊缝检测重叠距离（mm）
	int JudgeBufferSize;    //用于判断OK和NG的缓冲区大小
	double search_interval;    //搜寻间隔时间（ms）
	double sample_length;    //采样最小间隔（mm）
	double search_speed;    //搜寻速度（mm/s）
	double search_timeout;    //搜寻超时时间（ms）
	double deviation_distance;    //搜寻超偏差距离（mm）

	int XYZ_Correction;    //XYZ方向校正 0：不校正 1：校正
	double X_Shift;    //X向偏移（mm）
	double Y_Shift;    //Y向偏移（mm）
	double Z_Shift;    //Z向偏移（mm）
	int ActionAfterNonDetection;    //检测不到后的运动 0：不处理 1：停止
	double NonDetectionDistance;    //指定检测距离（mm）
	double DataAquisitionLength;    //数据采集长度（mm）
	double DetectErrLimit;    //错误检测限制（mm）

	double feedvector[3];    //焊缝跟踪进给矢量
	double Xvalue;    //激光传感器反馈的X方向值（mm）
	double Zvalue;    //激光传感器反馈的Z方向值（mm）
	int state;    //焊缝跟踪状态
	int isConnected;    //是否连接激光跟踪器 0：未连接 1：连接
	int isOpen;    //是否打开激光跟踪器 0：未打开 1：打开
}WeldTrackData;

typedef struct _LaserTrackData{
	int LaserTrackCraftNum;
	WeldTrackData WeldTrack_Data[LASER_TRACK_TECHNICS_COUNT];
}LaserTrackData;

int built_sensor_frame(WeldTrackData *SensorDataS);
int laser_on(int num, int ison);
int laser_search(int num);
int laser_trackon(int num, int ison);
int is_laser_used(void);
int is_start_searchseam(void);
int is_start_lasertrack(void);

void lasertrackdata_init(void);
int laser_connect_init(void);
int weld_data_init(int num);
int cmd_cxzk_laser_switch(unsigned char is_on);
int cmd_cxzk_trace_switch(unsigned char is_on);
int cmd_cxzk_search(void);
void smtrk_cxzk_exit(void);

#ifdef __cplusplus
}
#endif

#endif
