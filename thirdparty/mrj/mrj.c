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

cJSON* get_resource_data() {
	cJSON* root = cJSON_CreateObject();
	int i=0;
	for(;Table[i].func!=NULL; i++)
	{
		add_items(root, Table[i].name, Table[i].func);
	}
	return root;
}

int mrj_init() {
	return resource_init();
}
