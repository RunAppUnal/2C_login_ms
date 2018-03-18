package main

import (
	"go_rest_api/pkg/crypto"
	"go_rest_api/pkg/mongo"
	"go_rest_api/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	h := crypto.Hash{}
	u := mongo.NewUserService(ms.Copy(), "2C_login_db", "user", &h)
	s := server.NewServer(u)

	s.Start()
}
