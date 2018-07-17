#ifndef ITEMS_H
#define ITEMS_H

#include "share/resource.h"
#include "cJSON.h"

typedef cJSON* (*func_ptr)() RegisterFunc

typedef struct {
	char 			*name;
	RegisterFunc	func;
}Item;

static cJSON* get_dspInfor();

Item Table[] = {
	{"dspInfor", &get_dspInfor},
	{"ENDLINE", NULL},
};

#endif //ITEMS_H