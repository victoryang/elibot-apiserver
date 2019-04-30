// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_TOOLFRAME_SQL_MAPPER_H
#define MCSERVER_TOOLFRAME_SQL_MAPPER_H
#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取工具坐标参数列表
#define ELIBOT_COMMON_GET_TOOLFRAMES "elibot.common.toolframes"
//更新工具坐标参数
#define  ELIBOT_COMMON_UPDATE_TOOLFRAME "elibot.common.update.toolframe"
//获取自动校验点位数据
#define ELIBOT_COMMON_GET_TOOLPOS  "elibot.common.toolpos"
//更新自动校验点位数据
#define  ELIBOT_COMMON_UPDATE_TOOLPOS "elibot.common.update.toolpos"
//获取所有工具坐标状态
#define  ELIBOT_COMMON_GET_TOOLNAMEDATA "elibot.common.toolnamedata"
//获取所有工具数据
#define ELIBOT_COMMON_GET_ALL_TOOLFRAMES  "elibot.common.all.toolframes"
//获取所有工具点位数据
#define ELIBOT_COMMON_GET_ALL_TOOLPOS  "elibot.common.all.toolpos"
//获取所有工具坐标的使能状态
#define ELIBOT_TOOLFRAME_GET_ALL_ENABLE_STATUSES  "elibot.toolframe.getAllEnableStatuses"

// 获取弧焊参数映射器
// @id:映射器标识
// @dr:参数SQL映射器
sql_mapper *get_common_toolframe_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_TOOLFRAME_SQL_MAPPER_H
