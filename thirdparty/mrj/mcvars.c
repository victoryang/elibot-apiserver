#include "mcvars.h"

void get_cRobB_with_range(cJSON* root, int start, int end) {
	unsigned int *base;
	if (start > B_COUNT) {
		return NULL;
	}
	if (end > B_COUNT) {
		end = B_COUNT;
	}
	base = SHARE_RES(sysvar).cRobB + start;
	count = end - start + 1;
	cJSON_AddItemToObject(*root, "cRobB", cJSON_CreateIntArray((const int *)base, count));
}

void get_iRobI_with_range(cJSON* root, int start, int end) {
	int i=0;
	if (start > I_COUNT) {
		return NULL;
	}
	if (end > I_COUNT) {
		end = I_COUNT;
	}
	for (i=start; i<end; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateNumber(SHARE_RES(sysvar).iRobI[i]));
	}
	cJSON_AddItemToObject(*root, "iRobI", array);
}

void get_dRobD_with_range(cJSON* root, int start, int end) {
	double *base;
	if (start > D_COUNT) {
		return NULL;
	}
	if (end > D_COUNT) {
		end = D_COUNT;
	}
	base = SHARE_RES(sysvar).dRobD + start;
	count = end - start + 1;
	cJSON_AddItemToObject(*root, "dRobD", cJSON_CreateDoubleArray(base, count));
}

void get_dRobP_with_range(cJSON* root, int start, int end) {
	cJSON *array;
	int i;
	if (start > P_COUNT || end > P_COUNT) {
		return NULL;
	}
	
	array = cJSON_CreateArray();
	for (i=start; i<end, i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobP[i], AXIS_COUNT));
	}
	cJSON_AddItemToObject(*root, "dRobP", array);
}

void get_dRobV_with_range(cJSON* root, int start, int end) {
	cJSON *array;
	int i;
	if (start > V_COUNT || end > V_COUNT) {
		return NULL;
	}
	array = cJSON_CreateArray();
	for (i=start; i<end, i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobV[i], VSub_COUNT));
	}
	cJSON_AddItemToObject(*root, "dRobV", array);
}

getFunc getTable[] = {
	get_cRobB_with_range,
	get_iRobI_with_range,
	get_dRobD_with_range,
	get_dRobP_with_range,
	get_dRobV_with_range,
};

cJSON* get_data_with_range(int datatype, int start, int end) {
	getFunc gf;
	cJSON* root;
	switch datatype {
		case GetSysVarcRobB:
		case GetSysVariRobI:
		case GetSysVardRobD:
		case GetSysVardRobP:
		case GetSysVardRobV:
			root = cJSON_CreateObject();
			gf = getTable[datatype];
			gf(root, start, end);
			break;
		default:
			return NULL;
	}

	return root;
}
