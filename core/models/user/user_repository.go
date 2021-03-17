package user

import (
	"github.com/scylladb/gocqlx/v2"
)

type UserRepoInterface interface {
	InsertUser(*User) error
	RetrieveUser(string) (*User, error)
	UpdateUser(string, *User) error
	DeleteUser(string) error
}

type UserRepo struct {
	*gocqlx.Session
}

//InsertUser function will try to insert to db
func (ur *UserRepo) InsertUser(user *User) error {
	return ur.Session.Query(userTable.Insert()).BindStruct(*user).ExecRelease()
}

func (ur *UserRepo) RetrieveUser(user *User) error {
	return ur.Session.Query(userTable.Get()).BindStruct(*user).GetRelease(user)
}

func (ur *UserRepo) UpdateUser(username string, updatedUser *User, updatedColumns ...string) (bool, error) {
	return ur.Session.Query(userTable.Update(updatedColumns...)).BindStruct(*updatedUser).ExecCASRelease()
}

func (ur *UserRepo) DeleteUser(username string) error {
	user := User{
		Username: username,
	}
	return ur.Session.Query(userTable.Delete()).BindStruct(user).ExecRelease()
}
