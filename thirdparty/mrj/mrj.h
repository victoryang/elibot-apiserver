#ifndef MRJ_H
#define MRJ_H

#include "mcplc.h"
#include "mcresource.h"
#include "mcnv.h"
#include "mcvars.h"

char* get_sysvar_data(int datatype, int start, int end);
char* get_locvar_data(int datatype, int number, int start, int end);
char* get_plc_data();
char* get_resource_data();
char* get_nv_data();

int init_nv_ram();
int mrj_init();

#endif //MRJ_H