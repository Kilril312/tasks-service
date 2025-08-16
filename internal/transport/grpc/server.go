package grpc

import (
	"net"

	taskpb "github.com/Kilril312/project-protos/proto/task"
	userpb "github.com/Kilril312/project-protos/proto/user"
	"github.com/Kilril312/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.Service, uc userpb.UserServiceClient) error {
	netlist, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()

	handler := NewHandler(svc, uc)

	taskpb.RegisterTaskServiceServer(srv, handler)

	return srv.Serve(netlist)
}
