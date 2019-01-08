package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/user-service/proto/user"
)

func main() {
	db, err := CreateConnection()

	fmt.Printf("%+v\n", db)
	fmt.Printf("err: %v\n", err)

	defer db.Close()

	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	repo := &UserRepository{db}

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)

	s.Init()

	pb.RegisterUserServiceHandler(s.Server(), &handler{repo: repo})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}
}
