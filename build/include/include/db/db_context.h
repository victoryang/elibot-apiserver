// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef SQLITEDB_DB_CONTEXT_H
#define SQLITEDB_DB_CONTEXT_H

#include <stdint.h>

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

#define DB_OK 0
#define DB_ERR -1

// 隔离级别
enum { 
    // 延迟锁
    DB_DEFERRED =0, 
    // 立即锁
    DB_IMMEDIATE = 1,
    // 排它锁 
    DB_EXCLUSIVE = 2
};

// 数据库数据类型
enum { 
    // Int32
    DB_TYPE_INT32 = 1, 
    // Int64
    DB_TYPE_INT64 = 2,
    // Double 
    DB_TYPE_DOUBLE = 3,
    // String
    DB_TYPE_TEXT = 4,
    // NULL
    DB_TYPE_NULL = 5,
};

// 数据类型信息
enum {
    DATA_UINT8 = 1,
    DATA_UINT16 = 2,
    DATA_UINT32 = 3,
    DATA_UINT64 = 4,
    DATA_INT8 = 5,
    DATA_INT16 = 6,
    DATA_INT32 = 7,
    DATA_INT64 = 8,
    DATA_FLOAT = 9,
    DATA_DOUBLE = 10,
    DATA_STRING = 11
};

// SQL参数
typedef struct sql_parameter sql_parameter;
struct sql_parameter {
    // 参数名
    char *name;

    // 参数值
    union {
        // 整数值
        int32_t int_value;
        // 长整数值
        int64_t int64_value;
        // 浮点数值
        double double_value;
        // 字符串值
        char *string_value;
    } value;

    // 参数类型
    int32_t type;
};

// 数据库连接
typedef struct sql_conn sql_conn;
struct sql_conn {
    // 链接字符串
    char conn_str[256];  
    // 数据库指针
    void *db;

    // 关闭链接
    // @db:数据库句柄
    void (*close)(sql_conn *conn);
};

// SQL 事务
typedef struct sql_trans sql_trans;


// SQL数据读取器
typedef struct sql_data_reader sql_data_reader;

// 数据库上下文
typedef struct db_context db_context;
struct db_context {
    // 数据库链接
    sql_conn conn;
    // SQL事务
    sql_trans *trans;

    // 错误信息
    char *err_msg;

    // 打开数据库
    // @ctx: 数据库上下文句柄
    int (*open)(db_context *ctx);

    // 开启事务
    // @ctx: 数据库上下文句柄
    // @iso_level:事务隔离级别
    int (*begin_trans)(db_context *ctx, sql_trans *new_trans, int iso_level);

    // 查询操作
    // @ctx: 数据库上下文句柄
    // @dr:输出，数据读取器句柄
    // @sql: SQL查询脚本
    // @parameters: 参数列表
    // @return:
    int (*query)(db_context *ctx, sql_data_reader *dr, const char *sql);

    // 查询操作
    // @ctx: 数据库上下文句柄
    // @dr:输出，数据读取器句柄
    // @sql: SQL查询脚本
    // @parameters: 参数列表
    // @return:
    int (*query_by_params)(db_context *ctx, sql_data_reader *dr, const char *sql, sql_parameter parameters[], int param_size);
    
    // 执行语句操作
    // @ctx: 数据库上下文句柄
    // @sql: SQL查询脚本
    // @return 
    int (*execute)(db_context *ctx, const char *sql);

    // 执行语句操作
    // @ctx: 数据库上下文句柄
    // @sql: SQL查询脚本
    // @parameters: 参数列表
    // @param_size: 参数大小
    // @return 
    int (*execute_by_params)(db_context *ctx, const char *sql, sql_parameter parameters[], int param_size);

    // 关闭数据库上下文
    // @ctx: 数据库上下文句柄
    void (*close)(db_context *ctx);

    // 获取错误码
    // @ctx: 数据库上下文句柄
    int (*err_code)(db_context *ctx);

