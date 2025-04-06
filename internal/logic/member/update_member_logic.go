package member

import (
	"context"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/internal/utils/dberrorhandler"
	mms "github.com/qmcloud/game/types/game"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLogic {
	return &UpdateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMemberLogic) UpdateMember(in *mms.MemberInfo) (*mms.BaseResp, error) {
	query := l.svcCtx.DB.Member.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilUsername(in.Username).
		SetNotNilNickname(in.Nickname).
		SetNotNilRankID(in.RankId).
		SetNotNilMobile(in.Mobile).
		SetNotNilEmail(in.Email).
		SetNotNilAvatar(in.Avatar).
		SetNotNilWechatOpenID(in.WechatId).
		SetNotNilExpiredAt(pointy.GetTimeMilliPointer(in.ExpiredAt))

	if in.Password != nil && *in.Password != "" {
		query.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
