#ifndef MCQUERY_H
#define MCQUERY_H

#include "define.h"
#include "db/db_query.h"

db_query_req_option* new_db_query_req_option(int32_t type_handle_mode);

sql_parameter* get_sqlparam_index(db_query_req_parameter* req_params, int16_t i);
db_query_req_parameter* new_db_query_req_parameter(int16_t size);
void free_db_query_req_parameter(db_query_req_parameter* p);

db_query_req_page* new_db_query_req_page(int16_t page_start, int16_t page_size);

#endif //MCQUERY_H