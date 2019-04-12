// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef SQLITEDB_BACKUP_SQL_MAPPER_H
#define SQLITEDB_BACKUP_SQL_MAPPER_H

#include "db/sql_mapper.h"

#  ifdef __cplusplus
extern "C" {
#  endif /* __cplusplus */

// 备份全部参数
#define ELIBOT_BAK_BACKUP_PARAMS "elibot.bak.backupParams"

// 备份IO注释
#define ELIBOT_BAK_BACKUP_IO_COMMENTS "elibot.bak.backupIOComments"

/**
 * 获取备份映射器
 * @param id 映射器标识
 * @return 备份SQL映射器
 */
sql_mapper *get_bak_sql_mapper(const char *id);

#  ifdef __cplusplus
}
#  endif /* __cplusplus */
#endif //SQLITEDB_BACKUP_SQL_MAPPER_H
