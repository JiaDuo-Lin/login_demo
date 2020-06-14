package jwt

import (
	pb "demo/proto"
	"log"
)
import "context"

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (j *JWTService) CreateToken(ctx context.Context,
	req *pb.CreateTokenRequest) (resp *pb.CreateTokenResponse, err error) {

	token, err := CreateToken(req.Name, req.Id)
	if err != nil {
		log.Println(err)
		return
	}
	return &pb.CreateTokenResponse{Token: token.Token}, nil
}

func (j *JWTService) CheckToken(ctx context.Context,
	req *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {

	ok := CheckToken(req.Token)

	return &pb.CheckTokenResponse{Result: ok}, nil
}
