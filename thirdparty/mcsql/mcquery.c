#include "mcquery.h"

sql_parameter* getindexedsqlparam(db_query_req_parameter* req_params, int16_t i) {
    if(i >= req_params->param_size) {
        return NULL;
    }
    return &req_params->params[i];
}

sql_parameter* new_sql_parameter(int16_t size) {
    sql_parameter* params = (sql_parameter*) malloc(sizeof(sql_parameter)*size);
    if (NULL == params) {
        return NULL;
    }
    return params;
}

db_query_req_parameter* new_db_query_req_parameter() {
    db_query_req_parameter* param = (db_query_req_parameter*)malloc(sizeof(db_query_req_parameter));
    if (NULL==param) {
        return NULL;
    }

    param->params = NULL;
    param->param_size = 0;
    return param;
}

db_query_req_page* new_db_query_req_page(int page_start, int page_size) {
    db_query_req_page* page = (db_query_req_page*)malloc(sizeof(db_query_req_page));
    if (NULL==page) {
        return NULL;
    }
    page->page_start = page_start;
    page->page_size = page_size;
    return page;
}
