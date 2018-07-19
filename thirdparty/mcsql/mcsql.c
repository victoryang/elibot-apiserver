#include "mcsql.h"

char* mcsql_query(char* id, char* conn) {
	db_query_req* req = (db_query_req*) malloc(sizeof(db_query_req));
	if (NULL == req) {
		return NULL;
	}
	req->query_id = id;
	req->conn_str = conn;

	struct db_query_req_option option = {DB_QUERY_MODE_STANDARD};
	req->option = &option;

	req->parameter = NULL;
	req->page = NULL;

	cJSON* root = db_query(req);

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}