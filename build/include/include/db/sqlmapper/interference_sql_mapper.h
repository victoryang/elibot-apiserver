// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_INTERFERENCE_SQL_MAPPER_H
#define MCSERVER_INTERFERENCE_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有的干涉数据
#define ELIBOT_INTERFERENCE_GET_ALL "elibot.interference.getAll"

// 获取指定的干涉数据
#define ELIBOT_INTERFERENCE_GET_BY_NO "elibot.interference.getByNo"

// 更新干涉数据
#define ELIBOT_INTERFERENCE_UPDATE "elibot.interference.update"

// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_interference_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_INTERFERENCE_SQL_MAPPER_H
