package grpccli

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

func NewConn(target string) (*grpc.ClientConn, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(ctx, target,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithInsecure())

	return conn, err
}
