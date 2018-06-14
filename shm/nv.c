#include "share/resource.h"

#define NV_AXIS_COUNT AXIS_COUNT

int get_axis_count() {
	return NV_AXIS_COUNT;
}

char* get_main_file() {
	return GetMainfile();
}

int32_t get_zero_encode(const int x) {
	return Get_ZeroEncode(x);
}

int32_t get_cur_coordinate() {
	return GetCurCoordinate();
}

void init_nv_ram() {
	init_nvram(NULL);
}