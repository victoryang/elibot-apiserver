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

type UserInfo struct {
	User
	UserId			int
	Status			int
	CreatedTime		string
	ModifiedTime	string
}

type Shadow struct {
	UserId 		int
	Password 	string
}

var user_table = "elt_user"
var shadow_table = "elt_shadow"

// USER TABLE COMMAND
var userInsertCommand = "INSERT INTO " + user_table + "(name, mail, mobile, authority) VALUES(?,?,?,?)"
var userDeleteCommand = "DELETE FROM " + user_table + " WHERE name = ?"
var userQueryCommand = "SELECT name, mail, mobile, authority FROM " + user_table
var userUpdateCommand = "UPDATE " + user_table + " SET mail = ?, mobile = ?, authority = ? , mtime = datetime('now','localtime') WHERE name = ?"

// PASSWORD TABLE COMMAND
var shadowInsertCommand = "INSERT INTO " + shadow_table + " VALUES((SELECT userid FROM " + user_table + " WHERE name = ?), ?)"
var shadowDeleteCommand = "DELETE FROM " + shadow_table + " WHERE userid = (SELECT userid FROM " + user_table + " WHERE name = ?)"
var shadowQueryCommand = "SELECT password FROM " + shadow_table + " WHERE userid = (SELECT userid FROM " + user_table + " WHERE name = ?)"
var shadowUpdateCommand = "UPDATE " + shadow_table + " SET password = ? WHERE userid = (SELECT userid FROM " + user_table + " WHERE name = ?)"

func InitUserTable() error {
	table := user_table +
		"(userid INTEGER PRIMARY KEY AUTOINCREMENT, " +
	    "name TEXT UNIQUE NOT NULL, mail TEXT, mobile TEXT, authority INTEGER NOT NULL, status INTEGER," +
	    "ctime TimeStamp NOT NULL DEFAULT (datetime('now','localtime'))," +
	    "mtime TimeStamp NOT NULL DEFAULT (datetime('now','localtime')))"

    return InitializeTable(table)
}

func GetUsersList() []User {
	rows, err := DoQuery(userQueryCommand)
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
	return PrepareAndExecuteCommand(userUpdateCommand, u.Mail, u.Mobile, u.Authority, u.Name)
}

func AddUserWithPassword(u User, password string) bool {
	mu.Lock()
	defer mu.Unlock()

	if u.Name == "" || password == "" {
		return false
	}

	tx, err := conn.Begin()
	if err!=nil {
		Log.Error("failed to begin transaction")
		return false
	}
	defer tx.Rollback()

	stmtU, err1 := tx.Prepare(userInsertCommand)
	if err1!=nil {
		Log.Error("failed to prepare stmtU: ", err1)
		return false
	}
	defer stmtU.Close()

	if _, err2 := stmtU.Exec(u.Name, u.Mail, u.Mobile, u.Authority); err2!=nil {
		Log.Error("failed to exec stmtU: ", err2)
		return false
	}

	stmtS, err3 := tx.Prepare(shadowInsertCommand)
	if err!=nil {
		Log.Error("failed to prepare stmtS: ", err3)
		return false
	}
	defer stmtS.Close()

	if _, err4 := stmtS.Exec(u.Name, password); err4!=nil {
		Log.Error("failed to exec stmtS: ", err4)
		return false
	}

	if err = tx.Commit(); err!=nil {
		return false
	}

	return true
}

func RemoveUserAndPassword(name string) bool {
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

	stmtS, err1 := tx.Prepare(shadowDeleteCommand)
	if err1!=nil {
		Log.Error("failed to prepare stmtS: ", err1)
		return false
	}
	defer stmtS.Close()

	if _, err2 := stmtS.Exec(name); err2!=nil {
		Log.Error("failed to exec stmtS: ", err2)
		return false
	}

	stmtU, err3 := tx.Prepare(userDeleteCommand)
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
	table := shadow_table + "(userid INTEGER PRIMARY KEY NOT NULL, password TEXT NOT NULL)"

    return InitializeTable(table)
}

func VerifyPassword(name string, password string) bool {
	rows, err := DoQuery(shadowQueryCommand, name)
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
	
	return v == password
}

func ModifyPassword(name string, password string) error {
	return PrepareAndExecuteCommand(shadowUpdateCommand, password, name)
}