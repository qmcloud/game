package token

import (
	"context"
	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"
	"github.com/zeromicro/go-zero/core/logx"
)

type BlockUserAllTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBlockUserAllTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BlockUserAllTokenLogic {
	return &BlockUserAllTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BlockUserAllTokenLogic) BlockUserAllToken(in *mms.UUIDReq) (*mms.BaseResp, error) {
	return nil, nil
	/*
		err := l.svcCtx.DB.Token.Update().Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).SetStatus(0).Exec(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		tokenData, err := l.svcCtx.DB.Token.Query().
			Where(token.UUIDEQ(uuidx.ParseUUIDString(in.Id))).
			Where(token.StatusEQ(0)).
			Where(token.ExpiredAtGT(time.Now())).
			All(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		for _, v := range tokenData {
			expiredTime := v.ExpiredAt.Sub(time.Now())
			if expiredTime > 0 {
				err = l.svcCtx.Redis.Set(l.ctx, config.RedisTokenPrefix+v.Token, "1", expiredTime).Err()
				if err != nil {
					logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
					return nil, errorx.NewInternalError(i18n.RedisError)
				}
			}
		}

		return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil

	*/
}
