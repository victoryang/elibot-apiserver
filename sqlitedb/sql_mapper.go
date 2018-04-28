package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<db/db_context.h>

import "C"
import (
	"errors"
)

/*type SqlMapper struct {
	Mapper   *C.sql_mapper
}*/

func Get_bookprogram_sql_mapper() (*SqlMapper, error) {
	mapper := C.get_bookprogram_sql_mapper()
	if mapper == 0 {
		return nil, errors.Error("Getting sqlmapper fails")
	}
	//sql.Mapper = mapper
	C.register_sql_mapper(mapper)
	return mapper, nil
}