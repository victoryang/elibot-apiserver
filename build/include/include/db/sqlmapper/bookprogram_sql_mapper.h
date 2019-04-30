// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_BOOKPROGRAM_SQL_MAPPER_H
#define MCSERVER_BOOKPROGRAM_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有预约相关的参数
#define ELIBOT_BOOKPROGRAM_GET_ALL "elibot.bookprogram.getAll"

// 更新预约相关的参数
#define ELIBOT_BOOKPROGRAM_UPDATE "elibot.bookprogram.update"

// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_bookprogram_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_BOOKPROGRAM_SQL_MAPPER_H
