#include "mcresource.h"

bool getAutorunCycleMode(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(GET_AUTORUN_CYCLEMODE());
		return true;
	}
	if (item->valuedouble == GET_AUTORUN_CYCLEMODE()){
		return false;
	}
	cJSON_SetNumberValue(item, GET_AUTORUN_CYCLEMODE());
	return true;
}

bool getTeachMode(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(GET_TEACH_MODE());
		return true;
	}
	if (item->valuedouble == GET_TEACH_MODE()){
		return false;
	}
	cJSON_SetNumberValue(item, GET_TEACH_MODE());
	return true;
}

bool getMachinePos(cJSON* item) {
	size_t i = 0;
	bool changed = false;
	cJSON* temp;
	double now;
	if (NULL == item) {
		item = cJSON_CreateDoubleArray(RobotRes_MachPos_base, AXIS_COUNT);
		return true;
	}
	for(i = 0;i < (size_t)AXIS_COUNT; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = *(RobotRes_MachPos_base + i)
		if (temp->valuedouble != now) {
			cJSON_SetNumberValue(temp, now);
			changed = true;
		}
	}
	
	return changed;
}

bool getMachinePose(cJSON* item) {
	size_t i = 0;
	bool changed = false;
	cJSON* temp;
	double now;
	if (NULL == item) {
		item = cJSON_CreateDoubleArray(RobotRes_MachPose_base, 6);
		return true;
	}
	for(i = 0;i < 6; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = *(RobotRes_MachPose_base + i)
		if (temp->valuedouble != now) {
			cJSON_SetNumberValue(temp, now);
			changed = true;
		}
	}
	
	return changed;
}

bool getAbsPulse(cJSON* item) {
	size_t i = 0;
	bool changed = false;
	cJSON* temp;
	int32_t now;
	if (NULL == item) {
		item = cJSON_CreateIntArray((const int *)SHARE_RES(abs_pulse), AXIS_COUNT);
		return true;
	}
	for(i = 0;i < AXIS_COUNT; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = *((const int *)SHARE_RES(abs_pulse) + i)
		if (temp->valuedouble != (double)now) {
			cJSON_SetNumberValue(temp, (double)now);
			changed = true;
		}
	}
	
	return changed;
}

bool getAbzPulse(cJSON* item) {
	size_t i = 0;
	bool changed = false;
	cJSON* temp;
	int32_t now;
	if (NULL == item) {
		item = cJSON_CreateIntArray((const int *)SHARE_RES(abz_pulse), AXIS_COUNT);
		return true;
	}
	for(i = 0;i < AXIS_COUNT; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = *((const int *)SHARE_RES(abz_pulse) + i)
		if (temp->valuedouble != (double)now) {
			cJSON_SetNumberValue(temp, (double)now);
			changed = true;
		}
	}
	
	return changed;
}

bool getCurEncode(cJSON* item) {
	size_t i = 0;
	bool changed = false;
	cJSON* temp;
	int32_t now;
	if (NULL == item) {
		item = cJSON_CreateIntArray((const int *)SHARE_RES(cur_encode), AXIS_COUNT);
		return true;
	}
	for(i = 0;i < AXIS_COUNT; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = *((const int *)SHARE_RES(cur_encode) + i)
		if (temp->valuedouble != (double)now) {
			cJSON_SetNumberValue(temp, (double)now);
			changed = true;
		}
	}
	
	return changed;
}

bool getRobotState(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(GetRobotState());
		return true;
	}
	if (item->valuedouble == GetRobotState()){
		return false;
	}
	cJSON_SetNumberValue(item, GetRobotState());
	return true;
}

bool getServoReady(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(GetServoReady());
		return true;
	}
	if (item->valuedouble == GetServoReady()){
		return false;
	}
	cJSON_SetNumberValue(item, GetServoReady());
	return true;
}

bool getCurLine(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(GetCurLine());
		return true;
	}
	if (item->valuedouble == GetCurLine()){
		return false;
	}
	cJSON_SetNumberValue(item, GetCurLine());
	return true;
}

bool updateUint8ArrayElementIf(cJSON* array, volatile uint8_t* base) {
	size_t i = 0;
	bool changed = true;
	size_t num = cJSON_GetArraySize(array);
	cJSON* temp;
	volatile uint8_t now;
	for(i = 0;i < num; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = base[i];
		if ( temp->valuedouble!= (double)now) {
			cJSON_SetNumberValue(temp, (double)now);
			changed = true;
		}
	}
	return changed;
}

