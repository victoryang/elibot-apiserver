// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_DB_HASH_H
#define MCSERVER_DB_HASH_H

#define DB_HASH_OK 0
#define DB_HASH_ERR -1

/**
 * 哈希数据项
 */
typedef struct db_hash_entry db_hash_entry;
struct db_hash_entry {
    void *value;
    char *key;
    unsigned int hash_code;

    db_hash_entry *pre;
    db_hash_entry *next;
    unsigned int len;
};

/**
 * 哈希表
 */
typedef struct db_hash db_hash;
struct db_hash {
    db_hash_entry **entries;
    unsigned int size;
};


/**
 * MurmurHash2 哈希函数
 * @param key 哈希键
 * @param len 哈希键长
 * @return 哈希值
 */
unsigned int mmhash_32(const char *key, int len);

/**
 * 创建哈希表对象
 * @param size 哈希大小
 * @return 哈希对象句柄
 */
db_hash *db_hash_create(unsigned int size);

/**
 * 添加键值对
 * @param h 哈希表对象句柄
 * @param key 哈希键
 * @param value 哈希值
 * @return 0 标识成功 -1 标识失败
 */
int db_hash_add(db_hash *h, const char *key, void *value);

/**
 * 获取键值
 * @param h 哈希标对象句柄
 * @param key 哈希键
 * @return 哈希值句柄
 */
void *db_hash_get(db_hash *h, const char *key);

/**
 * 清除哈希表
 * @param h 哈希标对象句柄
 */
void db_hash_clear(db_hash *h);

/**
 * 释放哈希表
 * @param h 哈希标对象句柄
 */
void db_hash_release(db_hash *h);

#endif //MCSERVER_DB_HASH_H
