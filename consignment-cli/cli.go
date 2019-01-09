package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	Address         = "localhost:50051"
	DefaultInfoFile = "consignment-cli/consignment.json"
)

func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment

	err = json.Unmarshal(data, &consignment)

	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}

	return consignment, nil
}

func main() {
	//conn, err := grpc.Dial(Address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("connect error:%v", err)
	//}

	//defer conn.Close()
	cmd.Init()
	c := pb.NewShippingServiceClient("go.micro.srv.consignment", client.DefaultClient)

	infoFile := DefaultInfoFile

	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	resp, err := c.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}
	fmt.Println(resp)

	//
	//log.Printf("created : %t", resp.Created)
	//
	//resp, err = c.GetConsignments(context.Background(), &pb.GetRequest{})
	//
	//if err != nil {
	//	log.Fatalf("failed to list consignments: %v", err)
	//}
	//for _, c := range resp.Consignments {
	//	log.Printf("%+v", c)
	//}
}
