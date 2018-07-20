#ifndef MCMANAGER_H
#define MCMANAGER_H

#include "stdlib.h"
#include "db/db_manager.h"

int DBMgrExecute(struct db_manager* mgr, void *output);
db_manager* NewDBManager();

#endif //MCMANAGER_H