#include "share/resource.h"

int get_press_reset() {
	return PRESS_RESET;
}

int32_t get_zero_encode(const int x) {
	return Get_ZeroEncode(x);
}

void init_shm() {
	resource_init("/rbctrl/");
}