package token

import (
	"context"
	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTokenLogic) UpdateToken(in *mms.TokenInfo) (*mms.BaseResp, error) {
	return nil, nil
	/*
		token, err := l.svcCtx.DB.Token.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
			SetNotNilSource(in.Source).
			Save(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		if uint8(*in.Status) == common.StatusBanned {
			expiredTime := token.ExpiredAt.Sub(time.Now())
			if expiredTime > 0 {
				err = l.svcCtx.Redis.Set(l.ctx, config.RedisTokenPrefix+token.Token, "1", expiredTime).Err()
				if err != nil {
					logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
					return nil, errorx.NewInternalError(i18n.RedisError)
				}
			}
		} else if uint8(*in.Status) == common.StatusNormal {
			err := l.svcCtx.Redis.Del(l.ctx, config.RedisTokenPrefix+token.Token).Err()
			if err != nil {
				logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
				return nil, errorx.NewInternalError(i18n.RedisError)
			}
		}

		return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil

	*/
}
