package emaillog

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogLogic {
	return &UpdateEmailLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmailLogLogic) UpdateEmailLog(in *game.EmailLogInfo) (*game.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &game.BaseResp{}, nil
}
