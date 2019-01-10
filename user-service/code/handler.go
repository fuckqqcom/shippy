package code

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	pb "shippy/user-service/proto/user"
)

type Handler struct {
	Repo         Repository
	tokenService Authable
}

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
	u, err := h.Repo.GetByEmailAndPassword(req)

	if err != nil {
		return err
	}

	//进行密码验证
	//
	//if err := bcrypt.CompareHashAndPassword([]byte(req.Password), []byte(u.Password)); err != nil {
	//	return err
	//}

	t, err := h.tokenService.Encode(u)

	if err != nil {
		return err
	}
	resp.Token = t
	return nil
}

func (s *Handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	return nil
}
