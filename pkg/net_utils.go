package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

// ExtractResponse 解析响应结果
func ExtractResponse(response *resty.Response) error {
	if response.IsError() {
		return fmt.Errorf("request failed [%d]: %s",
			response.StatusCode(),
			response.Error())
	}

	if response.Body() == nil {
		return errors.New("result is nil")
	}
	var result BaseResponse
	err := json.Unmarshal(response.Body(), &result)
	if err != nil {
		return err
	}
	if result.Code != RequestSuccess {
		return errors.New(result.Msg)
	}
	return nil
}

// StructToMap 结构体转换为map
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

type BaseResponse struct {
	// 0为成功，其他都是失败
	Code string `json:"code,omitempty"`
	// 返回消息
	Msg string `json:"msg,omitempty"`
	// 数据
	Data interface{} `json:"data,omitempty"`
}

const (
	RequestSuccess = "0" //请求成功
)
