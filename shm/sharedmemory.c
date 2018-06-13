#include "share/resource.h"
#include "crc.h"

int watch_test() {
	return PRESS_RESET;
}

int32_t get_zero_encode(const int x) {
	return Get_ZeroEncode(x);
}

unsigned long watch_nv() {
	nvram_data 
	return NV_DATA_MAGIC;
}

void init_shm() {
	resource_init("/rbctrl/mcserver");
	init_nvram(NULL);
}