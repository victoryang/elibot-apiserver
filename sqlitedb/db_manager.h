// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_DB_MANAGER_H
#define MCSERVER_DB_MANAGER_H

#include <stdint.h>

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 数据库操作模式 1 备份 2 恢复　3 升级
enum {DB_BACKUP = 1, DB_RESTORE = 2, DB_UPGRADE = 3};

// 数据库管理器
typedef struct db_manager db_manager;
struct db_manager {
    /* 数据库联接串 */
    char *conn_string;
    /* 数据库目录 */
    char *db_dir;
    /* 数据库备份文件名 */
    char *db_bak_name;
    /* 数据库操作模式 */
    int op_mode;
    /* 升级文件路径,该参数仅升级模式下有效 */
    char *upgrade_pkg;
    /* 包解压密码 */
    char *passwd;

    /* 指令包 */
    char *commands;

    /* 强力恢复选项 */
    char force_store;

    /**
     * 启动数据库管理
     * @param mgr 管理器实例句柄
     * @param output 输出,根据实际情况而定，目前仅备份时输出备份文件
     * @return 操作码
     */
    int (*execute)(db_manager *mgr, void *output);

    /**
     * 启动数据库管理
     * @param mgr 管理器实例句柄
     * @param output 输出,根据实际情况而定，目前仅备份时输出备份文件
     * @return 操作码
     */
    int (*_execute)(db_manager *mgr, void *output);
};


/**
 * 创建数据库管理器
 * @param conn_str 联接字符串
 * @param db_dir 数据库目录
 * @param db_bak_name 数据库备份文件名
 * @param mode 数据操作模式
 * @param upgrade_pkg 升级文件路径
 * @param mgr 数据库管理器句柄
 * @param force_restore 强制恢复
 * @return 数据库管理器对象
 */
int new_db_manager(const char *conn_str
                            , const char *db_dir
                            , const char *db_bak_name
                            , int mode
                            , const char *upgrade_pkg
                            , db_manager *mgr
                            , char force_restore);

/**
 * 创建备份数据库管理器
 * @param conn_str 联接字符串
 * @param db_dir 数据库目录
 * @param mgr 数据库管理器句柄
 * @return 数据库管理器对象
 */
#define new_backup_db_manager(conn_str, db_dir, mgr) \
    new_db_manager(conn_str, db_dir, NULL, DB_BACKUP, NULL, mgr, 0)


/**
 * 创建恢复数据库管理器
 * @param conn_str 联接字符串
 * @param db_dir 数据库目录
 * @param db_bak_name 数据库备份文件名
 * @param mgr 数据库管理器句柄
 * @param force 强制恢复
 * @return 数据库管理器对象
 */
#define new_restore_db_manager(conn_str, db_dir, db_bak_name, mgr, force) \
    new_db_manager(conn_str, db_dir, db_bak_name, DB_RESTORE, NULL, mgr, force)

/**
 * 创建升级数据库管理器
 * @param conn_str 联接字符串
 * @param db_dir 数据库目录
 * @param upgrade_pkg 升级文件路径
 * @param mgr 数据库管理器句柄
 * @return 数据库管理器对象
 */
#define new_upgrade_db_manager(conn_str, db_dir, upgrade_pkg, mgr) \
    new_db_manager(conn_str, db_dir, NULL, DB_UPGRADE, upgrade_pkg, mgr, 0)

/**
 * 设置命令包
 * @param mgr 数据库管理器句柄
 * @param cmds 命令包句柄
 */
#define set_cmds(mgr, cmds) do{(mgr)->commands = cmds;}while(0)

/**
 * 获取数据库版本
 * @param conn_str 数据库联接字符串
 * @param version 版本号(输出）
 * @return 执行状态
 */
int get_db_version(const char *conn_str, char *version);

/**
 * 错误编码
 */
#define DB_MANAGER_OK 0                             // 操作成功
#define DB_MANAGER_ERR 1                            // 操作失败
#define DB_MANAGER_ERR_VERSION_LT -100              // 版本号低于当前版本
#define DB_MANAGER_ERR_VERSION_GT -110              // 版本号高于当前版本
#define DB_MANAGER_ERR_UPGRADE_FAILURE -120         // 升级失败
#define DB_MANAGER_ERR_RESTORE_FAILURE -130         // 恢复失败
#define DB_MANAGER_ERR_BACKUP_FAILURE -140          // 备份失败


#  ifdef __cplusplus
}
#  endif /* __cplusplus */

#endif //MCSERVER_DB_MANAGER_H
