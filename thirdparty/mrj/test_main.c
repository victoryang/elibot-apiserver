#include <stdio.h>
#include <stdlib.h>

#include <mrj.h>

int main(int argc, char const *argv[])
{
	resource_init("/rbctrl/mcserver");
	cJSON* data = get_resource_data();

	char* res = cJSON_Print(data);
	printf("%s\n", res);
    cJSON_Delete(data);
	return 0;
}