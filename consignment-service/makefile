build:
    #一定要注意makefile中的缩进,否则make build可能报错 Nothing to be done for build
    #proto命令前边是一个tab,不是四个或者八个空格
    proto -I.  --go_out=plugins=grpc:$(GOPATH)/src/shippy/consignment-service proto/consignment/consignment.proto