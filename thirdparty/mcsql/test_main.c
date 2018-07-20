#include <stdio.h>
#include <stdlib.h>

#include "mcsql.h"

int main(int argc, char const *argv[])
{
	sql_mapper * mapper = get_arc_sql_mapper(ELIBOT_ARC_GET_ALL_PARAMS);
    register_sql_mapper(mapper);
   
	char* res = mcsql_query(ELIBOT_ARC_GET_ALL_PARAMS, "/rbctrl/db/elibotDB.db");
	printf("%s\n", res);
	return 0;
}