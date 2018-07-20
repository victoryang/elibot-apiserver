#include "mcsql.h"

void add_int32_to_param(int32_t value, db_query_req_parameter* param) {
	(*param).type = DB_TYPE_INT32;
	(*param).value.int_value = value;
}

void add_int64_to_param(int64_t value, db_query_req_parameter* param) {
	(*param).type = DB_TYPE_INT64;
	(*param).value.int64_value = value;
}

void add_double_to_param(double value, db_query_req_parameter* param) {
	(*param).type = DB_TYPE_DOUBLE;
	(*param).value.double_value = value;
}

void add_string_to_param(char* value, db_query_req_parameter* param) {
	(*param).type = DB_TYPE_TEXT;
	(*param).value.string_value = value;
}

char* mcsql_query_with_param(char* id, char* conn, int32_t type_handle_mode, db_query_req_parameter* parameter, db_query_req_page* page) {
	db_query_req* req = (db_query_req*) malloc(sizeof(db_query_req));
	if (NULL == req) {
		return NULL;
	}
	req->query_id = id;
	req->conn_str = conn;

	db_query_req_option* option = new_db_query_req_option(type_handle_mode);
	if (NULL == option) {
		return NULL;
	}

	req->parameter = parameter;
	req->page = page;

	cJSON* root = db_query(req);

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

char* mcsql_query(char* id, char* conn, int32_t type_handle_mode) {
	return mcsql_query_with_param(id, conn, type_handle_mode, NULL, NULL);
}