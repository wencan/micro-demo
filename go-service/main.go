package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	mrpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	proto "github.com/wencan/micro-demo/proto/health"
)

// HealthService 健康检查服务
type HealthService struct {
}

// Check 检查指定服务的健康状况
func (healthService *HealthService) Check(ctx context.Context, req *proto.HealthCheckRequest, resp *proto.HealthCheckResponse) error {
	switch req.Service {
	case "go-service":
		resp.Status = proto.HealthCheckResponse_SERVING
	default:
		resp.Status = proto.HealthCheckResponse_UNKNOWN
	}
	return nil
}

// Watch 监测服务的健康状况
func (healthService *HealthService) Watch(ctx context.Context, req *proto.HealthCheckRequest, resp proto.Health_WatchStream) error {
	return grpc.Errorf(codes.Unimplemented, "懒")
}

func main() {
	service := mrpc.NewService(micro.Name("micro-demo.go-service"))

	err := proto.RegisterHealthHandler(service.Server(), new(HealthService))
	if err != nil {
		log.Fatalln(err)
	}

	err = service.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
