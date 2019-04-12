// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn


#ifndef MCSERVER_DYNAMIC_SQL_MAPPER_H
#define MCSERVER_DYNAMIC_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有的动力学数据
#define ELIBOT_DYNAMIC_GET_ALL "elibot.dynamic.getAll"

// 获取指定的干涉数据
#define ELIBOT_DYNAMIC_GET_BY_ID "elibot.dynamic.getById"

// 更新干涉数据
#define ELIBOT_DYNAMIC_UPDATE "elibot.dynamic.update"

// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_dynamic_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //MCSERVER_DYNAMIC_SQL_MAPPER_H
