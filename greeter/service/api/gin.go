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

type Say struct {
}

var cl hello.SayService
var u user.UserServiceClient

func Anything(c *gin.Context) {
	log.Print("received say.Anything api request")
	c.JSON(200, map[string]string{"msg": "hahahahahha"})

}

func Hello(c *gin.Context) {
	log.Printf("received say.Hello api request")
	name := c.Param("name")
	resp, err := cl.Hello(context.TODO(), &hello.Request{
		Name: name,
	})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, resp)
}

func CreateUser(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")
	pwd := c.Query("pwd")
	company := c.Query("company")

	resp, err := u.Create(context.TODO(), &user.User{
		Name: name, Email: email, Password: pwd, Company: company})

	if err != nil {
		log.Fatalf("call Create error : %v", err)
		c.JSON(200, err)
	}

	c.JSON(200, resp)

}

func GetAllUser(c *gin.Context) {
	log.Printf("received user.GetAllUser api request")

	fmt.Println(u)
	allResp, err := u.GetAll(context.Background(), &user.Request{})

	if err != nil {
		log.Fatalf("call GetAll error:%v", err)
		c.JSON(200, err)
	}

	for _, u := range allResp.Users {
		log.Printf("%v\n", u)
	}
	c.JSON(200, err)
}

func Auth(c *gin.Context) {
	email := c.Query("email")
	pwd := c.Query("pwd")
	fmt.Println("auth--->", email, pwd)
	authResp, err := u.Auth(context.Background(), &user.User{
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
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
		web.Name("go.micro.srv.user"),
	)
	service.Init()
	// init service obj
	cl = hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)
	u = user.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)

	r := gin.Default()
	r.GET("/greeter", Anything)
	r.GET("/greeter/:name", Hello)
	r.GET("/create", CreateUser)
	r.GET("/getAll", GetAllUser)
	r.GET("/auth", Auth)

	service.Handle("/", r)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
