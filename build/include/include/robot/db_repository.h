// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_DB_REPOSITORY_H
#define MCSERVER_DB_REPOSITORY_H

#include <stdint.h>
#include <stdio.h>

// 数据库链接
#define DB_CONN_STRINGS "db/elibotDB.db"

// 操作状态
#define REP_OK 0
#define REP_ERR -1

// 数据库对象仓储
typedef struct obj_db_repository obj_db_repository;
struct obj_db_repository {
    // 仓储名称
    char *name;
    // 工艺类型
    int64_t tech_type;
    // 对象仓储初始化
    // ＠return:0 表示成功　非0表示初始化失败
    int8_t (*init_func)(void);
};

// 仓储构建宏定义
#define BUILDIN_REPOSITORY(rep, tech)	 static obj_db_repository rep##_buildin_db_repository \
	__attribute__((section("buildin_db_repository"),aligned(sizeof(void*)),__used__))= \
{\
    .name = #rep,\
    .tech_type=tech,\
    .init_func=init_##rep##_repository\
}

// 数据库请求参数
typedef struct db_q_args db_q_args;
struct db_q_args {
    int argc;
    char **argv;
    FILE *pfin;
    FILE *pfout;
};

// 初始化所有仓储
// 初始化条件依据其所支持的工艺
void init_repositories();

#endif //MCSERVER_DB_REPOSITORY_H
