#ifndef MRJ_H
#define MRJ_H

#include "mcresource.h"
#include "mcnv.h"

char* get_resource_data();
char* get_nv_data();

int init_nv_ram();
int mrj_init();

#endif //MRJ_H