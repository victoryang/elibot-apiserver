package sqlitedb

type SqlMapper interface {
	RegisterSqlMapper(id string) ()
}

type BaseSqlMapper struct {
}