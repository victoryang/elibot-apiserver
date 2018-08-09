# sqlite3

## import sql file
	sqlite3 *.db ".read *.sql"

## check if table '{table_name}' exist
	SELECT name FROM sqlite_master WHERE type='table' AND name='{table_name}'