package memberrank

import (
	"context"
	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberRankListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberRankListLogic {
	return &GetMemberRankListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberRankListLogic) GetMemberRankList(in *mms.MemberRankListReq) (*mms.MemberRankListResp, error) {
	return nil, nil
	/*
		var predicates []predicate.MemberRank
		if in.Name != nil {
			predicates = append(predicates, memberrank.NameContains(*in.Name))
		}
		if in.Description != nil {
			predicates = append(predicates, memberrank.DescriptionContains(*in.Description))
		}
		if in.Remark != nil {
			predicates = append(predicates, memberrank.RemarkContains(*in.Remark))
		}
		result, err := l.svcCtx.DB.MemberRank.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		resp := &mms.MemberRankListResp{}
		resp.Total = result.PageDetails.Total

		for _, v := range result.List {
			resp.Data = append(resp.Data, &mms.MemberRankInfo{
				Id:          pointy.GetPointer(v.ID),
				CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				Name:        &v.Name,
				Description: &v.Description,
				Remark:      &v.Remark,
				Code:        &v.Code,
			})
		}

		return resp, nil

	*/
}
