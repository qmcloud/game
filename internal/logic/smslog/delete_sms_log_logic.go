package smslog

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSmsLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSmsLogLogic {
	return &DeleteSmsLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSmsLogLogic) DeleteSmsLog(in *game.UUIDsReq) (*game.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &game.BaseResp{}, nil
}
