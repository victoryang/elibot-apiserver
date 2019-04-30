// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_COMMENTS_SQL_MAPPER_H
#define MCSERVER_COMMENTS_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

//修改注释
#define ELIBOT_MODIFY_COMMENTS "elibot.modify.comments"

//获取注释
#define ELIBOT_GET_COMMENTS "elibot.get.comments"


// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_comments_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */



#endif //MCSERVER_COMMENTS_SQL_MAPPER_H
