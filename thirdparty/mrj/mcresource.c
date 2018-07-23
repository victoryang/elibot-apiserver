#include "mcresource.h"

cJSON* getAutorunCycleMode() {
	cJSON* item = cJSON_CreateNumber(GET_AUTORUN_CYCLEMODE());
	return item;
}

cJSON* getTeachMode() {
	cJSON* item = cJSON_CreateNumber(GET_AUTORUN_CYCLEMODE());
	return item;
}

cJSON* getMachinePos() {
	cJSON* item = cJSON_CreateDoubleArray(RobotRes_MachPos_base, AXIS_COUNT);
	return item;
}

cJSON* getMachinePose() {
	cJSON* item = cJSON_CreateDoubleArray(RobotRes_MachPose_base, 6);
	return item;
}

cJSON* getAbsPulse() {
	cJSON* item = cJSON_CreateIntArray((const int *)SHARE_RES(abs_pulse), AXIS_COUNT);
	return item;
}

cJSON* getAbzPulse() {
	cJSON* item = cJSON_CreateIntArray((const int *)SHARE_RES(abz_pulse), AXIS_COUNT);
	return item;
}

cJSON* getCurEncode() {
	cJSON* item = cJSON_CreateIntArray((const int *)SHARE_RES(cur_encode), AXIS_COUNT);
	return item;
}

cJSON* getRobotState() {
	cJSON* item = cJSON_CreateNumber(GetRobotState());
	return item;
}

cJSON* getServoReady() {
	cJSON* item = cJSON_CreateNumber(GetServoReady());
	return item;
}

cJSON* getCurLine() {
	cJSON* item = cJSON_CreateNumber(GetCurLine());
	return item;
}


cJSON* convertToIntArray(volatile uint8_t* base, int array_size) {
	int i=0;
	cJSON* array = cJSON_CreateArray();
	for (;i<array_size;i++) {
		cJSON_AddItemToArray(array, cJSON_CreateNumber((double)base[i]));
	}
	return array;
}
cJSON* getRobPLC() {
	cJSON* item = cJSON_CreateObject();
	cJSON_AddItemToObject(item, "PLC_IN", convertToIntArray(SHARE_RES(plc).PLC_IN, IO_IN_NUM));
	cJSON_AddItemToObject(item, "PLC_OUT", convertToIntArray(SHARE_RES(plc).PLC_OUT, IO_OUT_NUM));
	cJSON_AddItemToObject(item, "PLC_VIN", convertToIntArray(SHARE_RES(plc).PLC_VIN, IO_VIN_NUM));
	cJSON_AddItemToObject(item, "PLC_VOUT", convertToIntArray(SHARE_RES(plc).PLC_VOUT, IO_VOUT_NUM));
	cJSON_AddItemToObject(item, "PLC_M", convertToIntArray(SHARE_RES(plc).PLC_M, IO_M_NUM));
	return item;
}

cJSON* getSysvar() {
	cJSON *item, *array;
	int i=0;
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

cJSON* getRobotMode() {
	cJSON* item = cJSON_CreateNumber(GetRobotMode());
	return item;
}

cJSON* getEachLocavar(Robot_LOCVAR base) {
	cJSON *item, *array;
	int i=0;
	item = cJSON_CreateObject();

	cJSON_AddItemToObject(item, "cRobLB", cJSON_CreateIntArray(base.cRobLB, LB_COUNT));

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
	cJSON *item, *array;
	int i=0;
	item = cJSON_CreateObject();

	for (i=0; i<CALL_NEST_NUM; i++) {
		cJSON_AddItemToArray(array, getEachLocavar(SHARE_RES(locvar)[i]));
	}
	cJSON_AddItemToObject(item, "locvar", array);
	return item;
}

cJSON* getLocvarNum() {
	cJSON* item = cJSON_CreateNumber(SHARE_RES(locvar_num));
	return item;
}

cJSON* getAnalogIoInput() {
	cJSON* item = cJSON_CreateDoubleArray(SHARE_RES(analog_ioInput), ANALOG_IN_NUM);
	return item;
}

cJSON* getServoDirveMode() {
	cJSON* item = cJSON_CreateNumber(SHARE_RES(servo_dirve_mode));
	return item;
}

Item ResourceTable[] = {
	{"autorun_cycleMode", &getAutorunCycleMode},
	{"teach_mode", &getTeachMode},
	{"machinePos", &getMachinePos},
	{"machinePose", &getMachinePose},
	{"abs_pulse", &getAbsPulse},
	{"abz_pulse", &getAbzPulse},
	{"cur_encode", &getCurEncode},
	{"robotState", &getRobotState},
	{"servoReady", &getServoReady},
	{"currentLine", &getCurLine},
	{"plc", &getRobPLC},
	{"sysvar", &getSysvar},
	{"robotMode", &getRobotMode},
	{"locvar", &getLocvar},
	{"locvar_num", &getLocvarNum},
	{"analog_ioInput", &getAnalogIoInput},
	{"servo_dirve_mode", &getServoDirveMode},
	{"ENDLINE", NULL},
};