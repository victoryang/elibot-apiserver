// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef SQLITEDB_DB_QUERY_H
#define SQLITEDB_DB_QUERY_H

#include "db_context.h"
#include "cJSON.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 查询类型处理模式
enum {
    // 标准模式
    DB_QUERY_MODE_STANDARD = (1 << 0),
    // 自定义模式
    DB_QUERY_MODE_CUSTOM = (1 << 1),
    // 单对象模式
    DB_QUERY_MODE_CUSTOM_OBJECT = DB_QUERY_MODE_CUSTOM | (1 << 2),
    // 完全自定义
    DB_QUERY_MODE_CUSTOM_QUERY = DB_QUERY_MODE_CUSTOM | (1 << 3),
};

// 数据库查询请求选项
typedef struct db_query_req_option db_query_req_option;
struct db_query_req_option {
    // 类型处理模式
    int32_t type_handle_mode;
};

// 数据库查询请求分页
typedef struct db_query_req_page db_query_req_page;
struct db_query_req_page {
    // 查询起始页位置
    int16_t page_start;

    // 查询页大小
    int16_t page_size;
};

// 数据库查询请求参数
typedef struct db_query_req_parameter db_query_req_parameter;
struct db_query_req_parameter {
    // 参数
    sql_parameter *params;
    // 参数大小
    int16_t param_size;
};

// 数据库查询请求
typedef struct db_query_req db_query_req;
struct db_query_req {
    // 待查询的标识
    char *query_id;
    // 链接字符串
    char *conn_str;
    // 查询选项
    db_query_req_option *option;
    // 查询请求参数
    db_query_req_parameter *parameter;
    // 分页
    db_query_req_page *page;
};

// 查询请求
// @req:请求对象句柄
// @return:查询结果，以Json方式输出
// json格式：
// {
//      "headers":[
//          "id","para_name"
//      ],
//      "values":[
//          1512,"speed"
//      ],
//      "totalPageSize":32
// }
cJSON* db_query(db_query_req *req);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //SQLITEDB_DB_QUERY_H
