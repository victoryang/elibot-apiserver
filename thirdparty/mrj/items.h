#ifndef ITEMS_H
#define ITEMS_H

#include "share/resource.h"
#include "cJSON.h"

typedef RegisterFunc cJSON* (*func)();

typedef struct {
	char 			*name;
	RegisterFunc	func;
}Item;

const Item LASTITEM = {"ENDLINE", NULL}

Item Table[] = {
	{"dspInfor", get_dspInfor},
	LASTITEM
};

#endif //ITEMS_H