package min_app

import (
	"errors"
	"fmt"
	"meet_directly/pkg"
)

/*
微信小程序登录post请求封装
*/

func WChatLoginPost(data map[string]interface{}, uri string, url string) (*map[string]interface{}, error) {
	res, err := pkg.Post(data, uri, url)
	if err != nil {
		return nil, err
	}

	//检查错误码
	err = checkCodeErr(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

/*
微信小程序登录get请求封装
*/

func WChatLoginGet(data map[string]string, uri string, url string) (*map[string]interface{}, error) {
	res, err := pkg.Get(data, uri, url)
	if err != nil {
		return nil, err
	}

	//检查错误码
	err = checkCodeErr(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func checkCodeErr(res *map[string]interface{}) error {
	if res == nil {
		return errors.New("response is nil")
	}
	//读取map的值，如果键不存在
	if resErr, ok := (*res)["errcode"].(float64); ok {
		if resErr != 0 {
			// 如果 errcode 存在且不为 0，获取错误信息
			if errMsg, ok := (*res)["errmsg"].(string); ok {
				return fmt.Errorf("微信登录错误，错误码: %d, 错误信息: %s", int(resErr), errMsg)
			}
			return fmt.Errorf("微信登录错误，错误码: %d", int(resErr))
		}
	}
	return nil
}
