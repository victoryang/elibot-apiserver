#ifndef _PARAMETERS_H_
#define _PARAMETERS_H_

#include <stdint.h>
#include <common.h>
#include <robot/robotresource.h>

#define SPEED_LEVEL_LOW         0
#define SPEED_LEVEL_MEDIUM      1
#define SPEED_LEVEL_HIGH        2
#define SPEED_LEVEL_HIGHST      3

struct PARA_S2C{
	int32_t cube_inference_check;			//立方体干涉的开启的使能开关
	int32_t cube_softlimit_check[4];		//立方体软极限开启使能开关
	int32_t s_inference_ckeck;				//指定s轴干涉的检查
	int32_t disable_p_move_sig;				//禁止奇异点移动 0 不禁止1 禁止
	int32_t enable_corner_zero_wt_pl;		//在没有开启PL 或CR的情况下，拐角速度为零1 开启 0关闭
	int32_t forwart_step_units;				//前进操作时的步长单位0:在每个命令处停止	1:在移动命令处停止
	int32_t auto_check_run_onf;				//再现模式下(特殊运行)，(检查运行)开关
	int32_t start_addr_pos;					//指定控制电源投入时的地址0:恢复上次关机时的程序和地址。	1:指向主程序的地址 ( “0”行 )
	int32_t autorun_step_units;				//再现运行时，循环模式为"单步"时的步长		0:在每个命令处停止	1:在移动命令处停止
	int32_t enable_io_autorun;				//禁止外部启动 1:禁止0:允许
	int32_t enable_io_modchange;				//禁止外部模式转换 1:禁止0:允许
	int32_t enable_io_cyclechang;			//禁止外部循环转换 1:禁止0:允许
	int32_t enable_io_serve_on;				//禁止外部伺服上电 1(01):只接受示教盒
	                           							 //							2(10):只接受外部
														//							3(11)两者都行
	int32_t auto_low_start_onf;				//再现模式下，(低速启动)开关
	int32_t auto_dry_start_onf;				//再现模式下，(空运行)开关
	int32_t auto_speed_limit_onf;			//再现模式下，(限速运行)开关
	int32_t auto_machine_lock_onf;			//再现模式下，(机械锁运行)开关
	int32_t deadpause_cursor_control;		//急停光标前进控制功能
	int32_t deadpuase_smooth_movc;			//急停时光标指向平滑完成位置指定,单位:%
	int32_t remote_init_cycle_mode;			//远程初始循环模式
	int32_t power_launch_init_cycle_mode;	//电源投入时循环模式
	int32_t lacal_init_cycle_mode;			//本地初始循环模式	
};

struct PARA_S3C{
	int32_t cube_soft_limit_data[24];		//立方体软极限0~23 单位:毫米
	int32_t axis_interfrence[AXIS_COUNT];	//指定S 轴干涉	单位:1%度
	int32_t solid_interfrence_max[3];		//立体干涉区最大值 单位:毫米
	int32_t solid_interfrence_min[3];		//立体干涉区最小值x:0-2 单位:毫米
};


//位置等级区间(s) s:0-8	 			单位:1/1000毫米
#define PARA_POS_ZONE_LEVEL(s)		robot_para_data->data.pos_zone_level[s]	//位置等级区间(s) s:0-7	 			POSITIONING ZONE

//作业原点回归速度					单位:1%
#define PARA_SPEED_HOME_POS_RETURN	robot_para_data->data.speed_home_return	//作业原点回归速度				WORK HOME POSITION RETURN SPEED

//关节速度上限的值0~7				单位：度/s
#define PARA_SPEED_MAX_JOINT(x)	    robot_para_data->data.speed_max_joint[x]	//关节速度上限的值 		

//关节速度下限的值0~7				单位：0.01度/s
#define PARA_SPEED_MIN_JOINT(x)	    robot_para_data->data.speed_min_joint[x]	//关节速度上限的值 		

//关节加速度和减速度数值 ACC && DEC  s:0-5		单位：1度/s2                              
#define PARA_SPEED_CC(s)			robot_para_data->data.acc_joint[s]//关节加速度和减速度数值 ACC && DEC 
// 1degree/s3
#define PARA_SPEED_JCC(s)			robot_para_data->data.jcc_joint[s]

