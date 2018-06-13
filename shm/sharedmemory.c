#include "share/resource.h"
#include "crc.h"

int get_press_reset() {
	return PRESS_RESET;
}

int32_t get_zero_encode(const int x) {
	return Get_ZeroEncode(x);
}

void GetAll() {

}

void init_shm() {
	resource_init("/rbctrl/mcserver");
	init_nvram(NULL);
}