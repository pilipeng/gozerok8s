package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gozerok8s/app/usercenter/cmd/rpc/internal/svc"
	"gozerok8s/app/usercenter/cmd/rpc/pb"
	"gozerok8s/app/usercenter/cmd/rpc/usercenter"
	"os"
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

	//t := time.Now()
	//tm, _ := fmt.Printf("当0q445666666dd---5qq3345前的时间是: %d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	currentDir, _ := os.Getwd()

	var respUser usercenter.User
	respUser.Id = userId
	respUser.Mobile = "199999555s01"
	respUser.Avatar = "男"
	hostName, _ := os.Hostname()
	respUser.Info = hostName + " dir:" + currentDir

	respUser.Nickname = "xx"
	return &pb.GetUserInfoResp{
		User: &respUser,
	}, nil
}