cJSON* setElementToArray(volatile uint8_t* base, int array_size) {
	int i=0;
	cJSON* array = cJSON_CreateArray();
	for (;i<array_size;i++) {
		cJSON_AddItemToArray(array, cJSON_CreateNumber((double)base[i]));
	}
	return array;
}
bool getRobPLC(cJSON* item) {
	bool changed = false;
	if (NULL == item) {
		item = cJSON_CreateObject();
		cJSON_AddItemToObject(item, "PLC_IN", setElementToArray(SHARE_RES(plc).PLC_IN, IO_IN_NUM));
		cJSON_AddItemToObject(item, "PLC_OUT", setElementToArray(SHARE_RES(plc).PLC_OUT, IO_OUT_NUM));
		cJSON_AddItemToObject(item, "PLC_VIN", setElementToArray(SHARE_RES(plc).PLC_VIN, IO_VIN_NUM));
		cJSON_AddItemToObject(item, "PLC_VOUT", setElementToArray(SHARE_RES(plc).PLC_VOUT, IO_VOUT_NUM));
		cJSON_AddItemToObject(item, "PLC_M", setElementToArray(SHARE_RES(plc).PLC_M, IO_M_NUM));
		return true;
	}

	changed = changed | updateArrayElementIf(cJSON_GetObjectItem(item, "PLC_IN"), SHARE_RES(plc).PLC_IN);
	changed = changed | updateArrayElementIf(cJSON_GetObjectItem(item, "PLC_OUT"), SHARE_RES(plc).PLC_OUT);
	changed = changed | updateArrayElementIf(cJSON_GetObjectItem(item, "PLC_VIN"), SHARE_RES(plc).PLC_VIN);
	changed = changed | updateArrayElementIf(cJSON_GetObjectItem(item, "PLC_VOUT"), SHARE_RES(plc).PLC_VOUT);
	changed = changed | updateArrayElementIf(cJSON_GetObjectItem(item, "PLC_M"), SHARE_RES(plc).PLC_M);
	if(changed == true) {
		return item;
	}
	return NULL;
}

bool updateIntArrayElementIf(cJSON* array, const int* base) {
	size_t i = 0;
	bool changed = false;
	size_t num = cJSON_GetArraySize(array);
	cJSON* temp;
	int now;
	for(i = 0;i < num; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = base[i];
		if ( temp->valuedouble!= (double)now) {
			cJSON_SetNumberValue(temp, (double)now);
			changed = true;
		}
	}
	return changed;
}

bool updateShortArrayElementIf(cJSON* array, const short* base) {
	size_t i = 0;
	bool changed = false;
	size_t num = cJSON_GetArraySize(array);
	cJSON* temp;
	short now;
	for(i = 0;i < num; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = base[i];
		if ( temp->valuedouble!= (double)now) {
			cJSON_SetNumberValue(temp, (double)now);
			changed = true;
		}
	}
	return changed;
}

bool updateDoubleArrayElementIf(cJSON* array, const double* base) {
	size_t i = 0;
	bool changed = false;
	size_t num = cJSON_GetArraySize(array);
	cJSON* temp;
	double now;
	for(i = 0;i < num; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = base[i];
		if ( temp->valuedouble!= now) {
			cJSON_SetNumberValue(temp, now);
			changed = true;
		}
	}
	return changed;
}

bool update2DDoubleArrayElementIf(cJSON* array, const double** base) {
	size_t i = 0, j = 0;
	bool changed = false;
	size_t num = cJSON_GetArraySize(array);
	cJSON* sub_array;
	size_t sub_num;
	cJSON* temp;
	double now;
	for(i = 0;i < num; i++) {
		sub_array = cJSON_GetArrayItem(item, i);
		sub_num = cJSON_GetArraySize(sub_array);
		for (j = 0;j < sub_num; j++) {
			temp = cJSON_GetArrayItem(sub_array, j);
			now = base[i][j];
			if ( temp->valuedouble!= now) {
				cJSON_SetNumberValue(temp, now);
				changed = true;
			}
		}
	}
	return changed;
}

bool getSysvar(cJSON* item) {
	cJSON *array;
	int i=0;
	bool changed = false;
	if (NULL == item) {
		item = cJSON_CreateObject();

		cJSON_AddItemToObject(item, "cRobB", cJSON_CreateIntArray((const int *)SHARE_RES(sysvar).cRobB, B_COUNT));

		array = cJSON_CreateArray();
		for (i=0; i<I_COUNT; i++) {
			cJSON_AddItemToArray(array, cJSON_CreateNumber(SHARE_RES(sysvar).iRobI[i]));
		}
		cJSON_AddItemToObject(item, "iRobI", array);

		cJSON_AddItemToObject(item, "dRobD", cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobD, D_COUNT));

		array = cJSON_CreateArray();
		for (i=0; i<P_COUNT; i++) {
			cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobP[i], AXIS_COUNT));
		}
		cJSON_AddItemToObject(item, "dRobP", array);

		array = cJSON_CreateArray();
		for (i=0; i<V_COUNT; i++) {
			cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(SHARE_RES(sysvar).dRobP[i], VSub_COUNT));
		}
		cJSON_AddItemToObject(item, "dRobV", array);

		return item;
	}

	changed = changed | updateIntArrayElementIf(cJSON_GetObjectItem(item, "cRobB"), SHARE_RES(sysvar).cRobB);
	changed = changed | updateShortArrayElementIf(cJSON_GetObjectItem(item, "iRobI"), SHARE_RES(sysvar).iRobI);
	changed = changed | updateDoubleArrayElementIf(cJSON_GetObjectItem(item, "dRobD"), SHARE_RES(sysvar).dRobD);
	changed = changed | update2DDoubleArrayElementIf(cJSON_GetObjectItem(item, "dRobP"), SHARE_RES(sysvar).dRobP);
	changed = changed | update2DDoubleArrayElementIf(cJSON_GetObjectItem(item, "dRobV"), SHARE_RES(sysvar).dRobV);

	return changed;
}