//控制点线速度上限的值				单位：0.1mm/s
#define PARA_SPEED_MAX_LINER	    robot_para_data->data.speed_max_line	//控制点线速度上限的值 		

//控制点线加减速度上限的值			单位：0.01mm/s2
#define PARA_SPEED_MAX_LINER_CC		robot_para_data->data.acc_max_line	//控制点线加减速度上限的值	

//控制点线加加速度上限的值			单位：0.01mm/s3
#define PARA_SPEED_MAX_LINER_JCC		robot_para_data->data.jcc_max_line	

//控制点线速度下限的值				单位：0.1mm/s
#define PARA_SPEED_MIN_LINER	    robot_para_data->data.speed_min_line							

//修改再现速度的值
#define PARA_MODIFY_PLAY_SPEED		robot_para_data->data.speed_modify_play

//第一原点位置(s:0-5)				单位:0.01度
#define PARA_FRI_HOME_POS(s)		robot_para_data->data.first_home_pos[s]
//VR 最大角速度					单位:0.01度/s
#define PARA_MAX_VR_SPEED			robot_para_data->data.speed_max_vr
//VR 最小角速度					单位:0.01度/s
#define PARA_MIN_VR_SPEED			robot_para_data->data.speed_min_vr
//VR 加速度/减速度				单位:0.01度/s2
#define PARA_VR_CC					robot_para_data->data.acc_vr
#define PARA_VR_JCC					robot_para_data->data.Jcc_vr


//软极限最大位置(s:0-7)				单位:0.01度
#define PARA_MAX_LIMIT_POS(s)		robot_para_data->data.limit_max_pos[s]
//软极限最小位置(s:0-7)				单位:0.01度
#define PARA_MIN_LIMIT_POS(s)		robot_para_data->data.limit_min_pos[s]
//倍频比(s:0~5)
#define PARA_ROBOT_RATE(s)			robot_para_data->data.robot_rate[s]	

//分频比(s:0~5)						单位:0.1%
#define PARA_ROBOT_FRENQUENCY_RATE(s)		robot_para_data->data.robot_frenquency_rate[s]	

//伺服报警电平设置1:高0:低有效
#define PARA_SERVE_ENABLE_LEVEL(s)  		robot_para_data->data.serve_enable_level[s]

//生产线条数
#define PARA_LINE_NUMER  		robot_para_data->data.line_number 

//抓手类型
#define PARA_TOOL_TYPE  		robot_para_data->data.tool_type 

#define MOTOR_MAX_SPEED(x)	robot_para_data->data.max_motor_speed[x]

//抓取等待时间
#define PARA_CATCH_TIME  			robot_para_data->data.catch_wait_time
//放开等待时间
#define PARA_RELEASE_TIME  			robot_para_data->data.release_wait_time

#define PARA_SPEED_RATIO  			robot_para_data->data.pl_speed_ratio

//默认焊接规范
//#define PARA_SPORT_RULE  			robot_para_data->data.[345]
#define PARA_HAS_DAOWEI				robot_para_data->data.has_daowei_sig//是否有到位信号 1:有,0:无
#define PARA_HOW_TOHOME				robot_para_data->data.howto_home//回零方式 0:开关,1:绝对编码

//存绝对原点的伺服脉冲数S>=0,S<=5
#define ABZ_DATA_DIFF(s)				robot_para_data->data.abz_data_max_diff[s]
#define PARA_SYS_TECHNIC	robot_para_data->data.sys_technics

#define PARA_AXIS_DIRECT(a)		robot_para_data->data.axis_direct[a]
#define PARA_BACK_DIRECT(a)		robot_para_data->data.back_direct[a]
#define PARA_ENCODE_DIRECT(a)	robot_para_data->data.encode_direct[a]

