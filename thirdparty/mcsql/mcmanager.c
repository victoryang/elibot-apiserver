#include "mcmanager.h"

int execute(struct db_manager* mgr) {
	set_cmds(mgr, ELIBOT_BAK_BACKUP_PARAMS);
	void* output = malloc(sizeof(char) * BUF_SIZE);
	return mgr->execute(mgr, output);
}

int mcmanager_backup_db(char* conn, char* db_dir) {
	db_manager* mgr = (db_manager*)malloc(sizeof(db_manager));
    if(mgr == NULL) {
        return ERREMPTYMANAGER;
    }
    
    new_backup_db_manager(conn, db_dir, mgr);
    return execute(mgr);
}

int mcmanager_restore_db(char* conn, char* db_dir, char* db_bak_name, char force) {
	db_manager* mgr = (db_manager*)malloc(sizeof(db_manager));
    if(mgr == NULL) {
        return ERREMPTYMANAGER;
    }

    new_restore_db_manager(conn, db_dir, db_bak_name, mgr, force);
    return execute(mgr);
}

int mcmanager_upgrade_db(char* conn, char* db_dir, char* upgrade_pkg) {
	db_manager* mgr = (db_manager*)malloc(sizeof(db_manager));
    if(mgr == NULL) {
        return ERREMPTYMANAGER;
    }

    new_upgrade_db_manager(conn, db_dir, upgrade_pkg, mgr);
    return execute(mgr);
}