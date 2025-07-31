package hello

import (
	"context"

	"Music_api/internal/svc"
	"Music_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HellLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHellLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HellLogic {
	return &HellLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HellLogic) Hell(req *types.HelloReq) (resp *types.HelloResp, err error) {
	// todo: add your logic here and delete this line

	return
}
