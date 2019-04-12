// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_EXTAXIS_SQL_MAPPER_H
#define MCSERVER_EXTAXIS_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */


// 获取全部的外部轴数据
#define ELIBOT_EXTAXIS_GET_ALL "elibot.extaxis.getAll"
//获取单独外部轴数
#define ELIBOT_EXTAXIS_GET_FROM_INDEX "elibot.extaxis.from.index"
//更新外部轴参数
#define ELIBOT_EXTAXIS_UPDATE "elibot.extaxis.update"
//更新外部轴参数
#define ELIBOT_EXTAXIS_UPDATE_POS "elibot.extaxis.updatePos"

// 获取参数映射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_extaxis_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //MCSERVER_EXTAXIS_SQL_MAPPER_H
