package main

import (
	"context"
	pb "shippy/user-service/proto/user"
)

type handler struct {
	repo Repository
	//tokenService Authable
}

func (h *handler) Create(ctx context.Context, req *pb.User, resp *pb.Response) error {

	if err := h.repo.Create(req); err != nil {
		return nil
	}

	resp.User = req
	return nil
}

func (h *handler) Get(ctx context.Context, req *pb.User, resp *pb.Response) error {
	u, err := h.repo.Get(req.Id)

	if err != nil {
		return err
	}

	resp.User = u
	return nil
}

func (h *handler) GetAll(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	u, err := h.repo.GetAll()

	if err != nil {
		return err
	}

	resp.Users = u
	return nil
}

func (s *handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	_, err := s.repo.GetByEmailAndPassword(req)

	if err != nil {
		return err
	}

	resp.Token = "testing"
	return nil
}

func (s *handler) ValidateToken(ctx context.Context, req *pb.Token, resp *pb.Token) error {
	return nil
}
