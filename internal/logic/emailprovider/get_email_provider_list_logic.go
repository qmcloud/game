package emailprovider

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailProviderListLogic {
	return &GetEmailProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailProviderListLogic) GetEmailProviderList(in *game.EmailProviderListReq) (*game.EmailProviderListResp, error) {
	// todo: add your logic here and delete this line

	return &game.EmailProviderListResp{}, nil
}
