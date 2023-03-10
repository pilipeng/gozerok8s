// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "gozerok8s/app/usercenter/cmd/api/internal/handler/user"
	"gozerok8s/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: user.DetailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)
}
