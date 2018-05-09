// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_ZEROPOINT_SQL_MAPPER_H
#define MCSERVER_ZEROPOINT_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有零点相关的参数
#define ELIBOT_ZEROPOINT_GET_ALL "elibot.zeropoint.getAll"

// 更新零点参数
#define ELIBOT_ZEROPOINT_UPDATE "elibot.zeropoint.update"

// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_zeropoint_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_ZEROPOINT_SQL_MAPPER_H
