#include "mcsql.h"

void add_int32_to_param(int32_t value, sql_parameter* param) {
	(*param).type = DB_TYPE_INT32;
	(*param).value.int_value = value;
}

void add_int64_to_param(int64_t value, sql_parameter* param) {
	(*param).type = DB_TYPE_INT64;
	(*param).value.int64_value = value;
}

void add_double_to_param(double value, sql_parameter* param) {
	(*param).type = DB_TYPE_DOUBLE;
	(*param).value.double_value = value;
}

void add_string_to_param(char* value, sql_parameter* param) {
	(*param).type = DB_TYPE_TEXT;
	(*param).value.string_value = value;
}

char* mcsql_query_with_param(char* id, char* conn, int32_t type_handle_mode, db_query_req_parameter* parameter, db_query_req_page* page) {
	db_query_req_option opt = {
            type_handle_mode:type_handle_mode
    };

	db_query_req req = {
		query_id: id,
		conn_str: conn,
		option: &opt,
		parameter: parameter,
		page: page,
	};

	cJSON* root = db_query(&req);

	char *ret = cJSON_PrintUnformatted(root);
	cJSON_Delete(root);
	return ret;
}

char* mcsql_query(char* id, char* conn, int32_t type_handle_mode) {
	return mcsql_query_with_param(id, conn, type_handle_mode, NULL, NULL);
}