build:
#一定要注意makefile中的缩进,否则make build可能报错 Nothing to be done for build
#proto命令前边是一个tab,不是四个或者八个空格
#告知 Go 编译器生成二进制文件的目标环境：amd64 CPU 的 Linux 系统
#根据当前目录下的 Dockerfile 生成名为 consignment-service 的镜像
#在 Docker alpine 容器的 50001 端口上运行 consignment-service 服务
#可添加 -d 参数将微服务放到后台运行
	protoc -I.  --go_out=plugins=grpc:$(GOPATH)/src/shippy/consignment-service proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .
run:
	docker run -p 50051:50051 consignment-service
