package token

import (
	"context"
	mms "github.com/qmcloud/game/types/game"

	"github.com/qmcloud/game/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTokenLogic) CreateToken(in *mms.TokenInfo) (*mms.BaseUUIDResp, error) {
	return nil, nil
	/*
		result, err := l.svcCtx.DB.Token.Create().
			SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
			SetNotNilUUID(uuidx.ParseUUIDStringToPointer(in.Uuid)).
			SetNotNilToken(in.Token).
			SetNotNilSource(in.Source).
			SetNotNilUsername(in.Username).
			SetNotNilExpiredAt(pointy.GetTimeMilliPointer(in.ExpiredAt)).
			Save(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		return &mms.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil

	*/
}
