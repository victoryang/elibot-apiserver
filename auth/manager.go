package auth

import (
	"sync"

	"elibot-apiserver/db"
	Log "elibot-apiserver/log"
)

type Manager interface {
	AddUser(u db.User) bool
	RemoveUser(name string) bool
	EditUser(username string) bool
	GetUsersList() []db.User
}

type UserManager struct {
	Mutex		sync.Mutex
	Users 		[]db.User
}

var mgr *UserManager

func GetUserManager() *UserManager {
	return mgr
}

func InitManager() error {
	if err := db.InitUserTable(); err!=nil {
    	Log.Error("Could not initialize the user table")
    	return err
    }

    if err := db.InitShadow(); err!=nil {
    	Log.Error("Could not initialize the shadow table")
    	return err
    }

    mgr = new(UserManager)
    mgr.Users = make([]db.User, 0)
	mgr.UpdateUserList()

	return nil
}

func (m *UserManager) GetUsersList() []db.User {
	return m.Users
}

func (m *UserManager) UpdateUserList() {
	m.Users = db.GetUsersList()
}

func (m *UserManager) GetUser(name string, user *db.User) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.UpdateUserList()

	for _, u := range m.Users {
		if u.Name == name {
			user = &u
			return true
		}
	}

	return false
}

func (m *UserManager) AddUser(u db.User, secretkey string) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	if !db.AddUserWithSecretkey(u.Name, u.Mail, u.Mobile, u.Authority, secretkey) {
		Log.Error("Could not add item into user table")
		return false
	}
	
	m.UpdateUserList()
	return true
}

func (m *UserManager) RemoveUser(name string) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	if !db.RemoveUserWithSecretkey(name) {
		Log.Error("Could not remove item from user table")
		return false
	}
	
	m.UpdateUserList()
	return true
}

func (m *UserManager) ModifyUser(u db.User) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	if u.Name == "" {
		return false
	}

	var old db.User
	if m.GetUser(u.Name, &old) {
		if u.Mail == "" {
			u.Mail = old.Mail
		}

		if u.Mobile == "" {
			u.Mobile = old.Mobile
		}
	}

	if err := db.ModifyUser(u); err!=nil {
		Log.Error("Could not update user table: ", err)
		return false
	}
	
	m.UpdateUserList()
	return true
}

func (m *UserManager) ChangePassword(name string, password string) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	if name == "" || password == "" {
		return false
	}

	if err := db.ModifySecretkey(name, password); err!=nil {
		Log.Error("Could change password: ", err)
		return false
	}
	
	m.UpdateUserList()
	return true
}