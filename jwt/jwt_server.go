package jwt

import (
	"context"
	pb "demo/proto"
	"log"
)

type JWTServer struct {
	jwtService IJWTService
}

func NewJWTServer() *JWTServer {
	return &JWTServer{jwtService: NewJWTService()}
}

func (j *JWTServer) CreateToken(ctx context.Context,
	req *pb.CreateTokenRequest) (resp *pb.CreateTokenResponse, err error) {

	if resp, err = j.jwtService.CreateToken(ctx, req); err != nil {
		log.Fatal("err in create token")
	}
	return
}

func (j *JWTServer) CheckToken(ctx context.Context,
	req *pb.CheckTokenRequest) (resp *pb.CheckTokenResponse, err error) {

	if resp, err = j.jwtService.CheckToken(ctx, req); err != nil {
		log.Fatal("err in Check token")
	}
	return
}
