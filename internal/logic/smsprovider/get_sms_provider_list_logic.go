package smsprovider

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderListLogic {
	return &GetSmsProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsProviderListLogic) GetSmsProviderList(in *game.SmsProviderListReq) (*game.SmsProviderListResp, error) {
	// todo: add your logic here and delete this line

	return &game.SmsProviderListResp{}, nil
}
