package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"log"
	pb "shippy/consignment-service/proto/consignment"
	vesselProto "shippy/vessel-service/proto/vessel"
)

const (
	Port = ":50051"
)

//仓库接口
type Repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

//存放多批货物的仓库,实现了IRepository接口
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//定义微服务
type service struct {
	repo         Repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})

	fmt.Printf("Found vessel:%s \n", vesselResponse.Vessel.Name)

	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	consignment, err := s.repo.Create(req)

	if err != nil {
		return err
	}

	//接收新的货物
	consignment, err := s.repo.Create(req)

	if err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments := s.repo.GetAll()
	resp.Consignments = consignments
	return nil
}

func main() {

	repo := Repository{}
	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	srv.Init()
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

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
