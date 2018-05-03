// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef SQLITEDB_SQL_MAPPER_H
#define SQLITEDB_SQL_MAPPER_H

#include "db_context.h"
#include "cJSON.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// SQL 绑定参数
typedef struct sql_binding_parameter sql_binding_parameter;
struct sql_binding_parameter {
	// 参数名
	char *param_name;
	// 参数类型
	int32_t type;
};

// SQL 映射
typedef struct sql_mapper sql_mapper;
struct sql_mapper {
    // 映射唯一表示
    char *id;
    // SQL模板
    char *sql;

    // SQL 参数
    sql_binding_parameter *binding_params;
    // SQL 参数个数
    int32_t binding_params_size;

    // 对象类型处理器
    // @parent_object:父JSON对象
    // @dr:数据读取器句柄
    // @return:单JSON对象
    void (*object_type_handler)(cJSON *parent_object, sql_data_reader *dr);

    // 查询请求类型处理器
    // @dr:数据读取器句柄
    // @return:单JSON对象
    cJSON * (*query_type_handler)(sql_data_reader *dr);
};

// 注册SQL映射
// @q:SQL映射
uint8_t register_sql_mapper(sql_mapper *mapper);

// 获取SQL映射对象句柄
// @mapper_id:SQL映射标识
// @return:SQL映射对象句柄
sql_mapper* find_sql_mapper(const char *mapper_id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif/* SQLITEDB_SQL_MAPPER_H */
