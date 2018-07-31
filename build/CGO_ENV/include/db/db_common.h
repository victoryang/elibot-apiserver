// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef SQLITEDB_DB_COMMON_H
#define SQLITEDB_DB_COMMON_H

#include <stdint.h>

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

/**
 * 获取数组长度
 */
#define GET_ARRAY_LEN(array) (sizeof(array)/sizeof(array[0]))

/**
 * 时间记录器
 */
typedef struct time_sw time_sw;
struct time_sw {
    long start;
    long end;
    long ms;
};

// 时间记录开始
// @return:时间记录器对象
time_sw sw_start();

// 时间记录开始
// @sw:时间记录器句柄
// @return:返回记录时间 ms
long sw_stop(time_sw *sw);

// 时间记录开始
// @sw:时间记录器句柄
void sw_reset(time_sw *sw);

// 通过指定的分隔符讲字符串转化为浮点数数组
// @str:原始字符串
// @splitter:分隔符
// @out_array:输出，数组列表
void strtod_array(const char *str, const char *splitter, double *out_array);

// 通过指定的分隔符讲字符串转化为浮点数数组
// @str:原始字符串
// @splitter:分隔符
// @out_array:输出，数组列表
void strtoi_array(const char *str, const char *splitter, int *out_array);

// 通过指定的分隔符将字符串转化为字符串数组
// @str:原始字符串
// @splitter:分隔符
// @out_array:输出，数组列表
void strtostr_array(const char *str, const char *splitter, char *out_array[]);

// 将数组转化字符串
// @values:数组
// @len:数组长度
// @out_str:输出，字符串格式"{xxx,xxx}
void float_array_tostr(float values[], int32_t len, char *out_str);

// 将数组转化字符串
// @values:数组
// @len:数组长度
// @out_str:输出，字符串格式"{xxx,xxx}
void double_array_tostr(double values[], int32_t len, char *out_str);

// 将数组转化字符串
// @values:数组
// @len:数组长度
// @out_str:输出，字符串格式"{xxx,xxx}
void int_array_tostr(int values[], int32_t len, char *out_str);

// 将数组转化字符串
// @values:数组
// @len:数组长度
// @out_str:输出，字符串格式"{xxx,xxx}
void uint_array_tostr(uint32_t values[], int32_t len, char *out_str);

// 将数组转化字符串
// @values:数组
// @len:数组长度
// @out_str:输出，字符串格式"{xxx,xxx}
void short_array_tostr(short values[], int32_t len, char *out_str);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //SQLITEDB_DB_COMMON_H
