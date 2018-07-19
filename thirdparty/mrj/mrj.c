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

static char* add_items_to_body(Item items[], char *name) {
	cJSON* root = cJSON_CreateObject();
	cJSON* body = cJSON_CreateObject();
	cJSON_AddItemToObject(root, name, body);
	int i;
	for(i=0; items[i].func!=NULL; i++)
	{
		add_items(body, items[i].name, items[i].func);
	}

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

char* get_resource_data() {
	return add_items_to_body(ResourceTable, "resource");
}

char* get_nv_data() {
	cJSON* root = cJSON_CreateObject();
	cJSON* body = cJSON_CreateObject();
	cJSON_AddItemToObject(root, "nv", get_nv());

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

int mrj_init() {
	return resource_init("/rbctrl/mcserver");
}
