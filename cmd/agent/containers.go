package main

import (
	"context"
	"fmt"

	pb "github.com/XiovV/dokkup/grpc"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) DeployContainer(ctx context.Context, in *pb.DeployContainerRequest) (*empty.Empty, error) {
  fmt.Println(in) 
  
  return new(empty.Empty), nil
}