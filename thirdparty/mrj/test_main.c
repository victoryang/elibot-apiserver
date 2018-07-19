#include <stdio.h>
#include <stdlib.h>

#include <mrj.h>

int main(int argc, char const *argv[])
{
	resource_init("/rbctrl/mcserver");

	char* res = get_resource_data(data);
	printf("%s\n", res);
	return 0;
}