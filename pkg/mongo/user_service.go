package mongo

import (
	"2C_login_ms/pkg"
	//"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/night-codes/mgo-ai"
)

type UserService struct {
	collection *mgo.Collection
	hash       root.Hash
}

func NewUserService(session *Session, dbName string, collectionName string, hash root.Hash) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection, hash}
}

func (p *UserService) CreateUser(u *root.User) error {
	user := newUserModel(u)
	hashedPassword, err := p.hash.Generate(user.Password)

	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// connect AutoIncrement to collection "users"
	session, err := mgo.Dial("192.168.99.101:27017")
  ai.Connect(session.DB("2C_login_db").C("user"))
	user.UserId = ai.Next("user")

	return p.collection.Insert(&user)
}

func (p *UserService) DeleteByUsername(username string)  error{
	err := p.collection.Remove(bson.M{"username": username})
	return  err
}

func (p *UserService) GetByUsername(username string) (*root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}

func (p *UserService) Login(c root.Credentials) (error, root.User) {
  model := userModel{}
  err := p.collection.Find(bson.M{"username": c.Username}).One(&model)

  err = model.comparePassword(c.Password)
  if(err != nil) {
    return err, root.User{}
  }

  return err, root.User{
    Id: model.Id.Hex(),
		UserId: model.UserId,
    Username: model.Username,
    Password: "-",
	 	Email: model.Email,
		Name: model.Name,
	  LastName: model.LastName,
		Cellphone: model.Cellphone}
}
