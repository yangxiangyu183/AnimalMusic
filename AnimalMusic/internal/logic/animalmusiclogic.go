package logic

import (
	"context"

	"AnimalMusic/internal/svc"
	"AnimalMusic/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnimalMusicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAnimalMusicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnimalMusicLogic {
	return &AnimalMusicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AnimalMusicLogic) AnimalMusic(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
