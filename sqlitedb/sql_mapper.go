package sqlitedb

// #cgo LDFLAGS: -lsqlitedb
// #include<stdlib.h>
// #include<sql_mapper.h>
import "C"
import (
	"errors"
)

type SqlMapper struct {
	m   *C.sql_mapper
}

func Get_bookprogram_sql_mapper() (*SqlMapper, error) {
	mapper := C.get_bookprogram_sql_mapper()
	if mapper == 0 {
		return nil, errors.Error("Getting sqlmapper fails")
	}
	//sql.Mapper = mapper
	C.register_sql_mapper(mapper.m)
	return mapper, nil
}