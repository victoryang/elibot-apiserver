#include "share/resource.h"
#include "crc.h"

uint32_t crc = 0xFFFFFFFF;


int watch_test() {
	return PRESS_RESET;
}

int get_axis_count (){
	return AXIS_COUNT;
}

void init_worker_resource() {
	crc32_init();
}