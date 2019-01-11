package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/user-service/proto/user"
)

const topic = "user.created"

type Subscriber struct {
}

func main() {
	s := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	s.Init()

	micro.RegisterSubscriber(topic, s.Server(), new(Subscriber))
	if err := s.Run(); err != nil {
		fmt.Printf("srv run error: %v\n", err)
	}
}

//func (sub *Subscriber) Process666(ctx context.Context, user *pb.User) error {
//	log.Println("[Picked up a new message]")
//	log.Println("[Sending email to]:", user)
//	return nil
//}

func (sub *Subscriber) senEmail(user *pb.User) error {

	log.Printf("senEmail ---> [SENDING A EMAIL TO %s...]", user.Name)
	return nil
}
