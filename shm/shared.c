#include "share/resource.h"

unsigned char get_autorun_cyclemode() {
	return GET_AUTORUN_CYCLEMODE();
}

int init_shared_resource() {
	return resource_init("/rbctrl/mcserver");
}

