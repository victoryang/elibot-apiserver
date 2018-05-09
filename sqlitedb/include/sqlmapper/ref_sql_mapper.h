// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_REF_SQL_MAPPER_H
#define MCSERVER_REF_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有引用列表
#define ELIBOT_REF_GET_ALL "elibot.ref.getAll"

// 获取参数映射器
// @id:映射器标识
// @dr:参数SQL映射器
sql_mapper *get_ref_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //MCSERVER_REF_SQL_MAPPER_H
