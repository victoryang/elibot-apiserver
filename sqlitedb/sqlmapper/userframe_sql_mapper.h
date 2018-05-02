// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_USERFRAME_SQL_MAPPER_H
#define MCSERVER_USERFRAME_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 获取所有的用户坐标数据
#define ELIBOT_USER_FRAME_GET_ALL "elibot.userframe.getAll"

// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_userframe_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //MCSERVER_USERFRAME_SQL_MAPPER_H