//操作相关--------------------------------------------------
//立方体干涉的开启的使能开关
#define PARA_CUBE_INFERENCE_CHECK			ps2c.cube_inference_check
//立方体软极限的开启的使能开关0~3
#define PARA_CUBE_SOFTLIMIT_CHECK(x)			ps2c.cube_softlimit_check[x]	
//指定s轴干涉的检查
#define PARA_S_INFERENCE_CHECK					ps2c.s_inference_ckeck	
//禁止奇异点移动开启0 不禁止1 禁止
#define DISABLE_SING_P_MOVE					ps2c.disable_p_move_sig	
//在没有开启PL 或CR的情况下，拐角速度为零1 开启 0关闭
#define ENABLE_CORNER_ZERO_WT_PL				ps2c.enable_corner_zero_wt_pl	
//前进操作时的步长单位		0:在每个命令处停止	1:在移动命令处停止
#define PARA_OPEAR_UNITS_FORWARD 				ps2c.forwart_step_units	//EXECUTION UNITS AT “FORWARD” OPERATION

//S2C147  为开机默认的安全模式，0，1，2，三级

//      【再现模式】下的【特殊运行】，【检查运行】（开关）
#define PARA_OPEAR_CHECK_RUN_ONF    			ps2c.auto_check_run_onf

//指定控制电源投入时的地址
#define PARA_START_POS							ps2c.start_addr_pos	//0:控制电源投入时，恢复上次关机时的程序和地址。	1:指向主程序的地址 ( “0”行 )


//再现运行时，循环模式为"单步"时的步长      0:在每个命令处停止  1:在移动命令处停止
#define PARA_OPEAR_SETP_PLAY        			ps2c.autorun_step_units

//禁止外部启动 1:禁止0:允许
#define PARA_IO_AUTORUN				 			ps2c.enable_io_autorun

//禁止外部模式转换 1:禁止0:允许
#define PARA_IO_MODECHANGE				 		ps2c.enable_io_modchange

//禁止外部循环转换 1:禁止0:允许
#define PARA_IO_CYCLECHANGE				 		ps2c.enable_io_cyclechang

//禁止外部伺服上电 1(01):只接受示教盒
//							2(10):只接受外部
//							3(11)两者都行
#define PARA_IO_SERVIECE				 		ps2c.enable_io_serve_on 

//(新建)【再现模式】下的【特殊运行】，【低速启动】（开关）
#define PARA_OPEAR_LOW_START_ONF    			ps2c.auto_low_start_onf   

//(新建)【再现模式】下的【特殊运行】，【空运行】（开关）
#define PARA_OPEAR_DRY_RUN_ONF      			ps2c.auto_dry_start_onf

//(新建)【再现模式】下的【特殊运行】，【限速运行】（开关）
#define PARA_OPEAR_SPEED_LIMIT_ONF  			ps2c.auto_speed_limit_onf

//(新建)【再现模式】下的【特殊运行】，【机械锁定运行】（开关）
#define PARA_OPEAR_MACHINE_LOCK_ONF 			ps2c.auto_machine_lock_onf

//急停光标前进控制功能
#define PARA_DEADSTOP_CURSOR_FORWARD_CONTROL 	ps2c.deadpause_cursor_control

//急停时光标指向平滑完成位置指定,单位:%
#define PARA_DEADSTOP_SMOOTH_MOVC				ps2c.deadpuase_smooth_movc

//循环模式相关--------------------------------------------------

//远程初始循环模式
#define PARA_REMOTE_INIT_CYCLE_MODE			ps2c.remote_init_cycle_mode
//电源投入时循环模式
#define PARA_POWER_LAUNCH_INIT_CYCLE_MODE	ps2c.power_launch_init_cycle_mode

//本地初始循环模式
#define PARA_LOCAL_INIT_CYCLE_MODE			ps2c.lacal_init_cycle_mode


//S3C-----------------------------------------------------------------------
//立方体软极限S3C000 至 S3C023:0~23					单位:毫米
#define PARA_CUBE_SOFT_LIMIT(x)				ps3c.cube_soft_limit_data[x]
//指定S 轴干涉0~7									单位:1%度
#define PARA_SAXIS_INTERFERENCE(x)			ps3c.axis_interfrence[x]
//立体干涉区最大值x:0-2 单位:毫米
#define PARA_SOLID_INTERFERENCE_MAX(x)		ps3c.solid_interfrence_max[3]	
//立体干涉区最小值x:0-2 单位:毫米
#define PARA_SOLID_INTERFERENCE_MIN(x)		ps3c.solid_interfrence_min[3]

