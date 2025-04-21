package member

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLogic {
	return &DeleteMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMemberLogic) DeleteMember(in *mms.UUIDsReq) (*mms.BaseResp, error) {
	return nil, nil
	/*
		_, err := l.svcCtx.DB.Member.Delete().Where(member.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		return &mms.BaseResp{Msg: i18n.DeleteSuccess}, nil

	*/
}
