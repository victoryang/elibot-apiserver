//This file is generated from para_define.in. don't edit it!

#ifndef PARA_DEFINE_H
#define PARA_DEFINE_H
#ifdef __cplusplus
extern "C" {
#endif
    
#include "../config.h"
#include <stdint.h>

struct rbctrl_para{
	uint32_t speed_max_joint[AXIS_COUNT];	//关节速度上限的值(度/s)
	uint32_t speed_min_joint[AXIS_COUNT];	//关节速度下限的值(度/s)
	uint32_t acc_joint[AXIS_COUNT];	//关节加减速度数值 (度/s2 )
	uint32_t jcc_joint[AXIS_COUNT];	//关节加减速度数值 (度/s3 )
	uint32_t speed_max_line;	//线速度上限的值(mm/s)
	uint32_t speed_min_line;	//直线速度下限的值(mm/s)
	uint32_t acc_max_line;	//直线加减速度上限的值(mm/s2)
	uint32_t jcc_max_line;	//直线加加速上限的值(mm/s3)
	uint32_t speed_max_vr;	//VR 最大角速度(度/s)
	uint32_t speed_min_vr;	//VR 最小角速度(度/s)
	uint32_t acc_vr;	//VR 加速度/减速度(度/s2)
	uint32_t Jcc_vr;	//VR 加加速度/减速度(度/s2)
	uint32_t speed_home_return;	//作业原点回归速度(1%)
	uint32_t speed_modify_play;	//修改再现速度的值
	uint32_t howto_home;	//回零方式 0:开关,1:绝对编码
	uint32_t user_level;	//用户登陆权限
	double first_home_pos[AXIS_COUNT];	//原点位置(度-mm)
	double limit_max_pos[AXIS_COUNT];	//软极限最大位置(度-mm)
	double limit_min_pos[AXIS_COUNT];	//软极限最小位置(度-mm)
	uint32_t robot_rate[AXIS_COUNT];	//倍频比
	uint32_t robot_frenquency_rate[AXIS_COUNT];	//分频比(s:0~5)
	uint32_t pos_zone_level[8];	//PL位置等级区间(毫米)
	uint32_t line_number;	//生产线条数
	uint32_t tool_type;	//抓手类型(单，双抓手)
	uint32_t catch_wait_time;	//ms
	uint32_t release_wait_time;	//ms
	uint32_t has_daowei_sig;	//轨迹转接最大加速度
	uint32_t pl_speed_ratio;	//MOVJ加速时间
	uint32_t motor_type[AXIS_COUNT];	//电机编码器类型
	uint32_t max_motor_speed[AXIS_COUNT];	//电机最大转速
	uint32_t ele_gear[AXIS_COUNT];	//电机转一圈发送脉冲数
	uint32_t ele_gear_back[AXIS_COUNT];	//电机转一圈回馈脉冲数
	uint32_t motor_encode[AXIS_COUNT];	//电机转一圈绝对编码器的脉冲数
	uint32_t axis_direct[AXIS_COUNT];	//指令脉冲与机器人轴运行方向
	uint32_t back_direct[AXIS_COUNT];	//指令脉冲与反馈脉冲方向
	uint32_t encode_direct[AXIS_COUNT];	//指令脉冲与绝对值方向
	uint32_t abz_data_max_diff[AXIS_COUNT];	//发送和接收冲数差
	uint32_t serve_enable_level[AXIS_COUNT];	//伺服报警电平设置1:高0:低有效
	uint32_t lcd_backlight[2];	//LCD背光亮度0-100。index=0为全开亮度，1为半关闭亮度
	uint32_t lcd_time[2];	//LCD关屏时间senconds，0为永不关闭。index=0为半关闭时间，1为全关闭时间
	uint32_t sys_technics;	//工艺选择 
	double dh_stru[12];	//DH参数
	double dh_gear_ratio[AXIS_COUNT];	//齿轮比
	double dh_ouhe_rate[5];	//耦合比
	uint32_t robot_type;	//机器人类型
	double Minimum_AccTime[AXIS_COUNT];	//最小加速时间ms
	double Maximum_AccTime[AXIS_COUNT];	//最大加速时间ms
	double Time_rank[AXIS_COUNT];	//小距离速度加速度等级%
	double Motion_Smooth[AXIS_COUNT];	//运动平滑度
	double Deceleration_Factor;	//高速下减速倍速
	uint32_t collision_enabled;	//碰撞检测使能开关, 0 关　1 开
	uint32_t collision_sensitivity;	//碰撞检测灵敏度, 有效范围 10%~100%
	double collision_inverse_time;	//碰撞逆向运行时间, 范围 1-5秒
};

struct rbctrl_para_item{
	char* const name;
	char* const type_name;
	unsigned int offset;
	const void *min_data;
	const void *max_data;
};

#define PARAMETER_COUNT	(parameter_count)

struct rbctrl_para_item* get_para_item_info_index(unsigned int index);
struct rbctrl_para_item* get_para_item_info_name(const char *name);
void rbctrl_para_set_default(struct rbctrl_para *para);
extern unsigned int parameter_count;

#ifdef __cplusplus
}
#endif

#endif

