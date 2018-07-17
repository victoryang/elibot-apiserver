#include "items.h"

static cJSON* get_dspInfor() {
	cJSON* item;
	root = cJSON_CreateObject();
	cJSON_AddNumberToObject(item, "display_state", GetDisplayMsg_State());
	cJSON_AddStringToObject(item, "display_msg", GetDisplayMsg());
	return item;
}