package emaillog

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmailLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEmailLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmailLogLogic {
	return &DeleteEmailLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEmailLogLogic) DeleteEmailLog(in *game.UUIDsReq) (*game.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &game.BaseResp{}, nil
}
