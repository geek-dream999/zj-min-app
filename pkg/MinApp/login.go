package MinApp

import "context"

/*  微信小程序登录文档地址：https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getStableAccessToken.html    */

// MinLoginClient 接口定义
type MinLoginClient interface {
	// GetAccessToken 调用接口凭证
	GetAccessToken(ctx context.Context, req *AccessTokenReq) (resp *AccessTokenResp, err error)
	// Login 小程序登录
	Login(ctx context.Context, req *LoginReq) (resp *LoginResp, err error)
	// PhoneVerify 手机号快速验证
	PhoneVerify(ctx context.Context, req *PhoneVerifyReq) (resp *PhoneVerifyResp, err error)
}
