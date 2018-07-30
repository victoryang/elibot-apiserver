#ifndef DEFINE_H
#define DEFINE_H

#include "stdlib.h"

#include "share/resource.h"
#include "cJSON.h"

typedef int (*RegisterFunc)(cJSON** item);
typedef void (*getSysVarFunc)(cJSON*, int, int);
typedef void (*getLocVarFunc)(cJSON*, int, int, int);

typedef struct {
	char 			*name;
	cJSON 			*root;
	RegisterFunc	func;
}ResourceItem;

#endif //DEFINE_H