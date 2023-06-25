package internal

import (
	"log"
	"net"
	"os"

	"imlogic/internal/rpc/message"
	"imlogic/internal/rpc/user"
	"imlogic/internal/rpc/vendors"
	"imlogic/internal/service"

	_ "github.com/joho/godotenv/autoload"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type RpcServer struct {
	Host string
	Port string
}

func NewRpcServer() *RpcServer {
	port := os.Getenv("RPC_PORT")
	host := os.Getenv("RPC_HOST")
	return &RpcServer{
		Host: host,
		Port: port,
	}
}

func (s *RpcServer) Run() {
	log.Println("Server run...")

	address := s.Host + ":" + s.Port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("listening err:" + err.Error())
	}
	log.Println("0", "grpc listen ok ", address)
	defer listener.Close()

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcRecovery.UnaryServerInterceptor(),
			grpcPrometheus.UnaryServerInterceptor,
		)),
	)
	s.registerSrv(srv)
	defer srv.GracefulStop()
	if err := srv.Serve(listener); err != nil {
		panic("serve err:" + err.Error())
	}
}

func (s *RpcServer) registerSrv(srv *grpc.Server) {
	vendors.RegisterVendorsServiceServer(srv, service.NewVendorsService())
	user.RegisterUserServiceServer(srv, service.NewUserService())
	message.RegisterMessageServiceServer(srv, service.NewMessageService())
}