extern struct para_info *robot_para_data;
extern struct PARA_S2C ps2c;
extern struct PARA_S3C ps3c;

#define GET_ARRAY_INFO_DEFAULT(i, size, d, def)	((unsigned int)(i) >= (size) ? (def) : d(i))
#define GET_ARRAY_INFO(i, size, d)	GET_ARRAY_INFO_DEFAULT(i, size, d, 0)
#define GET_AXIS_INFO(i, d)	GET_ARRAY_INFO(i, AXIS_COUNT, d)
#define GET_AXIS_INFO_DEFAULT(i, d, default)	GET_ARRAY_INFO_DEFAULT(i, AXIS_COUNT, d, default)

//!参数初始化
void init_default_para(void);
void parameter_init(void* ram);

//再现模式下是否低速启动
//return int 0：无效  1：有效
#define isLowSpeedStart_Play()	PARA_OPEAR_LOW_START_ONF

//重置低速启动 
#define resetLowSpeedStart()	PARA_OPEAR_LOW_START_ONF = 0

//再现模式下是否限速运行
//return int 0：无效  1：有效
#define isSpeedLimit_Play()		PARA_OPEAR_SPEED_LIMIT_ONF

//再现模式下是否空运行
//return int 0：无效  1：有效
#define isDryRun_Play()			PARA_OPEAR_DRY_RUN_ONF

//再现模式下是否机械检查运行
//return int 0：无效  1：有效
#define isMachineLock_Play()	PARA_OPEAR_MACHINE_LOCK_ONF

//再现模式下是否检查运行
//return int 0：无效  1：有效
#define isChechRun_Play()	PARA_OPEAR_CHECK_RUN_ONF


//得到（控制点）线速度上限值
//return int 单位：1.0 mm/s or um/ms
#define getMaxLineSpeed()	((double)PARA_SPEED_MAX_LINER)

//得到（控制点）线速度下限值
//return int 单位：1.0 mm/s
#define getMinLineSpeed()	((double)PARA_SPEED_MIN_LINER)

//得到再现运行时，循环模式为"单步"时的步长
//return int  0:在每个命令处停止  1:在移动命令处停止
#define getSetpInPlay()		PARA_OPEAR_SETP_PLAY

//0:no use motor  1:abs motor  2:normal motor
#define MOTOR_TYPE(n,m)		(n<m?(robot_para_data->data.motor_type[n]):(robot_para_data->data.motor_type[m-1]))
#define GetMotorType(n)		MOTOR_TYPE(n,AXIS_COUNT)

//1~int.MAX default 131072
#define MOTOR_ENCODE(n,m)		(n<m?(robot_para_data->data.motor_encode[n]):(robot_para_data->data.motor_encode[m-1]))
#define GetMotorEncode(n)		MOTOR_ENCODE(n,AXIS_COUNT)

double getRotationSpeed_CylinderCoord_Sita(int speedlevel);
double getRotationAcc_CylinderCoord_Sita(int speedlevel);

//得到修改的再现速度1%-150%
//return int 单位：1%
#define GetModifyPlaySpeed()		((double)BOUND(1, PARA_MODIFY_PLAY_SPEED, 15000)/10000.0)

//得到第一原点位置
// @param pos       返回的位置数据
void getFirstHomePos(double pos[]);
void setFirstHomePos(double pos[]);

//获取软极限最大位置角度数据
//param pos    	单位:1.0 度
void getSoftMaxLimitPos(double pos[]);

//获取软极限最小位置脉冲数据
//param pos    	单位:1.0 度
void  getSoftMinLimitPos(double pos[]);

//获取最小关节速度
//return double  速度: 度/s
#define getMinJointSpeed(index)	((double)GET_AXIS_INFO(index, PARA_SPEED_MIN_JOINT) )

//获取插补算法的版本号
//return int  版本号: 
#define getInterpAlamVersion()	PARA_INTERP_ALAM_VERSION

