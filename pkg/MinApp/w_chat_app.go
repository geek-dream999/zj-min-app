package MinApp

import (
	"context"
	"meet_directly/pkg"
)

type WChatLogin struct {
	AppId   string
	Secret  string
	BaseUrl string
}

func NewWChatMinLogin(c WChatLogin) MinLoginClient {
	return &c
}

// GetAccessToken 调用接口凭证
func (c *WChatLogin) GetAccessToken(ctx context.Context, req *AccessTokenReq) (resp *AccessTokenResp, err error) {
	req.GrantType = ApiGrantType

	//req 转成map[string]interface{}格式
	data := c.getData(req)

	//调用微信小程序api
	res, err := WChatLoginPost(data, GetAccessTokenUrl, c.BaseUrl)
	if err != nil {
		return nil, err
	}

	//转成响应体
	if err = pkg.StructToMap(res, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Login 小程序登录
func (c *WChatLogin) Login(ctx context.Context, req *LoginReq) (resp *LoginResp, err error) {
	req.GrantType = GrantType

	//req 转成map[string]string格式
	data := c.getDataToString(req)

	res, err := WChatLoginGet(data, LoginUrl, c.BaseUrl)
	if err != nil {
		return nil, err
	}

	if err = pkg.StructToMap(res, &resp); err != nil {
		return nil, err
	}
	return nil, nil
}

// PhoneVerify 手机号快速验证
func (c *WChatLogin) PhoneVerify(ctx context.Context, req *PhoneVerifyReq) (resp *PhoneVerifyResp, err error) {

	//req 转成map[string]interface{}格式
	data := c.getData(req)

	//调用微信小程序api
	res, err := WChatLoginPost(data, PhoneVerifyUrl, c.BaseUrl)
	if err != nil {
		return nil, err
	}

	//转成响应体
	if err = pkg.StructToMap(res, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *WChatLogin) getData(in interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	err := pkg.StructToMap(in, &data)
	if err != nil {
		return nil
	}
	data["appid"] = c.AppId
	data["secret"] = c.Secret
	return data
}

func (c *WChatLogin) getDataToString(in interface{}) map[string]string {
	data := make(map[string]string)
	err := pkg.StructToMap(in, &data)
	if err != nil {
		return nil
	}
	data["appid"] = c.AppId
	data["secret"] = c.Secret
	return data
}

// 入参出参
type (
	AccessTokenResp struct {
		AccessToken string `json:"access_token"` //获取到的凭证
		ExpiresIn   int    `json:"expires_in"`   //凭证有效时间，单位：秒。目前是7200秒之内的值。
	}

	AccessTokenReq struct {
		AppId        string `json:"appid"`         //账号唯一凭证，即 AppID，可在「微信公众平台 - 设置 - 开发设置」页中获得。（需要已经成为开发者，且账号没有异常状态）
		Secret       string `json:"secret"`        //账号唯一凭证密钥，即 AppSecret，获取方式同 appid
		GrantType    string `json:"grant_type"`    //填写 client_credential
		ForceRefresh bool   `json:"force_refresh"` //非必传，默认使用 false。1. force_refresh = false 时为普通调用模式，access_token 有效期内重复调用该接口不会更新 access_token；2. 当force_refresh = true 时为强制刷新模式，会导致上次获取的 access_token 失效，并返回新的 access_token
	}

	LoginResp struct {
		SessionKey string `json:"session_key"` //会话密钥
		UnionId    string `json:"unionid"`     //用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台账号下会返回，详见 UnionID 机制说明。
		OpenId     string `json:"openid"`      //用户唯一标识
		ErrCode    int    `json:"errcode"`     //错误码
	}

	LoginReq struct {
		AppId     string `json:"appid"`      //小程序 appId
		Secret    string `json:"secret"`     //小程序 appSecret
		JsCode    string `json:"js_code"`    //登录时获取的 code，可通过wx.login获取
		GrantType string `json:"grant_type"` //授权类型，此处只需填写 authorization_code
	}

	PhoneVerifyResp struct {
		ErrCode   int       `json:"errcode"`    //错误码
		ErrMsg    string    `json:"errmsg"`     //错误信息
		PhoneInfo PhoneInfo `json:"phone_info"` //用户手机号信息
	}
	PhoneInfo struct {
		PhoneNumber     string    `json:"phoneNumber"`     //用户手机号
		PurePhoneNumber string    `json:"purePhoneNumber"` //没有区号的手机号
		CountryCode     string    `json:"countryCode"`     //区号
		Watermark       Watermark `json:"watermark"`       //数据水印
	}
	Watermark struct {
		Timestamp int    `json:"timestamp"` //用户获取手机号操作的时间戳
		AppId     string `json:"appid"`     //小程序appid
	}

	PhoneVerifyReq struct {
		AccessToken string `json:"access_token"` //接口调用凭证，该参数为 URL 参数，非 Body 参数。使用access_token或者authorizer_access_token
		Code        string `json:"code"`         //手机号获取凭证
		OpenId      string `json:"openid"`       //非必传，用户唯一标识
	}
)

// uri
const (
	GetAccessTokenUrl = "/cgi-bin/stable_token"
	LoginUrl          = "/sns/jscode2session"
	PhoneVerifyUrl    = "/wxa/business/getuserphonenumber"
)

// 固定参数
const (
	GrantType    = "authorization_code" //小程序登录
	ApiGrantType = "client_credential"  //调用api凭证
)
