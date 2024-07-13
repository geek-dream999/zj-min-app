package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

/**
通用post请求、get请求，传参 uri  domain
返回结果 *map[string]interface{}
如果有其它需求基于这一底层进行封装
*/

func Post(data map[string]interface{}, uri string, domain string) (resp *map[string]interface{}, err error) {
	response, err := resty.New().
		SetTimeout(2 * time.Second).
		SetBaseURL(domain). // 设置域名
		EnableTrace().
		SetDebug(true).
		NewRequest().
		SetBody(data).
		SetResult(&resp).
		Post(uri)
	if err != nil {
		return resp, err
	}

	//校验
	err = ExtractResponse(response)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Get(data map[string]string, uri string, domain string) (resp *map[string]interface{}, err error) {
	response, err := resty.New().
		SetTimeout(2 * time.Second).
		SetBaseURL(domain). // 设置域名
		EnableTrace().
		SetDebug(true).
		NewRequest().
		SetQueryParams(data).
		Get(uri)
	if err != nil {
		return resp, err
	}

	//校验
	err = ExtractResponse(response)
	if err != nil {
		return nil, err
	}

	// 解析响应体到 resp 变量
	if err = json.Unmarshal(response.Body(), &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// ExtractResponse 校验响应结果
func ExtractResponse(response *resty.Response) error {
	if response.IsError() {
		return fmt.Errorf("request failed [%d]: %s",
			response.StatusCode(),
			response.Error())
	}

	if response.Body() == nil {
		return errors.New("result is nil")
	}

	return nil
}

// StructToMap 结构体转换为map[string]interface
func StructToMap(s interface{}, m interface{}) error {
	if s == nil {
		return nil
	}
	// Convert map to json string
	jsonStr, err := json.Marshal(s)
	if err != nil {
		return err
	}

	// Convert json string to struct
	if err = json.Unmarshal(jsonStr, &m); err != nil {
		return err
	}
	return nil
}

// StructToStringMap 结构体转换为map[string][string]
func StructToStringMap(src interface{}, target map[string]string) error {
	if src == nil {
		return errors.New("source is nil")
	}

	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonStr, &target)
	if err != nil {
		return err
	}

	return nil
}

type BaseResponse struct {
	// 0为成功，其他都是失败
	Code string `json:"errcode,omitempty"`
	// 返回消息
	Msg string `json:"errmsg,omitempty"`
	// 数据
	Data interface{} `json:"data,omitempty"`
}

const (
	RequestSuccess    = "0" //请求成功
	RequestSuccessInt = 0   //请求成功

)
