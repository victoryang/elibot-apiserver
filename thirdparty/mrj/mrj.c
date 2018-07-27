#include "mrj.h"

cJSON* resource_root;

static cJSON* add_items_to_body(ResourceItem items[]) {
	cJSON* body = cJSON_CreateObject();
	int i;

	for(i=0; items[i].func!=NULL; i++)
	{
		(*(items[i].func))(&(items[i].root));
		/*cJSON_bool isInvalid = cJSON_IsInvalid(item);
		if(isInvalid == cJSON_Invalid) {
			return;
		}*/
		cJSON_AddItemToObject(body, items[i].name, items[i].root);
	}
	return body;
}

void init_resource_memory() {
	cJSON* body;
	resource_root = cJSON_CreateObject();
	body = add_items_to_body(ResourceTable);
	cJSON_AddItemToObject(resource_root, "resource", body);
	return;
}

void free_resource_memory() {
	cJSON_Delete(resource_root);
}

static int get_items(ResourceItem items[]) {
	int i;
	int changed = 0;
	for(i=0; items[i].func!=NULL; i++)
	{
		changed = changed | (*(items[i].func))(&(items[i].root));
	}
	return changed;
}

char* get_resource_data() {
	int changed = get_items(ResourceTable);
	if (changed != 0) {
		return cJSON_PrintUnformatted(resource_root);
	}
	return NULL;
}

char* get_nv_data() {
	cJSON* root = cJSON_CreateObject();
	cJSON_AddItemToObject(root, "nv", get_nv());

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

int init_nv_ram() {
	return init_nvram(NULL);
}

int mrj_init() {
	int ret = resource_init("/rbctrl/mcserver");
	if (ret != 0) {
		return ret;
	}

	init_resource_memory();
	return 0;
}
