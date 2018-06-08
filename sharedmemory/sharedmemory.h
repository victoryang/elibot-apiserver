#include "share/resource.h"

int get_press_reset() {
	return PRESS_RESET;
}

int32_t get_cur_line() {
	return GetCurLine();
}

void init_shm() {
	resource_init("/rbctrl/");
}