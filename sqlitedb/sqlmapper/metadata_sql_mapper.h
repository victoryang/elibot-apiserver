// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_METADATA_SQL_MAPPER_H
#define MCSERVER_METADATA_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有元数据列表
#define ELIBOT_METADATA_GET_ALL "elibot.metadata.getAll"

// 获取元数据映射器
// @id:映射器标识
// @dr:元数据SQL映射器句柄
sql_mapper *get_metadata_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_METADATA_SQL_MAPPER_H
