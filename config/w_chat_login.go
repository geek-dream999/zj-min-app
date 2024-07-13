package config

type WChatLogin struct {
	AppId   string `mapstructure:"appid" json:"appid" yaml:"appid"`
	Secret  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	BaseUrl string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
}
