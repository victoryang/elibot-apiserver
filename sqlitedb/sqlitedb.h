#ifndef SQLITEDB_H
#define SQLITEDB_H

#include "stdlib.h"
#include "string.h"
#include "db/db_query.h"
#include "db/db_manager.h"
#include "db/sqlmapper/arc_sql_mapper.h"
#include "db/sqlmapper/backup_sql_mapper.h"
#include "db/sqlmapper/bookprogram_sql_mapper.h"
#include "db/sqlmapper/dynamic_sql_mapper.h"
#include "db/sqlmapper/enum_sql_mapper.h"
#include "db/sqlmapper/extaxis_sql_mapper.h"
#include "db/sqlmapper/interference_sql_mapper.h"
#include "db/sqlmapper/io_sql_mapper.h"
#include "db/sqlmapper/metadata_sql_mapper.h"
#include "db/sqlmapper/parameter_sql_mapper.h"
#include "db/sqlmapper/ref_sql_mapper.h"
#include "db/sqlmapper/toolframe_sql_mapper.h"
#include "db/sqlmapper/userframe_sql_mapper.h"
#include "db/sqlmapper/zeropoint_sql_mapper.h"
#include "db/sqlmapper/recordoperation_sql_mapper.h"

void mcsql_set_db_file(char* dbname);

void register_all_sql_mappers();

char* mcsql_arc_get_all();

char* mcsql_arc_get_params(int32_t file_no, char* group);

char* mcsql_bookprogram_get_all();

char* mcsql_dynamics_get_all();

char* mcsql_dynamics_get_by_id(char* id);

char* mcsql_enum_get_all();

char* mcsql_extaxis_get_all();

char* mcsql_interference_get_all();

char* mcsql_ios_get_all(char* group, char* lang, int32_t auth, int32_t tech);

char* mcsql_metadata_get_all(char* lang);

char* mcsql_operation_record_get_all(int32_t created_time, int32_t start, int32_t page_size);

char* mcsql_params_get_params();

char* mcsql_params_get_valid_param_by_id(char* md_id);

char* mcsql_params_get_valid_param_by_group(char* group);

char* mcsql_ref_get_all();

char* mcsql_toolframe_get_all();

char* mcsql_toolframe_get_by_toolno(int32_t tool_no);

char* mcsql_userframe_get_all();

char* mcsql_userframe_get_by_userno(int32_t user_no);

char* mcsql_zeropoint_get_all();

#endif //SQLITEDB_H