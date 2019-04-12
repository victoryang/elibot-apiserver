// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_INSTALLSOLUTION_SQL_MAPPER_H
#define MCSERVER_INSTALLSOLUTION_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

//新建方案
#define ELIBOT_SOLUTION_NEW_BY_NAME "elibot.solution.newByName"

//删除方案
#define ELIBOT_SOLUTION_DEL_BY_NAME "elibot.solution.deleteByName"

//获取所有方案信息(非数据)
#define ELIBOT_SOLUTION_GET_ALL_SOLUTION_INFO "elibot.solution.getAllSolutionInfo"

// 设置为当前方案
#define ELIBOT_SOLUTION_ENABLE_SOLUTION_BY_NAME  "elibot.solution.enableSolutionByName"

// 更新方案数据
#define ELIBOT_SOLUTION_REFRESH_CURRENT_SOLUTION_DATA "elibot.solution.refreshCurrentSolutionData"

// 获取射器
// @id:映射器标识
// @dr:SQL映射器
sql_mapper *get_solution_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //MCSERVER_INSTALLSOLUTION_SQL_MAPPER_H
