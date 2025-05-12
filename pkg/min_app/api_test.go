package min_app

import (
	"context"
	"fmt"
	"testing"
)

// appid: wx427abf5416a7935d
// secret: d6d34b07c9f8943ceed38104bb0b2bab
// base-url: https://api.weixin.qq.com
// 测试调用接口凭证
func TestGetAccessToken(t *testing.T) {
	cli := NewWChatMinLogin(WChatLogin{
		AppId:   "wx427abf5416a7935d",
		Secret:  "d6d34b07c9f8943ceed38104bb0b2bab",
		BaseUrl: "https://api.weixin.qq.com",
	})
	ctx := context.Background()
	res, err := cli.GetAccessToken(ctx, &AccessTokenReq{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}

// 测试小程序登录
func TestLogin(t *testing.T) {
	cli := NewWChatMinLogin(WChatLogin{
		AppId:   "wx427abf5416a7935d",
		Secret:  "d6d34b07c9f8943ceed38104bb0b2bab",
		BaseUrl: "https://api.weixin.qq.com",
	})
	ctx := context.Background()
	res, err := cli.Login(ctx, &LoginReq{
		JsCode: "021K5y0",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
