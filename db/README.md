# sqlite3

## import sql file
	sqlite3 *.db ".read *.sql"

## check if table '{table_name}' exist
	SELECT name FROM sqlite_master WHERE type='table' AND name='{table_name}'


## ISSUE: SQLite-database disk image is malformed
	sqlite3 database_name
	PRAGMA integrity_check;
	sqlite>.output tmp.sql
	sqlite>.dump
	sqlite>.quit

	sqlite3 newdb.db
	sqlite>.read tmp.sql
	sqlite>.quit

## ISSUE: SQL error: database is locked
	$ fuser development.db
	> development.db: 5430

	kill -9 5430

## Check sqlite compile option
	PRAGMA compile_options;