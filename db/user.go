package db

import (
	Log "elibot-apiserver/log"
)

type User struct {
	Name 			string		`json:"name,omitempty"`
	Mail 			string 		`json:"mail,omitempty"`
	Mobile			string 		`json:"mobile,omitempty"`
	Authority		int 		`json:"authority"`
}

type Shadow struct {
	Name 		string
	SecretKey 	string
}

var user_table = "elt_users"
var shadow_table = "elt_shadow"
var userAddCommand = "INSERT INTO " + user_table + "(name, mail, mobile, authority) VALUES(?,?,?,?)"
var userRemoveCommand = "DELETE FROM " + user_table + " WHERE name = ?"
var secretAddCommand = "INSERT INTO " + shadow_table + " VALUES(?,?)"
var secretRemoveCommand = "DELETE FROM " + shadow_table + " WHERE name = ?"

func InitUserTable() error {
	command := "CREATE TABLE IF NOT EXISTS " + user_table + "(name TEXT PRIMARY KEY NOT NULL, mail TEXT, mobile TEXT,authority INTEGER)"

    return CreateTableIfNotExist(command)
}

func GetUsersList() []User {
	command := "SELECT * FROM " + user_table
	rows, err := DoQuery(command)
	if err!=nil {
		Log.Error("Could not get user list")
		return nil
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Name, &u.Mail, &u.Mobile, &u.Authority); err!=nil {
			Log.Error("scan rows fails: ", err)
			continue
		}

		users = append(users, u)
	}

	return users
}

func ModifyUser(u User) error {
	command := "UPDATE " + user_table + "SET mail = ?, mobile = ?, authority = ? WHERE name = ?"
	return PrepareAndExecuteCommand(command, u.Mail, u.Mobile, u.Authority, u.Name)
}

func AddUserWithSecretkey(name string, mail string, mobile string, authority int, secretkey string) bool {
	mu.Lock()
	defer mu.Unlock()

	if name == "" || secretkey == "" {
		return false
	}

	tx, err := conn.Begin()
	if err!=nil {
		Log.Error("failed to begin transaction")
		return false
	}
	defer tx.Rollback()

	stmtU, err1 := tx.Prepare(userAddCommand)
	if err1!=nil {
		Log.Error("failed to prepare stmtU: ", err1)
		return false
	}
	defer stmtU.Close()

	if _, err2 := stmtU.Exec(name, mail, mobile, authority); err2!=nil {
		Log.Error("failed to exec stmtU: ", err2)
		return false
	}

	stmtS, err3 := tx.Prepare(secretAddCommand)
	if err!=nil {
		Log.Error("failed to prepare stmtS: ", err3)
		return false
	}
	defer stmtS.Close()

	if _, err4 := stmtS.Exec(name, secretkey); err4!=nil {
		Log.Error("failed to exec stmtS: ", err4)
		return false
	}

	if err = tx.Commit(); err!=nil {
		return false
	}

	return true
}

func RemoveUserWithSecretkey(name string) bool {
	mu.Lock()
	defer mu.Unlock()

	if name == "" {
		return false
	}

	tx, err := conn.Begin()
	if err!=nil {
		Log.Error("failed to begin transaction")
		return false
	}
	defer tx.Rollback()

	stmtS, err1 := tx.Prepare(secretRemoveCommand)
	if err1!=nil {
		Log.Error("failed to prepare stmtS: ", err1)
		return false
	}
	defer stmtS.Close()

	if _, err2 := stmtS.Exec(name); err2!=nil {
		Log.Error("failed to exec stmtS: ", err2)
		return false
	}

	stmtU, err3 := tx.Prepare(userRemoveCommand)
	if err!=nil {
		Log.Error("failed to prepare stmtU: ", err3)
		return false
	}
	defer stmtU.Close()

	if _, err4 := stmtU.Exec(name); err4!=nil {
		Log.Error("failed to exec stmtU: ", err4)
		return false
	}

	if err = tx.Commit(); err!=nil {
		return false
	}

	return true
}

func InitShadow() error {
	command := "CREATE TABLE IF NOT EXISTS " + shadow_table + "(name TEXT PRIMARY KEY NOT NULL, secretkey TEXT NOT NULL)"

    return CreateTableIfNotExist(command)
}

func VerifySecretkey(name string, secretkey string) bool {
	command := "SELECT secretkey FROM " + shadow_table + " WHERE name = ?"
	rows, err := DoQuery(command, name)
	if err!=nil {
		Log.Error("failed to query: ", err)
		return false
	}
	defer rows.Close()

	var v string
	for rows.Next() {
		if err := rows.Scan(&v); err!=nil {
			Log.Error("scan rows fails: ", err)
			continue
		}
	}
	
	return v == secretkey
}

func ModifySecretkey(name string, secretkey string) error {
	command := "UPDATE " + user_table + "SET " + secretkey + " = ? WHERE name = ?"
	return PrepareAndExecuteCommand(command, secretkey, name)
}