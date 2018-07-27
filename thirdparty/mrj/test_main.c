#include <stdio.h>
#include <stdlib.h>

#include "mrj.h"

int main(int argc, char const *argv[])
{
	init_nv_ram();
	mrj_init();

	char* res = get_resource_data();
	char* nv = get_nv_data();
	printf("%s\n", res);
	printf("%s\n", nv);
	return 0;
}