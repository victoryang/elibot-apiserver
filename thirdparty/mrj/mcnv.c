#include "mcnv.h"

cJSON* get_zero_encode() {
	int i = 0;
	cJSON* item;
	item = cJSON_CreateArray();
	for(i=0; i<AXIS_COUNT; i++) {
		cJSON_AddItemToArray(item, cJSON_CreateNumber(Get_ZeroEncode(i)));
	}
	return item;
}

cJSON* get_origin() {
	int i = 0;
	cJSON* item;
	item = cJSON_CreateArray();
	for(i=0; i<AXIS_COUNT; i++) {
		cJSON_AddItemToArray(item, cJSON_CreateNumber(Get_RobotRes_Origin(i)));
	}
	return item;
}

cJSON* get_nv() {
	cJSON* item;
	item = cJSON_CreateObject();
	cJSON_AddItemToObject(item, "projectname", cJSON_CreateString(GetMainfile()));
	cJSON_AddItemToObject(item, "sys_technics", cJSON_CreateNumber(getTechnics()));
	cJSON_AddItemToObject(item, "curcoordinate", cJSON_CreateNumber(GetCurCoordinate()));
	cJSON_AddItemToObject(item, "manualspeedRate", cJSON_CreateNumber(GetManualSpeedRate()));
	cJSON_AddItemToObject(item, "cyclemode", cJSON_CreateNumber(GetCycleMode()));
	cJSON_AddItemToObject(item, "toolnum", cJSON_CreateNumber(GetToolNumber()));
	cJSON_AddItemToObject(item, "usernum", cJSON_CreateNumber(GetUserNumber()));
	cJSON_AddItemToObject(item, "zero_encode", get_zero_encode());
	cJSON_AddItemToObject(item, "system_period", cJSON_CreateNumber(GetSysPeriod()));
	cJSON_AddItemToObject(item, "system_ctrl_mode", cJSON_CreateNumber(GetSysCtrlMode()));
	cJSON_AddItemToObject(item, "origin", get_origin());
	return item;
}