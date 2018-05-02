// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_IO_SQL_MAPPER_H
#define MCSERVER_IO_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 格局分组获取有效的IO列表
#define ELIBOT_IO_GET_VALID_IOS_BY_GROUP "elibot.io.getValidIOsByGroup"

// 获取IO映射器
// @id:映射器标识
// @dr:参数SQL映射器
sql_mapper *get_io_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //MCSERVER_IO_SQL_MAPPER_H
