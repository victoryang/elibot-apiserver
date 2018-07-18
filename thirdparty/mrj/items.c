#include "items.h"

cJSON* get_dspInfor() {
	cJSON* item;
	item = cJSON_CreateObject();
	cJSON_AddItemToObject(item, "display_state", cJSON_CreateNumber(GetDisplayMsg_State()));
	cJSON_AddItemToObject(item, "display_msg", cJSON_CreateString(GetDisplayMsg()));
	return item;
}

Item Table[] = {
	{"dspInfor", &get_dspInfor},
	{"ENDLINE", NULL},
};