// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_DYNAMIC_PARAM_H
#define MCSERVER_DYNAMIC_PARAM_H

#define DRICTION_PARAM_COUNT 3

// 动力学参数
typedef struct inverse_dynamic_param inverse_dynamic_param;
struct inverse_dynamic_param {
    // 拖动参数
    unsigned int drap_speed;            // 拖动速度 done

    //摩擦参数
    double frictiondata[AXIS_COUNT][DRICTION_PARAM_COUNT];

    // 记录轨迹
    double tr_sampling_period;          // 记录轨迹采样周期
    double tr_sampling_time;            // 记录轨迹采样时间
    unsigned char track_play_mode;      // 轨迹再现模式, 0 路径轨迹再现　1 时间轨迹再现
#if 0
    unsigned char collision_enabled;    // 碰撞检测使能开关, 0 关　1 开 done
    unsigned int collision_sensitivity; // 碰撞检测灵敏度, 有效范围 10%~100% done
    double collision_inverse_time;      // 碰撞逆向运行时间, 范围 1-5秒
#endif
};

#endif //MCSERVER_DYNAMIC_PARAM_H
