package logic

import (
	"context"

	"Music_service/internal/svc"
	"Music_service/music_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *music_service.LoginRequest) (*music_service.LoginResponse, error) {
	// todo: add your logic here and delete this line

	return &music_service.LoginResponse{}, nil
}
