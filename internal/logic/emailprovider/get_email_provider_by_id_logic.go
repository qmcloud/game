package emailprovider

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailProviderByIdLogic {
	return &GetEmailProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailProviderByIdLogic) GetEmailProviderById(in *game.IDReq) (*game.EmailProviderInfo, error) {
	// todo: add your logic here and delete this line

	return &game.EmailProviderInfo{}, nil
}
