// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_DB_METADATA_H
#define MCSERVER_DB_METADATA_H

#include "db_hash.h"

#include <stdint.h>

/**
 * 元数据内存寻址映射器
 */
#define DEFINE_DB_MD_MEM_MAPPER_PHYSICAL_PARENT_DIM(_name) static uint32_t _name##_physical_parent_dim[]
#define DEFINE_DB_MD_MEM_MAPPER_PHYSICAL_DIM(_name) static uint32_t _name##_physical_dim[]
#define DEFINE_DB_MD_MEM_MAPPER_ELEM_TYPE_SIZE_LIST(_name) static uint32_t _name##_elem_type_size_list[]
#define DEFINE_DB_MD_MEM_MAPPER_ADDR_MODE_LIST(_name) static uint32_t _name##_addr_mode_list[]
#define DEFINE_DB_MD_MEM_MAPPER_OFFSET_LIST(_name) static uint32_t _name##_offset_list[]
#define DEFINE_DB_METADATA_MEM_MAPPER(_name, _id, _addr_times, _physical_dim_num, _physical_parent_dim_num) \
static db_metadata_mem_mapper _name##_md_mem_mapper = {\
    .md_id=_id,\
    .addr_times=_addr_times,\
    .elem_type_size_list = &_name##_elem_type_size_list[0],\
    .addr_mode_list = &_name##_addr_mode_list[0],\
    .offset_list = &_name##_offset_list[0],\
    .physical_dim_num=_physical_dim_num,\
    .physical_dim = &_name##_physical_dim[0],\
    .physical_parent_dim_num=_physical_parent_dim_num,\
    .physical_parent_dim = &_name##_physical_parent_dim[0],\
}

/**
 * 偏移量计算
 */
#define MD_OFFSET(type, elem) ((uint32_t)(&((type*)0)->elem))

// 变量字符串最大值
#define VAR_STR_MAX_LEN 128

// 数组维度最大值
#define MAX_ARR_DIM 5

/**
 * 元数据内存结构映射器
 */
typedef struct db_metadata_mem_mapper db_metadata_mem_mapper;
struct db_metadata_mem_mapper {
    // 参数唯一标识
    char *md_id;

    // 寻址次数
    int8_t addr_times;

    // 元素类型大小列表
    int32_t *elem_type_size_list;
    // 寻址模式列表
    int32_t *addr_mode_list;
    // 偏移列表
    int32_t *offset_list;

    // 物理内存数组维度信息
    int32_t physical_dim_num;
    int32_t *physical_dim;

    // 父级物理内存数组维度信息
    int32_t physical_parent_dim_num;
    int32_t *physical_parent_dim;
};

/**
 * 元数据类型生命
 */
typedef struct db_metadata db_metadata;
struct db_metadata {
    // 参数唯一标识
    char md_id[VAR_STR_MAX_LEN];
    // 分类信息
    char clazz[VAR_STR_MAX_LEN];
    // 参数分组
    char group[VAR_STR_MAX_LEN];
    // 参数名称
    char name[VAR_STR_MAX_LEN];
    // 工艺类型
    int tech;
    // 元素类型
    uint32_t elem_type;
    // 元素类型字符串表示
    char elem_type_str[VAR_STR_MAX_LEN];
    // 是否数组类型
    int8_t is_array;
    // 是否是数值类型
    int8_t is_number;
    // 数组维度个数
    int8_t dim_num;
    // 数组各维度值
    int8_t dim[MAX_ARR_DIM];
    // 元素大小
    int size;
    // 是否枚举
    int8_t is_enum;
    // 枚举类型
    char *enum_type[VAR_STR_MAX_LEN];
    // 是否引用
    int8_t is_ref;
    // 引用类型
    char ref_type[VAR_STR_MAX_LEN];
    // 权限等级
    uint32_t auth_level;
    // 读等级
    uint32_t read_level;
    // 写等级
    uint32_t write_level;
    // 字符串最大长度
    uint32_t str_len;

    // 最大值
    void *max_value;
    // 最小值
    void *min_value;
    // 默认值
    void *default_value;

    // 映射器
    db_metadata_mem_mapper *mapper;
};

/**
 * 元数据仓储
 */
typedef struct db_metadata_rep db_metadata_rep;
struct db_metadata_rep {
    db_hash *reps;
    int inited;
};

/**
 * 初始化仓储
 */
int init_metadatas_repository();

/**
 * 获取元数据句柄
 * @param id 元数据标识
 * @return 返回元数据句柄
 */
db_metadata *get_db_metadata_by_id(const char *id);

/**
 * 获取元数据句柄
 * @param group 分组名称
 * @param name 元数据名称
 * @return 返回元数据句柄
 */
db_metadata *get_db_metadata(const char *group, const char *name);

/**
 * 获取元数据内存映射器句柄
 * @param md_id 元数据标识
 * @return 元数据内存映射器句柄
 */
db_metadata_mem_mapper* find_metadata_mapper(char *md_id);

/**
 * 初始化元数据映射器缓存
 */
int init_metadata_mapper_cache();

/**
 * 查找内存地址
 * @param md 元数据句柄
 * @param base_addr 基地址
 * @param indexs 间接寻址索引列表
 * @param indexs_len 索引列表长度
 * @return 变量真实地址
 */
void *lookup_mem_addr(db_metadata *md, void *base_addr, unsigned int indexs[], int indexs_len);

/**
 * 查找内存地址
 * @param md 元数据句柄
 * @param base_addr 基地址
 * @param indexs 间接寻址索引列表
 * @param indexs_len 索引列表长度
 * @return 变量真实地址
 */
void *lookup_mem_addr_by_direct(db_metadata *md, void *base_addr);

#endif //MCSERVER_DB_METADATA_H
