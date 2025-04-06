package smsprovider

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSmsProviderLogic {
	return &DeleteSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSmsProviderLogic) DeleteSmsProvider(in *game.IDsReq) (*game.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &game.BaseResp{}, nil
}
