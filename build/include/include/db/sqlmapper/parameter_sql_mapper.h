// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef SQLITEDB_PARAMETER_SQL_MAPPER_H
#define SQLITEDB_PARAMETER_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取指定分组可用的参数列表SQL映射标识
#define ELIBOT_PARAMS_GET_VALID_PARAMS_BY_GROUP "elibot.params.getValidParamsByGroup"

// 获取指定可用的参数QL映射标识
#define ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID "elibot.params.getValidParamById"

// 获取参数列表SQL映射标识
#define ELIBOT_PARAMS_GET_PARAMS "elibot.params.getParams"

// 更新参数SQL映射标识
#define ELIBOT_PARAMS_UPDATE_PARAM "elibot.params.updateParam"

// 获取所有枚举列表
#define ELIBOT_ENUM_GET_ALL "elibot.enum.getAll"

// 获取所有引用列表
#define ELIBOT_REF_GET_ALL "elibot.ref.getAll"

// 获取参数映射器
// @id:映射器标识
// @dr:参数SQL映射器
sql_mapper *get_params_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //SQLITEDB_PARAMETER_SQL_MAPPER_H
