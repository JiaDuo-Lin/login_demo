package jwt

import (
	"context"
	pb "demo/proto"
)

type IJWTService interface {
	CreateToken(context.Context, *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error)
	CheckToken(context.Context, *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error)
}
