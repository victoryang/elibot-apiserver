#include "sqlitedb.h"

static char db_conn[128];

void mcsql_set_db_file(char* dbname) {
    strcpy(db_conn, dbname);
}

static void register_arc_sql_mappers() {
    // ELIBOT_ARC_GET_PARAMS
    sql_mapper *mapper = get_arc_sql_mapper(ELIBOT_ARC_GET_PARAMS);
    register_sql_mapper(mapper);

    // ELIBOT_ARC_GET_ALL_PARAMS
    mapper = get_arc_sql_mapper(ELIBOT_ARC_GET_ALL_PARAMS);
    register_sql_mapper(mapper);
}

static void register_backup_sql_mappers() {
    sql_mapper *mapper = get_bak_sql_mapper(ELIBOT_BAK_BACKUP_PARAMS);
    register_sql_mapper(mapper);
}

static void register_bookprogram_sql_mappers() {

    // ELIBOT_BOOKPROGRAM_GET_ALL
    sql_mapper *mapper = get_bookprogram_sql_mapper(ELIBOT_BOOKPROGRAM_GET_ALL);
    register_sql_mapper(mapper);
}

static void register_dynamics_sql_mappers() {

    // ELIBOT_BOOKPROGRAM_GET_ALL
    sql_mapper *mapper = get_dynamic_sql_mapper(ELIBOT_DYNAMIC_GET_ALL);
    register_sql_mapper(mapper);

    mapper = get_dynamic_sql_mapper(ELIBOT_DYNAMIC_GET_BY_ID);
    register_sql_mapper(mapper);
}

static void register_enum_sql_mappers() {
    sql_mapper *mapper = get_enum_sql_mapper(ELIBOT_ENUM_GET_ALL);
    register_sql_mapper(mapper);
}

static void register_extaxis_sql_mappers() {
    sql_mapper *mapper = get_extaxis_sql_mapper(ELIBOT_EXTAXIS_GET_ALL);
    register_sql_mapper(mapper);
}

static void register_interference_sql_mappers() {
    sql_mapper *mapper = get_interference_sql_mapper(ELIBOT_INTERFERENCE_GET_ALL);
    register_sql_mapper(mapper);
}

static void register_io_sql_mappers() {
    sql_mapper *mapper = get_io_sql_mapper(ELIBOT_IO_GET_VALID_IOS_BY_GROUP);
    register_sql_mapper(mapper);
}

static void register_metadata_sql_mappers() {
    sql_mapper *mapper = get_metadata_sql_mapper(ELIBOT_METADATA_GET_ALL);
    register_sql_mapper(mapper);
}

static void register_params_sql_mappers() {
    // ELIBOT_PARAMS_GET_PARAMS
    sql_mapper *mapper = get_params_sql_mapper(ELIBOT_PARAMS_GET_PARAMS);
    register_sql_mapper(mapper);

    // ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID
    mapper = get_params_sql_mapper(ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID);
    register_sql_mapper(mapper);

    // ELIBOT_PARAMS_GET_VALID_PARAMS_BY_GROUP
    mapper = get_params_sql_mapper(ELIBOT_PARAMS_GET_VALID_PARAMS_BY_GROUP);
    register_sql_mapper(mapper);
}

static void register_record_op_sql_mapper() {
    sql_mapper *mapper = get_record_sql_mapper(ELIBOT_RECORD_GET_ALL);
    register_sql_mapper(mapper);
}

static void register_ref_sql_mapper() {
    sql_mapper *mapper = get_ref_sql_mapper(ELIBOT_REF_GET_ALL);
    register_sql_mapper(mapper);
}

static void register_toolframe_sql_mappers() {

    // ELIBOT_COMMON_ALL_TOOLFRAME
    sql_mapper *mapper = get_common_toolframe_sql_mapper(ELIBOT_COMMON_GET_ALL_TOOLFRAMES);
    register_sql_mapper(mapper);

    // ELIBOT_COMMON_GET_TOOLFRAMES
    mapper = get_common_toolframe_sql_mapper(ELIBOT_COMMON_GET_TOOLFRAMES);
    register_sql_mapper(mapper);
}

static void register_userframe_sql_mappers() {
    sql_mapper *mapper = get_userframe_sql_mapper(ELIBOT_USER_FRAME_GET_ALL);
    register_sql_mapper(mapper);

    mapper = get_userframe_sql_mapper(ELIBOT_USER_FRAME_GET_BY_USER_NO);
    register_sql_mapper(mapper);
}

static void register_zeropoint_sql_mappers() {
    sql_mapper *mapper = get_zeropoint_sql_mapper(ELIBOT_ZEROPOINT_GET_ALL);
    register_sql_mapper(mapper);
}

void register_all_sql_mappers() {
    register_arc_sql_mappers();
    register_backup_sql_mappers();
    register_bookprogram_sql_mappers();
    register_dynamics_sql_mappers();
    register_enum_sql_mappers();
    register_extaxis_sql_mappers();
    register_interference_sql_mappers();
    register_io_sql_mappers();
    register_metadata_sql_mappers();
    register_params_sql_mappers();
    register_record_op_sql_mapper();
    register_ref_sql_mapper();
    register_toolframe_sql_mappers();
    register_userframe_sql_mappers();
    register_zeropoint_sql_mappers();
}

