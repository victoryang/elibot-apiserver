#include "share/resource.h"

#define NV_AXIS_COUNT AXIS_COUNT

int get_axis_count() {
	return NV_AXIS_COUNT;
}

char* get_main_file() {
	return GetMainfile();
}

int32_t get_sys_technics() {
	return getTechnics()
}

int32_t get_cur_coordinate() {
	return GetCurCoordinate();
}

int32_t get_mannual_speed_rate() {
	return GetManualSpeedRate();
}

int32_t get_cycle_mode() {
	return GetCycleMode();
}

int32_t get_tool_number() {
	return GetToolNumber();
}

int32_t get_user_number() {
	return GetUserNumber();
}

int32_t get_zero_encode(const int x) {
	return Get_ZeroEncode(x);
}

double get_system_period() {
	return GetSysPeriod();
}

int32_t get_sys_ctrl_mode() {
	return GetSysCtrlMode();
}

double get_origin(const int x) {
	return Get_RobotRes_Origin(x);
}

int init_nv_ram() {
	return init_nvram(NULL);
}