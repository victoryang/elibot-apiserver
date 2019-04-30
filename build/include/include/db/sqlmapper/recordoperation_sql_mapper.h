// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_RECORDOPERATION_SQL_MAPPER_H
#define MCSERVER_RECORDOPERATION_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有用户操作记录数据
#define ELIBOT_RECORD_GET_ALL "elibot.record.getAll"

//插入操作记录
#define ELIBOT_RECORD_INSERT "elibot.record.insert"

// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_record_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_RECORDOPERATION_SQL_MAPPER_H
