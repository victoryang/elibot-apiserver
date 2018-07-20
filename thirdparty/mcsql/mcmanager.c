#include "mcmanager.h"

int DBMgrExecute(struct db_manager* mgr, void *output) {
    return mgr->execute(mgr, output);
};

db_manager* NewDBManager() {
    db_manager* mgr = (db_manager*)malloc(sizeof(db_manager));
    if(mgr == NULL) {
        return NULL;
    }
    return mgr;
}