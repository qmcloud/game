package memberrank

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberRankLogic {
	return &DeleteMemberRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMemberRankLogic) DeleteMemberRank(in *mms.IDsReq) (*mms.BaseResp, error) {
	return nil, nil
	/*
		_, err := l.svcCtx.DB.MemberRank.Delete().Where(memberrank.IDIn(in.Ids...)).Exec(l.ctx)

		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		return &mms.BaseResp{Msg: i18n.DeleteSuccess}, nil

	*/
}
