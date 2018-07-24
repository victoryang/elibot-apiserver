#include "mcquery.h"

db_query_req_option* new_db_query_req_option(int32_t type_handle_mode) {
    db_query_req_option* option = (db_query_req_option*)malloc(sizeof(db_query_req_option));
    if (NULL==option) {
        return NULL;
    }
    option->type_handle_mode = type_handle_mode;
    return option;
}

sql_parameter* get_sqlparam_index(db_query_req_parameter* req_params, int16_t i) {
    if(i >= req_params->param_size) {
        return NULL;
    }
    return &req_params->params[i];
}

void free_db_query_req_parameter(db_query_req_parameter* p) {
	free(p->params);
	p->param_size = 0;

	free(p);
	p = NULL;
}

db_query_req_parameter* new_db_query_req_parameter(int16_t size) {
    db_query_req_parameter* param = (db_query_req_parameter*)malloc(sizeof(db_query_req_parameter));
    if (NULL==param) {
        return NULL;
    }

    param->params = (sql_parameter*) malloc(sizeof(sql_parameter)*size);
    if (NULL == param->params) {
    	return NULL;
    }
    param->param_size = size;
    return param;
}

db_query_req_page* new_db_query_req_page(int16_t page_start, int16_t page_size) {
    db_query_req_page* page = (db_query_req_page*)malloc(sizeof(db_query_req_page));
    if (NULL==page) {
        return NULL;
    }
    page->page_start = page_start;
    page->page_size = page_size;
    return page;
}
