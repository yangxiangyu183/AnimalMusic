package logic

import (
	"context"

	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_apiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_apiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_apiLogic {
	return &User_apiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_apiLogic) User_api(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
