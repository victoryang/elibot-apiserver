#include "items.h"

cJSON* get_dspInfor() {
	cJSON* item;
	item = cJSON_CreateObject();
	cJSON_AddNumberToObject(item, "display_state", GetDisplayMsg_State());
	cJSON_AddStringToObject(item, "display_msg", GetDisplayMsg());
	return item;
}

Item Table[] = {
	{"dspInfor", &get_dspInfor},
	{"ENDLINE", NULL},
};