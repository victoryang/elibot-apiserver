#ifndef MCSQLMAPPER_H
#define MCSQLMAPPER_H

#include "define.h"
#include "db/sqlmapper/arc_sql_mapper.h"
#include "db/sqlmapper/backup_sql_mapper.h"
#include "db/sqlmapper/bookprogram_sql_mapper.h"
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

void register_all_sql_mappers();

#endif //MCSQLMAPPER_H