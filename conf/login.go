package conf

// 登录设置:后台登录相关设置
type Login struct {
	Token      string `mapstructure:"token" json:"token" yaml:"token"`                //登录token验证
	Captcha    string `mapstructure:"captcha" json:"captcha" yaml:"captcha"`          //验证码
	Background string `mapstructure:"background" json:"background" yaml:"background"` //登录背景
}
