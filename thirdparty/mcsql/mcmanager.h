#ifndef MCMANAGER_H
#define MCMANAGER_H

#include "stdlib.h"
#include "define.h"
#include "db/db_manager.h"

#define BUF_SIZE 255
#define ERRNONE DB_OK
#define ERREMPTYMANAGER -1

int mcmanager_backup_db(char* conn, char* db_dir);

int mcmanager_restore_db(char* conn, char* db_dir, char* db_bak_name, char force);

int mcmanager_upgrade_db(char* conn, char* db_dir, char* upgrade_pkg);

#endif //MCMANAGER_H