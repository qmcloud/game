package base

import (
	"context"
	"github.com/qmcloud/game/internal/svc"
	mms "github.com/qmcloud/game/types/game"
	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *mms.Empty) (*mms.BaseResp, error) {
	//测试nats
	return &mms.BaseResp{
		Msg: "grpc return ok",
	}, nil
}

// insert init member data
func (l *InitDatabaseLogic) insertMemberData() error {
	return nil
	/*
		user := model.Member{Username: "normalMember", Nickname: "Normal Member", Email: "xxx@qq.com", Mobile: "18888888888", RankID: 1, Password: encrypt.BcryptEncrypt("qmcloud")}
		result := l.svcCtx.DB.Create(&user) // 通过数据的指针来创建
		if result.Error != nil {
			logx.Errorw("failed to insert member data for initialization", logx.Field("detail", result.Error))
			return dberrorhandler.DefaultEntError(l.Logger, result.Error, nil)
		} else {
			return nil
		}
	*/
}

// insert init member rank data
func (l *InitDatabaseLogic) insertMemberRankData() error {
	return nil
	/*var memberRanks []*ent.MemberRankCreate
	memberRanks = append(memberRanks, l.svcCtx.DB.MemberRank.Create().
		SetName("memberRank.normal").
		SetCode("001").
		SetDescription("普通会员 | Normal Member").
		SetRemark("普通会员 | Normal Member"),
	)

	memberRanks = append(memberRanks, l.svcCtx.DB.MemberRank.Create().
		SetName("memberRank.vip").
		SetCode("002").
		SetDescription("VIP").
		SetRemark("VIP"),
	)

	err := l.svcCtx.DB.MemberRank.CreateBulk(memberRanks...).Exec(l.ctx)
	if err != nil {
		logx.Errorw("failed to insert member rank data for initialization", logx.Field("detail", err))
		return dberrorhandler.DefaultEntError(l.Logger, err, nil)
	} else {
		return nil
	}*/
}

func (l *InitDatabaseLogic) insertProviderData() error {
	return nil
	/*var providers []*ent.OauthProviderCreate

	providers = append(providers, l.svcCtx.DB.OauthProvider.Create().
		SetName("google").
		SetClientID("your client id").
		SetClientSecret("your client secret").
		SetRedirectURL("http://localhost:3100/oauth/login/callback").
		SetScopes("email openid").
		SetAuthURL("https://accounts.google.com/o/oauth2/auth").
		SetTokenURL("https://oauth2.googleapis.com/token").
		SetAuthStyle(1).
		SetInfoURL("https://www.googleapis.com/oauth2/v2/userinfo?access_token=TOKEN"),
	)

	providers = append(providers, l.svcCtx.DB.OauthProvider.Create().
		SetName("github").
		SetClientID("your client id").
		SetClientSecret("your client secret").
		SetRedirectURL("http://localhost:3100/oauth/login/callback").
		SetScopes("email openid").
		SetAuthURL("https://github.com/login/oauth/authorize").
		SetTokenURL("https://github.com/login/oauth/access_token").
		SetAuthStyle(2).
		SetInfoURL("https://api.github.com/user"),
	)

	err := l.svcCtx.DB.OauthProvider.CreateBulk(providers...).Exec(l.ctx)
	if err != nil {
		logx.Errorw("failed to insert member's oauth provider data for initialization", logx.Field("detail", err))
		return dberrorhandler.DefaultEntError(l.Logger, err, nil)
	} else {
		return nil
	}*/
}
