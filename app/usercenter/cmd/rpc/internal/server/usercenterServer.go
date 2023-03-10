// Code generated by goctl. DO NOT EDIT!
// Source: usercenter.proto

package server

import (
	"context"

	"gozerok8s/app/usercenter/cmd/rpc/internal/logic"
	"gozerok8s/app/usercenter/cmd/rpc/internal/svc"
	"gozerok8s/app/usercenter/cmd/rpc/pb"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}
