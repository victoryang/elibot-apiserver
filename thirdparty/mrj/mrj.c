#include "mrj.h"

cJSON* root;

static void add_items(cJSON* r, const char *name, RegisterFunc func) {
	cJSON* item = func();
	cJSON_bool ok = cJSON_IsInvalid(item);
	if(ok == cJSON_Invalid) {
		cJSON_Delete(item);
		return;
	}
	cJSON_AddItemToObject(r, name, item);
	return;
}

cJSON* get_resource_data() {
	int i=0;
	for(;Table[i].name!="ENDLINE"; i++)
	{
		add_items(root, Table[i].name, Table[i].func);
	}
	return root;
}

void init_resource() {
	root = cJSON_CreateObject();
	return 0;
}