//获取分割插补算法大插补周期
//return int  大插补周期: 
#define getInterpBigPeriod()	PARA_INTERP_BIG_PERIOD

//获取远程初始循环模式
//return int  循环模式0: 保持设定	1:单步	2:连续	3:单循环
#define getRemoteCycleMode()	PARA_REMOTE_INIT_CYCLE_MODE

//获取本地初始循环模式
//return int  循环模式0: 保持设定	1:单步	2:连续	3:单循环
#define getLocalCycleMode()		PARA_LOCAL_INIT_CYCLE_MODE

//获取电源投入时初始循环模式
//return int  循环模式0: 保持设定	1:单步	2:连续	3:单循环
#define getPowerLaunchCycleMode()	PARA_POWER_LAUNCH_INIT_CYCLE_MODE

//获取示教初始循环模式
//return int  循环模式0: 保持设定	1:单步	2:连续	3:单循环
#define getTeachCycleMode()	PARA_TEACH_INIT_CYCLE_MODE


//获取再现初始循环模式
//return int  循环模式0: 保持设定	1:单步	2:连续	3:单循环
#define getPlayCycleMode()	PARA_PLAY_INIT_CYCLE_MODE

//获得位置等级范围
//return double 等级范围单位:us
#define getPosZoneLevel(plvalue)	((double)GET_ARRAY_INFO(plvalue, 8, PARA_POS_ZONE_LEVEL))

/*
得到手动动作线速度
param speedlevel 速度等级
return double 		单位： 100.0%
*/
double getJogLinkSpeed(int speedlevel);

//得到控制点线加减速度上限的值
//return double 		单位：1.0 mm/s2
#define getMaxLineCC(void)	(((double)PARA_SPEED_MAX_LINER_CC))
#define getMaxLineJCC(void)	(((double)PARA_SPEED_MAX_LINER_JCC))

//焊接使能
//return double  单位: 秒
#define GetSpeedRatio()		(PARA_SPEED_RATIO*0.001) //ms 2 s

//有无到位检测信号
//return double  单位: 秒
#define GetTransAcc()		PARA_HAS_DAOWEI

//回零方式选择 0，外部开关回零 1，绝对编码回零
//return double  单位: 秒
#define ToHomeChoose()	PARA_HOW_TOHOME

//得到伺服有效电平
//return 1:高有效,0:低有效
#define GetServeEnableLevel(i) PARA_SERVE_ENABLE_LEVEL(i)

//抓手类型
//return 0,1(默认):单抓手,2:双抓手3:3双抓手
#define GetToolType()	PARA_TOOL_TYPE

//得到立方体软极限
//param index 标示:(0-23)
//return int 		1.0 毫米
#define getCubeSoftLimit(index)		GET_ARRAY_INFO(index, 24, PARA_CUBE_SOFT_LIMIT)

//获取最大VR 速度
//return double :VR 速度 单位:1.0 度
#define getMaxVRSpeed()	((double)PARA_MAX_VR_SPEED)

//获取最小VR 速度
//return double :VR 速度 单位:1.0 度
#define getMinVRSpeed()	((double)PARA_MIN_VR_SPEED)

//获取VR 加速度和减速速度
//return double :VR 加速度/减速度 单位:1.0 度/S2

#define getVRCC()			((double)PARA_VR_CC)
#define getVRJCC()		((double)PARA_VR_JCC)

#define SENDPLUSE_PER_CIRCLE(n)	robot_para_data->data.ele_gear[n]
#define BACKPLUSE_PER_CIRCLE(n)	robot_para_data->data.ele_gear_back[n]

#define GetSendPlusePerCircle(n)	SENDPLUSE_PER_CIRCLE(n)
#define GetBackPlusePerCircle(n)	BACKPLUSE_PER_CIRCLE(n)

//获得立体干涉区最大值(单位:毫米)
//param double  干涉最大值
void getMaxInterferenceValue(double interference[]);

//获得立体干涉区最小值(单位:毫米)
//param double  干涉最小值
void getMinInterferenceValue(double interference[]);

