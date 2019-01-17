package cli_c

import (
	"github.com/micro/go-micro/client"
	hello "shippy/greeter/service/proto"
	user "shippy/user-service/proto/user"
)

var (
	Cl hello.SayService
	U  user.UserServiceClient
)

func Init() {
	Cl = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)
	U = user.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)

}
