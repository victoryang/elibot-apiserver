package auth

import (
	"sync"
	"encoding/base64"
	"crypto/md5"

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

func (m *UserManager) GetUserAuthority(name string, authority *int) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for _, u := range m.Users {
		if u.Name == name {
			*authority = u.Authority
			return true
		}
	}

	return false
}

func (m *UserManager) GetUser(name string, user *db.User) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	for _, u := range m.Users {
		if u.Name == name {
			*user = u
			return true
		}
	}

	return false
}

func (m *UserManager) AddUser(u db.User, pwd string) bool {
	var epwd string
	if err := encodePassword(pwd, &epwd); err!=nil {
		return false
	}

	if !db.AddUserWithPassword(u, epwd) {
		Log.Error("Could not add item into user table")
		return false
	}

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.UpdateUserList()
	return true
}

func (m *UserManager) RemoveUser(name string) bool {
	if !db.RemoveUserAndPassword(name) {
		Log.Error("Could not remove item from user table")
		return false
	}

	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.UpdateUserList()
	return true
}

func (m *UserManager) ModifyUser(u db.User) bool {
	var old db.User
	if m.GetUser(u.Name, &old) {
		if u.Mail == "" {
			u.Mail = old.Mail
		}

		if u.Mobile == "" {
			u.Mobile = old.Mobile
		}
	} else {
		Log.Error("Could not modify user: not found")
        return false
	}

	if err := db.ModifyUser(u); err!=nil {
		Log.Error("Could not update user: ", err)
		return false
	}

	m.Mutex.Lock()
    defer m.Mutex.Unlock()

	m.UpdateUserList()
	return true
}

func encodePassword(password string, epassword *string) error {
	pw, err := base64.URLEncoding.DecodeString(password);
	if err == nil {
		sli := md5.Sum(pw)
		*epassword = base64.StdEncoding.EncodeToString(sli[:])
	}

	return err
}

func (m *UserManager) VerifyPassword(name string, pwd string) bool {
	var epwd string
	if err := encodePassword(pwd, &epwd); err!=nil {
		return false
	}

	return db.VerifyPassword(name, epwd)
}

func (m *UserManager) ChangePassword(name string, pwd string) bool {
	var epwd string
	if err := encodePassword(pwd, &epwd); err!=nil {
		return false
	}

	if err := db.ModifyPassword(name, epwd); err!=nil {
		Log.Error("Could change password: ", err)
		return false
	}

	return true
}