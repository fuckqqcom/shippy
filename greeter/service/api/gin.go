package main

import (
	"context"
	"fmt"
	"log"

	hello "shippy/greeter/service/proto"
	user "shippy/user-service/proto/user"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"
)

func main() {
	service := web.NewService(web.Name("go.micro.api.greeter"))
	service.Handle("/", New().r)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}

type Engine struct {
	srv *ServiceClient
	r   *gin.Engine
}

func New() (e *Engine) {
	e = &Engine{
		srv: NewServiceClient(),
		r:   gin.Default(),
	}
	e.r.GET("/greeter", Anything)
	e.r.GET("/greeter/:name", e.Hello)
	e.r.GET("/create", e.CreateUser)
	e.r.GET("/getAll", e.GetAllUser)
	e.r.GET("/auth", e.Auth)
	return
}

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

func (e *Engine) Hello(c *gin.Context) {
	log.Printf("received say.Hello api request")
	name := c.Param("name")
	resp, err := e.srv.cl.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func (e *Engine) CreateUser(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")
	pwd := c.Query("pwd")
	company := c.Query("company")

	resp, err := e.srv.u.Create(context.TODO(), &user.User{
		Name: name, Email: email, Password: pwd, Company: company})

	if err != nil {
		log.Fatalf("call Create error : %v", err)
		c.JSON(200, err)
	}

	c.JSON(200, resp)

}

func (e *Engine) GetAllUser(c *gin.Context) {
	log.Printf("received user.GetAllUser api request")
	fmt.Println("cl--->", e.srv.u)

	allResp, err := e.srv.u.GetAll(context.Background(), &user.Request{})
	if err != nil {
		log.Fatalf("call GetAll error:%v", err)
		c.JSON(200, err)
	}

	for _, u := range allResp.Users {
		log.Printf("%v\n", u)
	}
	c.JSON(200, allResp)
}

func (e *Engine) Auth(c *gin.Context) {
	email := c.Query("email")
	pwd := c.Query("pwd")
	fmt.Println("auth--->", email, pwd)
	authResp, err := e.srv.u.Auth(context.Background(), &user.User{
		Email:    email,
		Password: pwd,
	})

	if err != nil {
		log.Fatalf("auth failed : %v", err)
	}
	log.Println("token: ", authResp.Token)
	c.JSON(200, authResp.Token)
}
