package token

import (
	"context"
	"github.com/qmcloud/game/ent"
	"github.com/qmcloud/game/internal/utils/dberrorhandler"
	"github.com/qmcloud/game/internal/utils/entx"
	mms "github.com/qmcloud/game/types/game"
	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"time"

	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/qmcloud/game/ent/token"

	"github.com/qmcloud/game/internal/svc"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *mms.UUIDsReq) (*mms.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		tokenData, err := tx.Token.Query().Where(token.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).All(l.ctx)

		if err != nil {
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return err
		} else {
			for _, v := range tokenData {
				expiredTime := v.ExpiredAt.Sub(time.Now())
				if expiredTime > 0 {
					err = l.svcCtx.Redis.Set(l.ctx, config.RedisTokenPrefix+v.Token, "1", expiredTime).Err()
					if err != nil {
						logx.Errorw(logmsg.RedisError, logx.Field("detail", err.Error()))
						return err
					}
				}
			}
		}

		_, err = tx.Token.Delete().Where(token.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)
		return err
	})

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
