package smslog

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsLogByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsLogByIdLogic {
	return &GetSmsLogByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsLogByIdLogic) GetSmsLogById(in *game.UUIDReq) (*game.SmsLogInfo, error) {
	// todo: add your logic here and delete this line

	return &game.SmsLogInfo{}, nil
}
