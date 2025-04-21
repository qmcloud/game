package memberrank

import (
	"context"
	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRankByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankByIdLogic {
	return &GetMemberRankByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRankByIdLogic) GetMemberRankById(in *mms.IDReq) (*mms.MemberRankInfo, error) {
	return nil, nil
	/*result, err := l.svcCtx.DB.MemberRank.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.MemberRankInfo{
		Id:          pointy.GetPointer(result.ID),
		CreatedAt:   pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:   pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Name:        &result.Name,
		Code:        &result.Code,
		Description: &result.Description,
		Remark:      &result.Remark,
	}, nil
	*/
}
