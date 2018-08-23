#include "mcsqlmapper.h"

void register_arc_sql_mappers() {
	// ELIBOT_ARC_GET_PARAMS
    sql_mapper *mapper = get_arc_sql_mapper(ELIBOT_ARC_GET_PARAMS);
    register_sql_mapper(mapper);

    // ELIBOT_ARC_GET_ALL_PARAMS
    mapper = get_arc_sql_mapper(ELIBOT_ARC_GET_ALL_PARAMS);
    register_sql_mapper(mapper);
};

void register_backup_sql_mappers() {
    sql_mapper *mapper = get_bak_sql_mapper(ELIBOT_BAK_BACKUP_PARAMS);
    register_sql_mapper(mapper);
};

void register_bookprogram_sql_mappers() {

    // ELIBOT_BOOKPROGRAM_GET_ALL
    sql_mapper *mapper = get_bookprogram_sql_mapper(ELIBOT_BOOKPROGRAM_GET_ALL);
    register_sql_mapper(mapper);
};

void register_enum_sql_mappers() {
    sql_mapper *mapper = get_enum_sql_mapper(ELIBOT_ENUM_GET_ALL);
    register_sql_mapper(mapper);
};

void register_extaxis_sql_mappers() {
    sql_mapper *mapper = get_extaxis_sql_mapper(ELIBOT_EXTAXIS_GET_ALL);
    register_sql_mapper(mapper);
};

void register_interference_sql_mappers() {
    sql_mapper *mapper = get_interference_sql_mapper(ELIBOT_INTERFERENCE_GET_ALL);
    register_sql_mapper(mapper);
};

void register_io_sql_mappers() {
    sql_mapper *mapper = get_io_sql_mapper(ELIBOT_IO_GET_VALID_IOS_BY_GROUP);
    register_sql_mapper(mapper);
};

void register_metadata_sql_mappers() {
    sql_mapper *mapper = get_metadata_sql_mapper(ELIBOT_METADATA_GET_ALL);
    register_sql_mapper(mapper);
};

void register_params_sql_mappers() {
    // ELIBOT_PARAMS_GET_PARAMS
    sql_mapper *mapper = get_params_sql_mapper(ELIBOT_PARAMS_GET_PARAMS);
    register_sql_mapper(mapper);

    // ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID
    mapper = get_params_sql_mapper(ELIBOT_PARAMS_GET_VALID_PARAM_BY_ID);
    register_sql_mapper(mapper);

    // ELIBOT_PARAMS_GET_VALID_PARAMS_BY_GROUP
    mapper = get_params_sql_mapper(ELIBOT_PARAMS_GET_VALID_PARAMS_BY_GROUP);
    register_sql_mapper(mapper);
};

void register_ref_sql_mapper() {
	sql_mapper *mapper = get_ref_sql_mapper(ELIBOT_REF_GET_ALL);
    register_sql_mapper(mapper);
};

void register_toolframe_sql_mappers() {

    // ELIBOT_COMMON_ALL_TOOLFRAME
    sql_mapper *mapper = get_common_toolframe_sql_mapper(ELIBOT_COMMON_GET_ALL_TOOLFRAMES);
    register_sql_mapper(mapper);

    // ELIBOT_COMMMON_GET_ALL_TOOLPOS
    mapper = get_common_toolframe_sql_mapper(ELIBOT_COMMON_GET_ALL_TOOLPOS);
    register_sql_mapper(mapper);
};

void register_userframe_sql_mappers() {
    sql_mapper *mapper = get_userframe_sql_mapper(ELIBOT_USER_FRAME_GET_ALL);
    register_sql_mapper(mapper);
};

void register_zeropoint_sql_mappers() {
    sql_mapper *mapper = get_zeropoint_sql_mapper(ELIBOT_ZEROPOINT_GET_ALL);
    register_sql_mapper(mapper);
};

void register_all_sql_mappers() {
	register_arc_sql_mappers();
	register_backup_sql_mappers();
	register_bookprogram_sql_mappers();
	register_enum_sql_mappers();
	register_extaxis_sql_mappers();
	register_interference_sql_mappers();
	register_io_sql_mappers();
	register_metadata_sql_mappers();
	register_params_sql_mappers();
	register_ref_sql_mapper();
	register_toolframe_sql_mappers();
	register_userframe_sql_mappers();
	register_zeropoint_sql_mappers();
}