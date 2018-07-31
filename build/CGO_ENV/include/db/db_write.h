// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef SQLITEDB_DB_WRITE_H
#define SQLITEDB_DB_WRITE_H


#include "db_context.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 数据库写请求参数
typedef struct db_write_req_parameter db_write_req_parameter;
struct db_write_req_parameter {
    // 参数
    sql_parameter *params;
    // 参数大小
    int16_t param_size;
};

// 数据库写请求
typedef struct db_write_req db_write_req;
struct db_write_req {
    // 待查询的标识
    char *query_id;
    // 链接字符串
    char *conn_str;

    // 查询请求参数
    db_write_req_parameter *parameter;
};

// 数据库写请求
// @req:请求对象句柄
// @return:0 表示正确　非0表示错误
int db_write(db_write_req *req);

// 数据库脚本写请求
// @conn_str:联接字符串
// @sql:sql脚本
// @return:0 表示正确　非0表示错误
int db_write_by_sql(const char *conn_str, char *sql);

// 数据库脚本写请求
// @conn_str:联接字符串
// @sql:sql脚本
// @return:0 表示正确　非0表示错误
int db_write_by_file(const char *conn_str, char *script_path_Name);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //SQLITEDB_DB_WRITE_H
