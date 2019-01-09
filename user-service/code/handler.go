package code

import (
	"context"
	pb "shippy/user-service/proto/user"
)

type Handler struct {
	Repo Repository
	//tokenService Authable
}

func (h *Handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {

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

func (s *Handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	_, err := s.Repo.GetByEmailAndPassword(req)

	if err != nil {
		return err
	}

	resp.Token = "testing"
	return nil
}

func (s *Handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	return nil
}
