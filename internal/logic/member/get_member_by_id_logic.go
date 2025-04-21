package member

import (
	"context"
	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByIdLogic {
	return &GetMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByIdLogic) GetMemberById(in *mms.UUIDReq) (*mms.MemberInfo, error) {
	return nil, nil
	/*
		result, err := l.svcCtx.DB.Member.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		return &mms.MemberInfo{
			Id:        pointy.GetPointer(result.ID.String()),
			CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(result.Status)),
			Username:  &result.Username,
			Password:  &result.Password,
			Nickname:  &result.Nickname,
			RankId:    &result.RankID,
			Mobile:    &result.Mobile,
			Email:     &result.Email,
			Avatar:    &result.Avatar,
			ExpiredAt: pointy.GetPointer(result.ExpiredAt.UnixMilli()),
		}, nil

	*/
}
