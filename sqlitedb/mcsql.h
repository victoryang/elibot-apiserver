#ifndef MCSQL_H
#define MCSQL_H

#include "mcquery.h"

char* mcsql_query_with_param(char* id, char* conn, int32_t type_handle_mode, db_query_req_parameter* parameter, db_query_req_page* page);

char* mcsql_query(char* id, char* conn, int32_t type_handle_mode);

#endif //MCSQL_H