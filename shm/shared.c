#include "share/resource.h"
#include "cJSON.h"

unsigned char get_autorun_cyclemode() {
	return GET_AUTORUN_CYCLEMODE();
}

static cJSON* get_struct_dspInfor() {
	cJSON* item;
	item = cJSON_CreateObject();
	cJSON_AddNumberToObject(item, "display_state", GetDisplayMsg_State());
	cJSON_AddStringToObject(item, "display_msg", GetDisplayMsg());
	return item;
}

cJSON* get_resource_data() {
	cJSON* root;
	root = cJSON_CreateObject();
	cJSON_AddItemToObject(root, "dspInfor", get_dspInfor());
	return root;
}

int init_shared_resource() {
	return resource_init("/rbctrl/mcserver");
}

