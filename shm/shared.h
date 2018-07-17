#ifndef SHARED_H
#define SHARED_H
#include "db/cJSON.h"

unsigned char get_autorun_cyclemode();

cJSON* get_resource_data();

int init_shared_resource();

#endif //SHARED_H