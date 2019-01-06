package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/micro/go-micro"
	pb "shippy/vessel-service/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

//接口实现

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, v := range repo.vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("no vessel can't be use")
}

//定义服务

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	v, err := s.repo.FindAvailable(spec)

	if err != nil {
		return err
	}

	resp.Vessel = v
	return nil
}

func main() {
	vessels := []*pb.Vessel{{Id: "vessel0001", Capacity: 500, MaxWeight: 200000, Name: "Boaty McBoatface"}}
	repo := &VesselRepository{vessels: vessels}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo: repo})

	if err := srv.Run(); err != nil {
		fmt.Println("err--->", err)
	}
}
