package member

import (
	"context"
	"fmt"
	"github.com/qmcloud/game/internal/model"
	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/internal/utils/dberrorhandler"
	mms "github.com/qmcloud/game/types/game"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMemberByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberByUsernameLogic {
	return &GetMemberByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberByUsernameLogic) GetMemberByUsername(in *mms.UsernameReq) (*mms.MemberInfo, error) {
	user := model.Members{Username: in.Username}
	result := l.svcCtx.DB.Table("mms_members").First(&user)
	if result.Error != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, result.Error, in)
	}
	fmt.Println(user)
	ID := uint32(user.ID)
	return &mms.MemberInfo{
		Id:        &ID,
		CreatedAt: pointy.GetPointer(user.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(user.UpdatedAt.UnixMilli()),
		Status:    pointy.GetPointer(uint32(user.Status)),
		Username:  &user.Username,
		Password:  &user.Password,
		Nickname:  &user.Nickname,
		RankId:    &user.RankID,
		Mobile:    &user.Mobile,
		Email:     &user.Email,
		Avatar:    &user.Avatar,
		ExpiredAt: pointy.GetPointer(user.ExpiredAt.UnixMilli()),
	}, nil
}
