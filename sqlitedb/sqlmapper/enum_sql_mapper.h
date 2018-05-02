// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_ENUM_SQL_MAPPER_H
#define MCSERVER_ENUM_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有枚举列表
#define ELIBOT_ENUM_GET_ALL "elibot.enum.getAll"

// 获取参数映射器
// @id:映射器标识
// @dr:参数SQL映射器
sql_mapper *get_enum_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_ENUM_SQL_MAPPER_H
