#include <stdio.h>
#include <stdlib.h>

#include "mcsql.h"

void test_case1() {
	sql_mapper * mapper = get_arc_sql_mapper(ELIBOT_ARC_GET_ALL_PARAMS);
	register_sql_mapper(mapper);

	char* res = mcsql_query(ELIBOT_ARC_GET_ALL_PARAMS, "/rbctrl/db/elibotDB.db", DB_QUERY_MODE_STANDARD);
	printf("%s\n", res);
}

void test_case2() {
	sql_mapper * mapper = get_arc_sql_mapper(ELIBOT_ARC_GET_PARAMS);
    register_sql_mapper(mapper);

    db_query_req_parameter* param = new_db_query_req_parameter(2);
    param->params[0].name = "file_no";
    param->params[0].type = DB_TYPE_INT32;
    param->params[0].value.int_value = 7;
    param->params[1].name = "group";
    param->params[1].type = DB_TYPE_TEXT;
    param->params[1].value.string_value = "arc.weaving";
    char* res = mcsql_query_with_param(ELIBOT_ARC_GET_PARAMS, "/rbctrl/db/elibotDB.db", DB_QUERY_MODE_CUSTOM);
    printf("%s\n", res);
}

int main(int argc, char const *argv[])
{
	test_case1();
	test_case2();
	return 0;
}