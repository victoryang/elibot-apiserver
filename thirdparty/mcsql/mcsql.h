#ifndef MCSQL_H
#define MCSQL_H

#include "mcquery.h"

void add_int32_to_param(int32_t value, sql_parameter* param);
void add_int64_to_param(int64_t value, sql_parameter* param);
void add_double_to_param(double value, sql_parameter* param);
void add_string_to_param(char* value, sql_parameter* param);

char* mcsql_query_with_param(char* id, char* conn, int32_t type_handle_mode, db_query_req_parameter* parameter, db_query_req_page* page);

char* mcsql_query(char* id, char* conn, int32_t type_handle_mode);

#endif //MCSQL_H