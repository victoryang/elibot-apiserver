package sqlitedb

// SqlMapper defines methods of a sql mapper
type SqlMapper interface {
	RegisterSqlMapper()		error
	GetID()		string
}

// BaseSqlMapper should be inherited by sql mappers
type BaseSqlMapper struct {

}