bool getRobotMode(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(GetRobotMode());
		return true;
	}
	if (item->valuedouble == GetRobotMode()){
		return false;
	}
	cJSON_SetNumberValue(item, GetRobotMode());
	return true;
}

cJSON* getEachLocavar(Robot_LOCVAR base) {
	cJSON *item, *array;
	int i=0;
	item = cJSON_CreateObject();

	cJSON_AddItemToObject(item, "cRobLB", cJSON_CreateIntArray((const int *)base.cRobLB, LB_COUNT));

	cJSON_AddItemToObject(item, "iRobLI", cJSON_CreateIntArray(base.iRobLI, LI_COUNT));

	cJSON_AddItemToObject(item, "dRobLD", cJSON_CreateDoubleArray(base.dRobLD, LD_COUNT));

	array = cJSON_CreateArray();
	for (i=0; i<LP_COUNT; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(base.dRobLP[i], AXIS_COUNT));
	}
	cJSON_AddItemToObject(item, "dRobLP", array);

	array = cJSON_CreateArray();
	for (i=0; i<LV_COUNT; i++) {
		cJSON_AddItemToArray(array, cJSON_CreateDoubleArray(base.dRobLV[i], VSub_COUNT));
	}
	cJSON_AddItemToObject(item, "dRobLV", array);
	return item;
}

cJSON* getLocvar() {
	cJSON *item;
	int i=0;
	item = cJSON_CreateArray();

	for (i=0; i<CALL_NEST_NUM; i++) {
		cJSON_AddItemToArray(item, getEachLocavar(SHARE_RES(locvar)[i]));
	}
	return item;
}

bool getLocvarNum(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(SHARE_RES(locvar_num));
		return true;
	}
	
	if (item->valuedouble == SHARE_RES(locvar_num)) {
		return false;
	}
	item->valuedouble = SHARE_RES(locvar_num);
	return true;
}

bool getAnalogIoInput(cJSON* item) {
	size_t i = 0;
	bool changed = false;
	cJSON* temp;
	if (NULL == item) {
		item = cJSON_CreateDoubleArray(SHARE_RES(analog_ioInput), ANALOG_IN_NUM);
		return true;
	}
	for(i = 0;i < (size_t)ANALOG_IN_NUM; i++) {
		temp = cJSON_GetArrayItem(item, i);
		now = *(SHARE_RES(analog_ioInput) + i)
		if (temp->valuedouble != now) {
			cJSON_SetNumberValue(temp, now);
			changed = true;
		}
	}
	
	return changed;
}

bool getServoDirveMode(cJSON* item) {
	if (NULL == item) {
		item = cJSON_CreateNumber(SHARE_RES(servo_dirve_mode));
		return true;
	}
	if (item->valuedouble == SHARE_RES(servo_dirve_mode)) {
		return false;
	}
	item->valuedouble = SHARE_RES(servo_dirve_mode);
	return true;
}

ResourceItem ResourceTable[] = {
	{"autorun_cycleMode", NULL, &getAutorunCycleMode},
	{"teach_mode", NULL, &getTeachMode},
	{"machinePos", NULL, &getMachinePos},
	{"machinePose", NULL, &getMachinePose},
	{"abs_pulse", NULL, &getAbsPulse},
	{"abz_pulse", NULL, &getAbzPulse},
	{"cur_encode", NULL, &getCurEncode},
	{"robotState", NULL, &getRobotState},
	{"servoReady", NULL, &getServoReady},
	{"currentLine", NULL, &getCurLine},
	{"plc", NULL, &getRobPLC},
	{"sysvar", NULL, &getSysvar},
	{"robotMode", NULL, &getRobotMode},
	//{"locvar", NULL, &getLocvar},
	{"locvar_num", NULL, &getLocvarNum},
	{"analog_ioInput", NULL, &getAnalogIoInput},
	{"servo_dirve_mode", NULL, &getServoDirveMode},
	{"ENDLINE", NULL, NULL},
};