#include "share/resource.h"
#include "crc.h"
#include "nv.h"

uint32_t crc = 0xFFFFFFFF;

int watch_test() {
	return PRESS_RESET;
}

int32_t get_zero_encode(const int x) {
	return Get_ZeroEncode(x);
}

int get_axis_count (){
	return AXIS_COUNT;
}

int watch_nv() { 
	int c = crc32(crc, (unsigned char *)(void *)t4, sizeof(struct test));
	return c;
}

void init_shm() {
	resource_init("/rbctrl/mcserver");
	init_nvram(NULL);
	crc32_init();
}