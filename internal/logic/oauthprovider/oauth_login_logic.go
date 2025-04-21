package oauthprovider

import (
	"context"
	mms "github.com/qmcloud/game/types/game"
	"golang.org/x/oauth2"

	"github.com/qmcloud/game/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

var providerConfig = make(map[string]oauth2.Config)

// userInfoURL used to store infoURL in database | 用来存储获取用户信息网址数据
var userInfoURL = make(map[string]string)

type OauthLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOauthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OauthLoginLogic {
	return &OauthLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OauthLoginLogic) OauthLogin(in *mms.OauthLoginReq) (*mms.OauthRedirectResp, error) {
	return nil, nil
	/*
		var config oauth2.Config
		if v, ok := providerConfig[in.Provider]; ok {
			config = v
		} else {
			p, err := l.svcCtx.DB.OauthProvider.Query().Where(oauthprovider.NameEQ(in.Provider)).First(l.ctx)
			if err != nil {
				return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
			}

			providerConfig[p.Name] = oauth2.Config{
				ClientID:     p.ClientID,
				ClientSecret: p.ClientSecret,
				Endpoint: oauth2.Endpoint{
					AuthURL:   p.AuthURL,
					TokenURL:  p.TokenURL,
					AuthStyle: oauth2.AuthStyle(p.AuthStyle),
				},
				RedirectURL: p.RedirectURL,
				Scopes:      strings.Split(p.Scopes, " "),
			}
			config = providerConfig[p.Name]

			if _, ok := userInfoURL[p.Name]; !ok {
				userInfoURL[p.Name] = p.InfoURL
			}

		}

		url := config.AuthCodeURL(in.State)

		return &mms.OauthRedirectResp{Url: url}, nil

	*/
}
