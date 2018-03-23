package mongo

import (
  "2C_login_ms/pkg"
  "gopkg.in/mgo.v2/bson"
  "gopkg.in/mgo.v2"
  "golang.org/x/crypto/bcrypt"
)

type userModel struct {
  Id           bson.ObjectId `bson:"_id,omitempty"`
  UserId       uint64
  Username     string
  Password     string
  Email        string
  Name         string
  LastName     string
  Cellphone    int
}

func userModelIndex() mgo.Index {
  return mgo.Index{
    Key:        []string{"username"},
    Unique:     true,
    DropDups:   true,
    Background: true,
    Sparse:     true,
  }
}

func newUserModel(u *root.User) *userModel {
  return &userModel{
    UserId:   u.UserId,
    Username: u.Username,
    Password: u.Password,
    Email:    u.Email,
    Name:     u.Name,
    LastName: u.LastName,
    Cellphone:u.Cellphone}
}

func(u *userModel) toRootUser() *root.User {
  return &root.User{
    Id:       u.Id.Hex(),
    UserId:   u.UserId,
    Username: u.Username,
    Password: u.Password,
    Email:    u.Email,
    Name:     u.Name,
    LastName: u.LastName,
    Cellphone:u.Cellphone}
}

func(u *userModel) comparePassword(password string) error {
  incoming := []byte(password)
  existing := []byte(u.Password)
  err := bcrypt.CompareHashAndPassword(existing, incoming)
  return err
}
