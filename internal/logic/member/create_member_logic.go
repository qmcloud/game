package member

import (
	"context"
	"fmt"
	"github.com/qmcloud/game/internal/model"
	"github.com/qmcloud/game/internal/svc"
	"github.com/qmcloud/game/internal/utils/dberrorhandler"
	mms "github.com/qmcloud/game/types/game"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
	"time"
)

type CreateMemberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMemberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMemberLogic {
	return &CreateMemberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMemberLogic) CreateMember(in *mms.MemberInfo) (*mms.BaseUUIDResp, error) {
	if in.Mobile != nil {
		var Finduser = model.Members{Mobile: *in.Mobile}
		result := l.svcCtx.DB.First(&Finduser)
		if result.Error != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, result.Error, in)
		}

		if result.RowsAffected > 0 {
			return nil, errorx.NewInvalidArgumentError("login.mobileExist")
		}
	}

	if in.Email != nil {
		checkEmail := model.Members{Email: *in.Email}
		resEmail := l.svcCtx.DB.First(&checkEmail)
		if resEmail.Error != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, resEmail.Error, in)
		}

		if resEmail.RowsAffected > 0 {
			return nil, errorx.NewInvalidArgumentError("login.signupUserExist")
		}
	}
	if in.Username != nil {
		//Username: *in.Username, Nickname: "test", Email: "xxx@xx.com", Mobile: "", RankID: 1, Password: encrypt.BcryptEncrypt(*in.Password)
		user := model.Members{}
		user.Username = *in.Username
		user.Password = encrypt.BcryptEncrypt("qq407193275->TG:@qmcloud" + *in.Password)
		user.Nickname = "test"
		user.ExpiredAt = time.Now()
		user.RankID = uint64(rand.Int63n(9999999999))
		result := l.svcCtx.DB.Table("mms_members").Create(&user) // 通过数据的指针来创建
		if result.Error != nil {
			fmt.Println(result.Error)
			return nil, dberrorhandler.DefaultEntError(l.Logger, result.Error, in)
		}
		return &mms.BaseUUIDResp{Msg: i18n.CreateSuccess}, result.Error
	}

	return nil, errorx.NewInvalidArgumentError("login.loginTypeForbidden")

}
