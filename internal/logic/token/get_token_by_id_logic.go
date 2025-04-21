package token

import (
	"context"
	mms "github.com/qmcloud/game/types/game"

	"github.com/qmcloud/game/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTokenByIdLogic) GetTokenById(in *mms.UUIDReq) (*mms.TokenInfo, error) {
	return nil, nil
	/*
		result, err := l.svcCtx.DB.Token.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		return &mms.TokenInfo{
			Id:        pointy.GetPointer(result.ID.String()),
			CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(result.Status)),
			Uuid:      pointy.GetPointer(result.UUID.String()),
			Token:     &result.Token,
			Source:    &result.Source,
			Username:  &result.Username,
			ExpiredAt: pointy.GetPointer(result.ExpiredAt.UnixMilli()),
		}, nil

	*/
}
