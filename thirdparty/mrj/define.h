#ifndef DEFINE_H
#define DEFINE_H

#include "stdlib.h"

#include "share/resource.h"
#include "cJSON.h"

typedef cJSON* (*RegisterFunc)();
typedef int (*RegisterFunc1)(cJSON** item);
typedef void (*getFunc)(cJSON*, int, int);

typedef struct {
	char 			*name;
	RegisterFunc	func;
}Item;

typedef struct {
	char 			*name;
	cJSON 			*root;
	RegisterFunc1	func;
}ResourceItem;

#endif //DEFINE_H