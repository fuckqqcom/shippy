package main

import (
	"log"

	"shippy/user-service/code"
	pb "shippy/user-service/proto/user"

	"github.com/micro/go-micro"
)

func main() {
	db, err := code.CreateConnection()
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.LogMode(true)
	db.SingularTable(true)

	db.AutoMigrate(&pb.User{})
	repo := &code.UserRepository{DB: db}

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	s.Init()

	pb.RegisterUserServiceHandler(s.Server(), &code.Handler{Repo: repo, TokenService: &code.TokenService{Repo: repo}})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}
}
