#include "mrj.h"

char* get_sysvar_data(int datatype, int start, int end) {
	cJSON* root;
	char* ret;
	root = get_sysvar_with_range(datatype, start, end);
	ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

char* get_locvar_data(int datatype, int number, int start, int end) {
	cJSON* root;
	char* ret;
	root = get_locvar_with_range(datatype, number, start, end);
	ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

char* get_plc_data() {
	cJSON* root = cJSON_CreateObject();
	cJSON_AddItemToObject(root, "plc", get_plc());

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

char* get_resource_data() {
	cJSON* root = cJSON_CreateObject();
	cJSON_AddItemToObject(root, "resource", get_resource());

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
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
	return resource_init("/rbctrl/mcserver");
}