//指定控制电源投入时的地址
//0:控制电源投入时，恢复上次关机时的程序和地址。	1:指向主程序的地址
#define startPosIsMainJob()	PARA_START_POS

//获得急停时光标指向平滑完成位置指定参数
//return double :指定平滑位置区域
#define getParaDeadStopSmoothMovc()	((double)BOUND(20, PARA_DEADSTOP_SMOOTH_MOVC, 100)/100.0)

//得到外部启动使能变量值
#define getIO_AutoRun()		PARA_IO_AUTORUN

//设置外部启动使能变量值
#define setIO_AutoRun(flag)	PARA_IO_AUTORUN=flag

//得到外部模式转换使能变量值
#define getIO_ModeChange() PARA_IO_MODECHANGE

//设置外部模式转换使能变量值
#define setIO_ModeChange(flag)	PARA_IO_MODECHANGE=flag

//得到外部循环模式转换使能变量值
#define getIO_CycleModeChange()	PARA_IO_CYCLECHANGE

//设置外部循环模式转换使能变量值
#define setIO_CycleModeChange(flag)	PARA_IO_CYCLECHANGE=flag

//得到外部伺服上电变量值
#define getIO_Serviece()	PARA_IO_SERVIECE

//设置外部伺服上电变量值
#define setIO_Serviece(flag)	PARA_IO_SERVIECE=flag;

//获取立体软极限检测标志
//return int :0:代表不检测1代表检测
#define getCubeSoftLimitCheck(index)	GET_ARRAY_INFO(index, 4, PARA_CUBE_SOFTLIMIT_CHECK)

//获取立体干涉检测标志
//return int :0:代表不检测1代表检测
#define getCubeInterferenceCheck()	(PARA_CUBE_INFERENCE_CHECK==1)

//获取S轴干涉检测标志
//return int :0:代表不检测1代表检测
#define getSAxisInterferenceCheck()	(PARA_S_INFERENCE_CHECK==1)


//获取S轴干涉数据
#define getSAxisInterference(index)	((double)GET_ARRAY_INFO(index, 8, PARA_SAXIS_INTERFERENCE)/100.0)

//获得急停光标前进控制功能
//return int : 0 此控制功能不使用1使用此控制功能
#define getDeadStopCursorForwardControl()	(PARA_DEADSTOP_CURSOR_FORWARD_CONTROL==1)

//禁止奇异点移动
//return int : 0 不禁止奇异点移动1禁止奇异点移动
#define IsProhbitSingPointMove()	(DISABLE_SING_P_MOVE==1)

//没有PL 或CR过度的情况，是否开启拐角速度为零的信号
//return int : 0 禁止1启用
#define IsEnableCornerSpeed()	(ENABLE_CORNER_ZERO_WT_PL==1)
 
//设置各轴脉冲值	
void SetServoData(int ser,unsigned int axis);
void SetAllServoData(int ser[]);

#define GET_PARA_BIT(p,b)	((p>>b)&0x01)

#define getAxisDir(axis)	PARA_AXIS_DIRECT(axis) //GET_PARA_BIT(PARA_MIRROR_SHIFT_SIGN,axis)
#define getEnCodeDir(axis)	PARA_ENCODE_DIRECT(axis) //GET_PARA_BIT(PARA_MIRROR_SHIFT_SIGN,(axis+8))
#define getAbzDir(axis)		PARA_BACK_DIRECT(axis) //GET_PARA_BIT(PARA_MIRROR_SHIFT_SIGN,(axis+16))

//得到各轴脉冲值	
//return int 单位：1脉冲
#define Get_ABZ_MAX_Diff(axis)		GET_AXIS_INFO(axis, ABZ_DATA_DIFF)
#define GetJointDec(s)			GetJointAcc(s)
//作业原点回归速度					单位:1%
#define Get_Return_Home_Speed() (PARA_SPEED_HOME_POS_RETURN/100.0)

double GetMaxJointSpeed(int axis);
double GetJointAcc(int axis);
void set_default_zeropos(int type);
double GetJointJcc(int axis);
//void setModualFunc(int v);

void initServoDriveMode(const char *str_mode);
void setServoDriveMode(int mode, int option);
void setIP(int ip);

#endif
