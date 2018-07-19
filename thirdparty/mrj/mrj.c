#include "mrj.h"

static void add_items(cJSON* r, const char *name, RegisterFunc func) {
	cJSON* item = (*func)();
	/*cJSON_bool isInvalid = cJSON_IsInvalid(item);
	if(isInvalid == cJSON_Invalid) {
		return;
	}*/
	cJSON_AddItemToObject(r, name, item);
	return;
}

char* get_resource_data() {
	cJSON* root = cJSON_CreateObject();
	int i=0;
	for(;Table[i].func!=NULL; i++)
	{
		add_items(root, Table[i].name, Table[i].func);
	}

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

int mrj_init() {
	return resource_init("/rbctrl/mcserver");
}
