#include "mcvars.h"

void get_cRobB_with_range(cJSON* root, int start, int end) {
	unsigned int *base;
	int count;
	if (start > B_COUNT) {
		return;
	}
	if (end > B_COUNT) {
		end = B_COUNT;
	}
	base = SHARE_RES(sysvar).cRobB + start;
	count = end - start;
	cJSON_AddItemToObject(root, "cRobB", cJSON_CreateIntArray((const int *)base, count));
}

void get_iRobI_with_range(cJSON* root, int start, int end) {
	int i=0;
	cJSON *array;
	if (start > I_COUNT) {
		return;
	}
	if (end > I_COUNT) {
		end = I_COUNT;
	}

	array = cJSON_CreateArray();
	for (i=start; i<end; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateNumber(SHARE_RES(sysvar).iRobI[i]));
	}
	cJSON_AddItemToObject(root, "iRobI", array);
}

void get_dRobD_with_range(cJSON* root, int start, int end) {
	double *base;
	int count;
	if (start > D_COUNT) {
		return;
	}
	if (end > D_COUNT) {
		end = D_COUNT;
	}
	base = SHARE_RES(sysvar).dRobD + start;
	count = end - start;
	cJSON_AddItemToObject(root, "dRobD", cJSON_CreateDoubleArray(base, count));
}

void get_dRobP_with_range(cJSON* root, int start, int end) {
	cJSON *array;
	int i;
	if (start > P_COUNT || end > P_COUNT) {
		return;
	}
	
	array = cJSON_CreateArray();
	for (i=start; i<end; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobP[i], AXIS_COUNT));
	}
	cJSON_AddItemToObject(root, "dRobP", array);
}

void get_dRobV_with_range(cJSON* root, int start, int end) {
	cJSON *array;
	int i;
	if (start > V_COUNT || end > V_COUNT) {
		return;
	}

	array = cJSON_CreateArray();
	for (i=start; i<end; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobV[i], VSub_COUNT));
	}
	cJSON_AddItemToObject(root, "dRobV", array);
}

getSysVarFunc sysVarTable[] = {
	&get_cRobB_with_range,
	&get_iRobI_with_range,
	&get_dRobD_with_range,
	&get_dRobP_with_range,
	&get_dRobV_with_range,
};

cJSON* get_sysvar_with_range(int datatype, int start, int end) {
	getSysVarFunc gf;
	cJSON* root=NULL;
	switch (datatype) {
		case GetSysVarcRobB:
		case GetSysVariRobI:
		case GetSysVardRobD:
		case GetSysVardRobP:
		case GetSysVardRobV:
			root = cJSON_CreateObject();
			gf = sysVarTable[datatype];
			gf(root, start, end);
			break;
		default:
			return NULL;
	}

	return root;
}

void get_cRobLB_with_range(cJSON* root, int num, int start, int end) {
	unsigned int *base;
	int count;
	if (start >= LB_COUNT || num >= CALL_NEST_NUM) {
		return;
	}
	if (end > LB_COUNT) {
		end = LB_COUNT;
	}
	base = SHARE_RES(locvar)[num].cRobLB + start;
	count = end - start;
	cJSON_AddItemToObject(root, "cRobLB", cJSON_CreateIntArray((const int *)base, count));
}

void get_iRobLI_with_range(cJSON* root, int num, int start, int end) {
	int i=0;
	cJSON *array;
	if (start >= LI_COUNT || num >= CALL_NEST_NUM) {
		return;
	}
	if (end > LI_COUNT) {
		end = LI_COUNT;
	}

	array = cJSON_CreateArray();
	for (i=start; i<end; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateNumber(SHARE_RES(locvar)[num].iRobLI[i]));
	}
	cJSON_AddItemToObject(root, "iRobLI", array);
}

void get_dRobLD_with_range(cJSON* root, int num, int start, int end) {
	double *base;
	int count;
	if (start >= LD_COUNT || num >= CALL_NEST_NUM) {
		return;
	}
	if (end > LD_COUNT) {
		end = LD_COUNT;
	}
	base = SHARE_RES(locvar)[num].dRobLD + start;
	count = end - start;
	cJSON_AddItemToObject(root, "dRobLD", cJSON_CreateDoubleArray(base, count));
}

void get_dRobLP_with_range(cJSON* root, int num, int start, int end) {
	cJSON *array;
	int i;
	if (start > LP_COUNT || end > LP_COUNT || num >= CALL_NEST_NUM) {
		return;
	}
	
	array = cJSON_CreateArray();
	for (i=start; i<end; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(locvar)[num].dRobLP[i], AXIS_COUNT));
	}
	cJSON_AddItemToObject(root, "dRobLP", array);
}

void get_dRobLV_with_range(cJSON* root, int num, int start, int end) {
	cJSON *array;
	int i;
	if (start > LV_COUNT || end > LV_COUNT || num >= CALL_NEST_NUM) {
		return;
	}

	array = cJSON_CreateArray();
	for (i=start; i<end; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(locvar)[num].dRobLV[i], VSub_COUNT));
	}
	cJSON_AddItemToObject(root, "dRobLV", array);
}

getLocVarFunc locVarTable[] = {
	&get_cRobLB_with_range,
	&get_iRobLI_with_range,
	&get_dRobLD_with_range,
	&get_dRobLP_with_range,
	&get_dRobLV_with_range,
}

cJSON* get_locvar_with_range(int datatype, int number, int start, int end) {
	getLocVarFunc gf;
	cJSON* root=NULL;
	switch (datatype) {
		case GetSysVarcRobLB:
		case GetSysVariRobLI:
		case GetSysVardRobLD:
		case GetSysVardRobLP:
		case GetSysVardRobLV:
			root = cJSON_CreateObject();
			gf = locVarTable[datatype];
			gf(root, number, start, end);
			break;
		default:
			return NULL;
	}

	return root;
}