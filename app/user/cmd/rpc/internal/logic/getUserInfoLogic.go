package logic

import (
	"context"
	"gozerok8s/app/user/cmd/rpc/internal/svc"
	"gozerok8s/app/user/cmd/rpc/pb"
	"gozerok8s/app/user/cmd/rpc/usercenter"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	userId := in.GetId()
	var respUser usercenter.User
	respUser.Id = userId
	respUser.Mobile = "19999999999"
	respUser.Avatar = "男"
	hostName, _ := os.Hostname()
	respUser.Info = hostName
	respUser.Nickname = "xx"
	return &pb.GetUserInfoResp{
		User: &respUser,
	}, nil
}
