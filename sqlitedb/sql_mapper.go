package sqlitedb

// SqlMapper defines methods of a sql mapper
type SqlMapper interface {
	RegisterSqlMapper() error
}

// BaseSqlMapper should be inherited by sql mappers
type BaseSqlMapper struct {
	id        string
}