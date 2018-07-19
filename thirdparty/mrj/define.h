#ifndef DEFINE_H
#define DEFINE_H

#include "stdlib.h"

#include "share/resource.h"
#include "cJSON.h"

typedef cJSON* (*RegisterFunc)();

typedef struct {
	char 			*name;
	RegisterFunc	func;
}Item;

#endif //DEFINE_H