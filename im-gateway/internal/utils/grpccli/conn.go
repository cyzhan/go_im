package grpccli

import (
	"os"

	"google.golang.org/grpc"
)

var (
	imLogicConn *grpc.ClientConn
)

func NewGrpc() {
	apiServer := os.Getenv("LOGIC_SERVER")
	conn, err := NewConn(apiServer)
	if err != nil {
		panic(err)
	}

	imLogicConn = conn
}

func GetApiConn() *grpc.ClientConn {
	return imLogicConn
}