static char* mcsql_db_query(db_query_req* req) {
    cJSON* root = db_query(req);

    char *ret = cJSON_PrintUnformatted(root);
    cJSON_Delete(root);

    return ret;
}

char* mcsql_arc_get_all() {
    const char *q_id = ELIBOT_ARC_GET_ALL_PARAMS;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_arc_get_params(int32_t file_no, char* group) {
    const char *q_id = ELIBOT_ARC_GET_PARAMS;

    sql_parameter sql_params[] = {
            {name:"file_no", value:{ int_value: file_no}, type:DB_TYPE_INT32},
            {name:"group", value:{ string_value: group}, type:DB_TYPE_TEXT},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 2
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_OBJECT
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_bookprogram_get_all() {
    const char *q_id = ELIBOT_BOOKPROGRAM_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_dynamics_get_all() {
    const char *q_id = ELIBOT_DYNAMIC_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_dynamics_get_by_id(char* id) {
    const char *q_id = ELIBOT_DYNAMIC_GET_BY_ID;

    sql_parameter sql_params[] = {
    {name:"md_id", value:{ string_value: id}, type:DB_TYPE_TEXT}
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 1
    };

    db_query_req_option opt = {
        type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_enum_get_all() {
    const char *q_id = ELIBOT_ENUM_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_extaxis_get_all() {
    const char *q_id = ELIBOT_EXTAXIS_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_OBJECT
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_interference_get_all() {
    const char *q_id = ELIBOT_INTERFERENCE_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_ios_get_all(char* group, char* lang, int32_t auth, int32_t tech) {
    const char *q_id = ELIBOT_IO_GET_VALID_IOS_BY_GROUP;

    sql_parameter sql_params[] = {
            {name:"group", value:{ string_value: group}, type:DB_TYPE_TEXT},
            {name:"lang", value:{ string_value: lang}, type:DB_TYPE_TEXT},
            {name:"auth", value:{ int_value: auth}, type:DB_TYPE_INT32},
            {name:"tech", value:{ int_value: tech}, type:DB_TYPE_INT32},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 4
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_metadata_get_all(char* lang) {
    const char *q_id = ELIBOT_METADATA_GET_ALL;

    sql_parameter sql_params[] = {
            {name:"lang", value:{ string_value: lang}, type:DB_TYPE_TEXT},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 1
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_OBJECT
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_operation_record_get_all(int32_t created_time, int32_t start, int32_t page_size) {
    const char *q_id = ELIBOT_RECORD_GET_ALL;

    sql_parameter sql_params[] = {
            {name:"created_time", value:{ int_value: created_time}, type:DB_TYPE_INT32},
            {name:"start", value:{int_value: start}, type:DB_TYPE_INT32},
            {name:"pageSize", value:{int_value: page_size}, type: DB_TYPE_INT32},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 3
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_params_get_params() {
    const char *q_id = ELIBOT_PARAMS_GET_PARAMS;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_params_get_valid_param_by_id(char* md_id) {
    const char *q_id = ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID;

    sql_parameter sql_params[] = {
            {name:"md_id", value:{ string_value: md_id}, type:DB_TYPE_TEXT},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 1
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_OBJECT
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_params_get_valid_param_by_group(char* group) {
    const char *q_id = ELIBOT_PARAMS_GET_VALID_PARAMS_BY_GROUP;

    sql_parameter sql_params[] = {
            {name:"group", value:{ string_value: group}, type:DB_TYPE_TEXT},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 1
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_OBJECT
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_ref_get_all() {
    const char *q_id = ELIBOT_REF_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_QUERY
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_toolframe_get_all() {
    const char *q_id = ELIBOT_COMMON_GET_ALL_TOOLFRAMES;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_toolframe_get_by_toolno(int32_t tool_no) {
    const char *q_id = ELIBOT_COMMON_GET_TOOLFRAMES;

    sql_parameter sql_params[] = {
            {name:"tool_no", value:{ int_value: tool_no}, type:DB_TYPE_INT32},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 1
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_OBJECT
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_userframe_get_all() {
    const char *q_id = ELIBOT_USER_FRAME_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_userframe_get_by_userno(int32_t user_no) {
    const char *q_id = ELIBOT_USER_FRAME_GET_BY_USER_NO;

    sql_parameter sql_params[] = {
            {name:"user_no", value:{ int_value: user_no}, type:DB_TYPE_INT32},
    };

    db_query_req_parameter q_params = {
            params: sql_params,
            param_size: 1
    };

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_CUSTOM_OBJECT
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:&q_params,
            page:NULL,
    };

    return mcsql_db_query(&req);
}

char* mcsql_zeropoint_get_all() {
    const char *q_id = ELIBOT_ZEROPOINT_GET_ALL;

    db_query_req_option opt = {
            type_handle_mode:DB_QUERY_MODE_STANDARD
    };

    db_query_req req = {
            query_id:(char *)q_id,
            conn_str:db_conn,
            option:&opt,
            parameter:NULL,
            page:NULL,
    };

    return mcsql_db_query(&req);
}