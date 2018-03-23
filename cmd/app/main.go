package main

import (
	"2C_login_ms/pkg/crypto"
	"2C_login_ms/pkg/mongo"
	"2C_login_ms/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("192.168.99.101:27017")
	if err != nil {
		log.Println(err)
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	h := crypto.Hash{}
	u := mongo.NewUserService(ms.Copy(), "2C_login_db", "user", &h)
	s := server.NewServer(u)

	s.Start()
}
