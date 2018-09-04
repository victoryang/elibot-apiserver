#include "mcresource.h"

cJSON* get_resource() {
	cJSON* item;
	item = cJSON_CreateObject();
	cJSON_AddItemToObject(item, "autorun_cycleMode", cJSON_CreateNumber(GET_AUTORUN_CYCLEMODE()));
	cJSON_AddItemToObject(item, "teach_mode", cJSON_CreateNumber(GET_TEACH_MODE()));
	cJSON_AddItemToObject(item, "machinePos", cJSON_CreateDoubleArray(RobotRes_MachPos_base, AXIS_COUNT));
	cJSON_AddItemToObject(item, "machinePose", cJSON_CreateDoubleArray(RobotRes_MachPose_base, 6));
	cJSON_AddItemToObject(item, "abs_pulse", cJSON_CreateIntArray((const int *)SHARE_RES(abs_pulse), AXIS_COUNT));
	cJSON_AddItemToObject(item, "abz_pulse", cJSON_CreateIntArray((const int *)SHARE_RES(abz_pulse), AXIS_COUNT));
	cJSON_AddItemToObject(item, "cur_encode", cJSON_CreateIntArray((const int *)SHARE_RES(cur_encode), AXIS_COUNT));
	cJSON_AddItemToObject(item, "robotState", cJSON_CreateNumber(GetRobotState()));
	cJSON_AddItemToObject(item, "servoEnabled", cJSON_CreateNumber(GetServoReady()));
	cJSON_AddItemToObject(item, "canMotorRun", cJSON_CreateNumber(GETMOTOR_RUN_STATE()));
	cJSON_AddItemToObject(item, "currentLine", cJSON_CreateNumber(GetCurLine()));
	cJSON_AddItemToObject(item, "robotMode", cJSON_CreateNumber(GetRobotMode()));
	cJSON_AddItemToObject(item, "locvar_num", cJSON_CreateNumber(SHARE_RES(locvar_num)));
	cJSON_AddItemToObject(item, "analog_ioInput", cJSON_CreateDoubleArray(SHARE_RES(analog_ioInput), ANALOG_IN_NUM));
	cJSON_AddItemToObject(item, "servo_dirve_mode", cJSON_CreateNumber(SHARE_RES(servo_dirve_mode)));
	cJSON_AddItemToObject(item, "speed_modify_play", cJSON_CreateNumber(NVRAM_PARA(speed_modify_play)));
	return item;
}