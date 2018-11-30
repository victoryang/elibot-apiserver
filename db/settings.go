package db

import (
	"database/sql"

	Log "elibot-apiserver/log"
)

func iterateSettings(rows *sql.Rows) map[string]string {
	maps := make(map[string]string)
	for rows.Next() {
		var key string
		var value string
		if err := rows.Scan(&key, &value); err!=nil {
			Log.Error("scan rows fails: ", err)
			continue
		}
		maps[key] = value
	}
	return maps
}

func DoQueryCommand(command string, args ...interface{}) (map[string]string, error){
	mu.Lock()
	defer mu.Unlock()
	rows, err := conn.Query(command, args...)
	if err!=nil {
		Log.Error("failed to query: ", err)
		return nil, err
	}
	defer rows.Close()

	return iterateSettings(rows), nil
}