package user

import (
	pb "demo/proto"
	"log"
)
import (
	"context"
)

type UserService struct {}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) IsRegistered(ctx context.Context, 
	req *pb.IsRegisteredRequest) (*pb.IsRegisteredResponse, error)  {
	
	user := NewUser(req.Id, "", nil)
	ok := user.IsExist()
	return &pb.IsRegisteredResponse{Ok:ok}, nil
}

func (u *UserService) Add(ctx context.Context, 
	req *pb.AddRequest) (*pb.AddResponse, error)  {

	user := NewUser(req.Id, req.Name, req.Tags)
	user.Add()
	return &pb.AddResponse{Ok:true}, nil
}

func (u *UserService)  Load(ctx context.Context,
	req *pb.LoadRequest) (resp *pb.LoadResponse, err error) {

	user := NewUser(req.Id, "", nil)
	err = user.Load()
	if err != nil {
		log.Println(err)
		return
	}

	resp =  &pb.LoadResponse{
		Id:user.ID,
		Name:user.Name,
		Tags:user.Tags,
	}
	return
}