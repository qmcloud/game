package oauthprovider

import (
	"context"
	mms "github.com/qmcloud/game/types/game"

	"github.com/qmcloud/game/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetOauthProviderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOauthProviderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOauthProviderListLogic {
	return &GetOauthProviderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOauthProviderListLogic) GetOauthProviderList(in *mms.OauthProviderListReq) (*mms.OauthProviderListResp, error) {
	return nil, nil
	/*
		var predicates []predicate.OauthProvider
		if in.Name != nil {
			predicates = append(predicates, oauthprovider.NameContains(*in.Name))
		}
		result, err := l.svcCtx.DB.OauthProvider.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		resp := &mms.OauthProviderListResp{}
		resp.Total = result.PageDetails.Total

		for _, v := range result.List {
			resp.Data = append(resp.Data, &mms.OauthProviderInfo{
				Id:           &v.ID,
				CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				Name:         &v.Name,
				ClientId:     &v.ClientID,
				ClientSecret: &v.ClientSecret,
				RedirectUrl:  &v.RedirectURL,
				Scopes:       &v.Scopes,
				AuthUrl:      &v.AuthURL,
				TokenUrl:     &v.TokenURL,
				AuthStyle:    &v.AuthStyle,
				InfoUrl:      &v.InfoURL,
			})
		}

		return resp, nil

	*/
}
