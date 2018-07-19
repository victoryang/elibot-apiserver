#include "mcsql.h"

char* mcsql_query(char* id, char* conn) {
	db_query_req_parameter* req_param = new_db_query_req_parameter();
	db_query_req_option* req_option =  new_db_query_req_option(DB_QUERY_MODE_STANDARD);
	db_query_req req = db_query_req{id, conn, req_option, NULL, NULL};
	cJSON* root = db_query(&req);
	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret
}