    // 获取错误信息
    // @ctx: 数据库上下文句柄
    const char* (*err_message)(db_context *ctx);
};

// SQL数据读取器
struct sql_data_reader {
    // 数据库上下文句柄
    db_context *db_ctx;
    // 语句句柄
    void *stmt;
    // 标志是否存在行
    uint32_t has_row;
    // 标志是结束标志
    uint32_t done;

    // 读取32位整形列值
    // @dr:读取器对象句柄
    // @column:列名称
    // @return:返回指定列整数值
    int32_t (*read_int32)(sql_data_reader *dr, const char *column);

    // 读取64位整形列值
    // @dr:读取器对象句柄
    // @column:列名称
    // @return:返回指定列整数值
    int64_t (*read_int64)(sql_data_reader *dr, const char *column);

    // 读取双精度浮点列值
    // @dr:读取器对象句柄
    // @column:列名称
    // @return:返回指定列双精度浮点数值
    double (*read_double)(sql_data_reader *dr, const char *column);

    // 读取文本列值
    // @dr:读取器对象句柄
    // @column:列名称
    // @return:返回指定列文本值
    const char * (*read_text)(sql_data_reader *dr, const char *column);

    // 读取32位整形列值
    // @dr:读取器对象句柄
    // @icol:列索引
    // @return:返回指定列整数值
    int32_t (*read_int32_by_index)(sql_data_reader *dr, uint32_t icol);

    // 读取64位整形列值
    // @dr:读取器对象句柄
    // @icol:列索引
    // @return:返回指定列整数值
    int64_t (*read_int64_by_index)(sql_data_reader *dr, uint32_t icol);

    // 读取双精度浮点列值
    // @dr:读取器对象句柄
    // @icol:列索引
    // @return:返回指定列双精度浮点数值
    double (*read_double_by_index)(sql_data_reader *dr, uint32_t icol);

    // 读取文本列值
    // @dr:读取器对象句柄
    // @icol:列索引
    // @return:返回指定列文本值
    const char * (*read_text_by_index)(sql_data_reader *dr, uint32_t icol);

    // 获取列类型
    // @dr:读取器对象句柄
    // @column:列名称
    // @return:返回指定列类型
    int (*get_type)(sql_data_reader *dr, const char *column);

    // 获取列类型
    // @dr:读取器对象句柄
    // @icol:列索引
    // @return:返回指定列类型
    int (*get_type_by_index)(sql_data_reader *dr, uint32_t icol);

    // 获取查询列大小
    // @dr:读取器对象句柄
    // @return:返回查询列总个数
    int (*get_columns_count)(sql_data_reader *dr);

    // 获取制定列的列名称
    // @dr:读取器对象句柄
    // @icol:列索引
    // @return:返回列名称
    const char* (*get_column_name)(sql_data_reader *dr, int32_t icol);

    // 检测是否存在下一行，如果存在，则获取下一行数据，并将游标指向该行
    // @dr:读取器对象句柄
    // @return:1 代表存在下一行 0 代表不存在下一行
    int (*next)(sql_data_reader *dr);

    // 关闭读取器
    // @dr:读取器对象句柄
    void (*close)(sql_data_reader *dr);
};

// SQL事务
struct sql_trans {
    // 提交事务
    // @ctx:数据库上下文句柄
    // @return:创建状态，０表示成功，非０表示失败
    int (*commit)(db_context *ctx);

    // 回滚事务
    // @ctx:数据库上下文句柄
    // @return:创建状态，０表示成功，非０表示失败
    int (*rollback)(db_context *ctx);
};

// 创建数据库上下文对象
// @conn_str:数据库链接字符串
// @ctx:待初始化的数据上下文对象句柄
// @return:创建状态，０表示成功，非０表示失败
int new_db_context(const char *conn_str, db_context *ctx);

// 数据库资源释放
void db_release();

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif/* SQLITEDB_DB_CONTEXT_H */
