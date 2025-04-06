package smsprovider

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSmsProviderLogic {
	return &UpdateSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSmsProviderLogic) UpdateSmsProvider(in *game.SmsProviderInfo) (*game.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &game.BaseResp{}, nil
}
