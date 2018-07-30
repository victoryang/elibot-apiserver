#ifndef MCVARS_H
#define MCVARS_H

#include "define.h"
#define GetSysVarcRobB 0
#define GetSysVariRobI 1
#define GetSysVardRobD 2
#define GetSysVardRobP 3
#define GetSysVardRobV 4

cJSON* get_data_with_range(int datatype, int start, int end);

#endif //MCVARS_H