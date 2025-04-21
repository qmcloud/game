package oauthprovider

import (
	"context"
	mms "github.com/qmcloud/game/types/game"

	"github.com/qmcloud/game/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOauthProviderLogic {
	return &DeleteOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOauthProviderLogic) DeleteOauthProvider(in *mms.IDsReq) (*mms.BaseResp, error) {
	return nil, nil
	/*
		_, err := l.svcCtx.DB.OauthProvider.Delete().Where(oauthprovider.IDIn(in.Ids...)).Exec(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		return &mms.BaseResp{Msg: i18n.DeleteSuccess}, nil

	*/
}
