#ifndef DEFINE_H
#define DEFINE_H

#include "stdlib.h"

#include "share/resource.h"
#include "cJSON.h"

typedef void (*getSysVarFunc)(cJSON*, int, int);
typedef void (*getLocVarFunc)(cJSON*, int, int, int);

#endif //DEFINE_H