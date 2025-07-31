package logic

import (
	"context"

	"user_rpc/internal/svc"
	"user_rpc/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type User_rpcLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUser_rpcLogic(ctx context.Context, svcCtx *svc.ServiceContext) *User_rpcLogic {
	return &User_rpcLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *User_rpcLogic) User_rpc(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
