package user

import (
	pb "demo/proto"
	"log"
)
import "context"

type UserServer struct {
	userService IUserService
}

func NewUserSever() *UserServer {
	return &UserServer{
		userService:NewUserService(),
	}
}

func (u *UserServer) IsRegistered(ctx context.Context,
	req *pb.IsRegisteredRequest) (*pb.IsRegisteredResponse, error) {

	resp, err :=  u.userService.IsRegistered(ctx, req)
	if err != nil {
		log.Printf("UserServer IsRegistered has an err:[%v]", err)
	}
	return resp, err
}

func (u *UserServer) Add(ctx context.Context,
	req *pb.AddRequest) (*pb.AddResponse, error)  {

	resp, err := u.userService.Add(ctx, req)
	if err != nil {
		log.Printf("UserServer Add has an err:[%v]", err)
	}
	return resp, err
}

func (u *UserServer)  Load(ctx context.Context,
	req *pb.LoadRequest) (*pb.LoadResponse, error) {

	resp, err := u.userService.Load(ctx, req)
	if err != nil {
		log.Printf("UserServer Load has an err:[%v]", err)
	}
	return resp, err
}