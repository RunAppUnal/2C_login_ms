package root

type User struct {
  Id           string  `json:"id"`
  UserId       uint64  `json:"userid"`
  Username     string  `json:"username"`
  Password     string  `json:"password"`
  Email        string  `json:"email"`
  Name         string  `json:"name"`
  LastName     string  `json:"lastname"`
  Cellphone    int     `json:"cellphone`
}

type UserService interface {
  CreateUser(u *User) error
  GetByUsername(username string) (*User,error)
  DeleteByUsername(username string) error
  Login(c Credentials) (error, User)
}
