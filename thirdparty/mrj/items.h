#ifndef ITEMS_H
#define ITEMS_H

#include "share/resource.h"
#include "cJSON.h"

typedef cJSON* (*RegisterFunc)();

typedef struct {
	char 			*name;
	RegisterFunc	func;
}Item;

extern Item Table[];

#endif //ITEMS_H