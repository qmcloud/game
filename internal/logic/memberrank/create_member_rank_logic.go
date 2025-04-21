package memberrank

import (
	"context"

	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMemberRankLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMemberRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMemberRankLogic {
	return &CreateMemberRankLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMemberRankLogic) CreateMemberRank(in *mms.MemberRankInfo) (*mms.BaseIDResp, error) {
	return nil, nil
	/*result, err := l.svcCtx.DB.MemberRank.Create().
		SetNotNilName(in.Name).
		SetNotNilCode(in.Code).
		SetNotNilDescription(in.Description).
		SetNotNilRemark(in.Remark).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil

	*/
}
