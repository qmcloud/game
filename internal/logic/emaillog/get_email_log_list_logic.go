package emaillog

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailLogListLogic {
	return &GetEmailLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailLogListLogic) GetEmailLogList(in *game.EmailLogListReq) (*game.EmailLogListResp, error) {
	// todo: add your logic here and delete this line

	return &game.EmailLogListResp{}, nil
}
