#ifndef MCVARS_H
#define MCVARS_H

#include "define.h"
#define GetSysVarcRobB 0
#define GetSysVariRobI 1
#define GetSysVardRobD 2
#define GetSysVardRobP 3
#define GetSysVardRobV 4
#define GetSysVarcRobLB GetSysVarcRobB
#define GetSysVariRobLI GetSysVariRobI
#define GetSysVardRobLD GetSysVardRobD
#define GetSysVardRobLP GetSysVardRobP
#define GetSysVardRobLV GetSysVardRobV

cJSON* get_sysvar_with_range(int datatype, int start, int end);
cJSON* get_locvar_with_range(int datatype, int number, int start, int end);

#endif //MCVARS_H