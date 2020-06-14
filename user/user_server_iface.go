package user

import pb "demo/proto"
import (
	"context"
)

type IUserService interface {
	IsRegistered(context.Context, *pb.IsRegisteredRequest) (*pb.IsRegisteredResponse, error)
	Add(context.Context, *pb.AddRequest) (*pb.AddResponse, error)
	Load(context.Context, *pb.LoadRequest) (*pb.LoadResponse, error)
}
