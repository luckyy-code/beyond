package logic

import (
	"beyond/application/applet/internal/code"
	"beyond/application/user/rpc/user"
	"beyond/pkg/encrypt"
	"beyond/pkg/jwt"
	"beyond/pkg/xcode"
	"context"
	"strings"

	"beyond/application/applet/internal/svc"
	"beyond/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	//将最前面和最后面的ASCII定义的空格去掉
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.LoginMobileEmpty
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}
	err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		return nil, err
	}
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("EncMobile mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{Mobile: mobile})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u == nil || u.UserId == 0 {
		return nil, xcode.AccessDenied
	}

	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": u.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)

	return &types.LoginResponse{
		UserId: u.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
