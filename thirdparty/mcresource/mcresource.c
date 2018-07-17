#include "mcresource.h"

cJSON* root;

static int register_struct(cJSON* r, const char *name, cJSON* item) {
	cJSON_bool ok = cJSON_IsInvalid(item);
	if(ok == cJSON_Invalid) {
		return -1;
	}
	cJSON_AddItemToObject(r, name, item);
	return 0;
}

cJSON* get_resource_data()
{
	int i=0;
	while(Table[i].name!=LASTITEM.name)
	{
		cJSON* item = (Table[i].func)();
		int ret = cJSON_AddItemToObject(root, Table[i].name, item);
		if(ret!=0)
			break;
	}
	return root;
}

void init() {
	root = cJSON_CreateObject();
	return 0;
}