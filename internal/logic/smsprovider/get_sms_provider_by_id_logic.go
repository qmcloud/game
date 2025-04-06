package smsprovider

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderByIdLogic {
	return &GetSmsProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsProviderByIdLogic) GetSmsProviderById(in *game.IDReq) (*game.SmsProviderInfo, error) {
	// todo: add your logic here and delete this line

	return &game.SmsProviderInfo{}, nil
}
