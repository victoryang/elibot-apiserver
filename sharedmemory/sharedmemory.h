#include "share/resource.h"

int get_press_reset() {
	return PRESS_RESET;
}

int32_t get_zero_encode() {
	return Get_ZeroEncode(0);
}

void init_shm() {
	resource_init("/rbctrl/");
}