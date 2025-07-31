package logic

import (
	"context"

	"Music_service/internal/svc"
	"Music_service/music_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSmsLogic) SendSms(in *music_service.SendSmsRequest) (*music_service.SendSmsResponse, error) {
	// todo: add your logic here and delete this line

	return &music_service.SendSmsResponse{}, nil
}
