package main

import (
	"context"
	"github.com/micro/go-micro"
	"log"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	Port = ":50051"
)

//仓库接口
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

//存放多批货物的仓库,实现了IRepository接口
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//定义微服务
type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {

	//接收新的货物
	consignment, err := s.repo.Create(req)

	if err != nil {
		return err
	}

	resp = &pb.Response{Created: true, Consignment: consignment}
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	allConsignments := s.repo.GetAll()
	resp = &pb.Response{Consignments: allConsignments}
	return nil
}

func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	server.Init()
	repo := Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo: repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	//listener, err := net.Listen("tcp", Port)
	//
	//if err != nil {
	//	log.Fatalf("failed to listen:%v", err)
	//}
	//
	//log.Printf("listen on : %s \n", Port)
	//
	//server := grpc.NewServer()
	//repo := Repository{}
	//
	//pb.RegisterShippingServiceServer(server, &service{repo})
	//
	//if err := server.Serve(listener); err != nil {
	//	log.Fatalf("failed to save :%v", err)
	//}
}
