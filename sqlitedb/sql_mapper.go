package sqlitedb

// SqlMapper defines methods of a sql mapper
type SqlMapper interface {
	RegisterSqlMapper(mode int)		error
	GetID()		string
}

const (
	ELIBOT_GET_ALL_PARAMS = 0
)

// BaseSqlMapper should be inherited by sql mappers
type BaseSqlMapper struct {

}