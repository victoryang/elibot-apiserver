#ifndef MCQUERY_H
#define MCQUERY_H

#include "define.h"
#include "db/db_query.h"

db_query_req_parameter* new_db_query_req_parameter();
db_query_req_option* new_db_query_req_option(int32_t type_handle_mode);

#endif //MCQUERY_H