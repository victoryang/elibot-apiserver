#include <stdio.h>
#include <stdlib.h>

#include <mcsql.h>

int main(int argc, char const *argv[])
{
	char* res = mcsql_query(ELIBOT_ARC_GET_ALL_PARAMS, "/rbctrl/db/elibotDB.db");
	printf("%s\n", res);
	return 0;
}