package code

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	"log"

	pb "shippy/user-service/proto/user"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	Repo         Repository
	TokenService Authable
	Publisher    micro.Publisher
}

const Topic = "user.created"

func (h *Handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	req.Password = string(pwd)

	if err := h.Repo.Create(req); err != nil {
		return nil
	}

	resp.User = req

	//发布带有用户所有信息的消息
	if err := h.Publisher.Publish(ctx, req); err != nil {
		return err
	}
	return nil
}

func (h *Handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	u, err := h.Repo.Get(req.Uid)

	if err != nil {
		return err
	}

	resp.User = u
	return nil
}

func (h *Handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	u, err := h.Repo.GetAll()

	if err != nil {
		return err
	}

	resp.Users = u
	return nil
}

func (h *Handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	log.Println("Logging in with:", req.Email, req.Password)

	u, err := h.Repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	hashPwd := []byte(u.Password)
	pwd := []byte(req.Password)

	fmt.Println(1111, string(hashPwd))
	fmt.Println(2222, string(pwd))
	//进行密码验证
	if err := bcrypt.CompareHashAndPassword(hashPwd, pwd); err != nil { //注意：第二个参数一定是原始密码字符串，而非hash
		log.Printf("bcrypt.CompareHashAndPassword error:%v", err)
		return err
	}

	t, err := h.TokenService.Encode(u)
	if err != nil {
		return err
	}
	resp.Token = t
	return nil
}

func (srv *Handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	// spew.Dump(req.Token)
	// Decode token
	claims, err := srv.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	log.Println("get token", claims)

	if claims.User.Uid == "" {
		return errors.New("invalid user")
	}

	resp.Valid = true
	return nil
}
