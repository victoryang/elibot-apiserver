#include "share/resource.h"

unsigned char get_autorun_cyclemode() {
	return GET_AUTORUN_CYCLEMODE()
}

void init_shared_resource() {
	resource_init("/rbctrl/mcserver");
}