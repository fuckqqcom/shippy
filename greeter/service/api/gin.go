package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
	"log"
	hello "shippy/greeter/service/proto"
	user "shippy/user-service/proto/user"
)

type ServiceClient struct {
	cl hello.SayService
	u  user.UserServiceClient
}

func NewServiceClient() *ServiceClient {
	return &ServiceClient{
		cl: hello.NewSayService("go.micro.srv.greeter", client.DefaultClient),
		u:  user.NewUserServiceClient("go.micro.srv.user", client.DefaultClient),
	}
}

func Anything(c *gin.Context) {
	log.Print("received say.Anything api request")
	c.JSON(200, map[string]string{"msg": "hahahahahha"})

}

func (s ServiceClient) Hello(c *gin.Context) {
	log.Printf("received say.Hello api request")
	name := c.Param("name")

	resp, err := s.cl.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func (s ServiceClient) CreateUser(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")
	pwd := c.Query("pwd")
	company := c.Query("company")

	resp, err := s.u.Create(context.TODO(), &user.User{
		Name: name, Email: email, Password: pwd, Company: company})

	if err != nil {
		log.Fatalf("call Create error : %v", err)
		c.JSON(200, err)
	}

	c.JSON(200, resp)

}

func (s ServiceClient) GetAllUser(c *gin.Context) {
	log.Printf("received user.GetAllUser api request")
	fmt.Println("cl--->", s.u)

	allResp, err := s.u.GetAll(context.Background(), &user.Request{})

	if err != nil {
		log.Fatalf("call GetAll error:%v", err)
		c.JSON(200, err)
	}

	for _, u := range allResp.Users {
		log.Printf("%v\n", u)
	}
	c.JSON(200, allResp)
}

func (s ServiceClient) Auth(c *gin.Context) {
	email := c.Query("email")
	pwd := c.Query("pwd")
	fmt.Println("auth--->", email, pwd)
	fmt.Println("cl--->", s.u)
	authResp, err := s.u.Auth(context.Background(), &user.User{
		Email:    email,
		Password: pwd,
	})

	if err != nil {
		log.Fatalf("auth failed : %v", err)
	}
	log.Println("token: ", authResp.Token)
	c.JSON(200, authResp.Token)
}

func main() {
	service := web.NewService(web.Name("go.micro.api.greeter"))

	service.Init()
	serviceClient := NewServiceClient()
	fmt.Println("serviceClient--->", serviceClient.cl, serviceClient.u)
	s := ServiceClient{serviceClient.cl, serviceClient.u}

	r := gin.Default()
	r.GET("/greeter", Anything)
	r.GET("/greeter/:name", s.Hello)
	r.GET("/create", s.CreateUser)
	r.GET("/getAll", s.GetAllUser)
	r.GET("/auth", s.Auth)

	service.Handle("/", r)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
