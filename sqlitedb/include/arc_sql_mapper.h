// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn


#ifndef MCSERVER_ARC_SQL_MAPPER_H
#define MCSERVER_ARC_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取焊接参数列表
#define ELIBOT_ARC_GET_PARAMS "elibot.arc.getParams"

// 获取所有焊接参数列表
#define ELIBOT_ARC_GET_ALL_PARAMS "elibot.arc.getAllParams"

// 更新焊接参数
#define ELIBOT_ARC_UPDATE_PARAM  "elibot.arc.updateParam"

// 获取弧焊参数映射器
// @id:映射器标识
// @dr:参数SQL映射器
sql_mapper *get_arc_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_ARC_SQL_MAPPER_H
