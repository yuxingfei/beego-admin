package conf

// 其他额外配置
type Other struct {
	LogAesKey string `mapstructure:"log_aes_key" json:"log_aes_key" yaml:"log_aes_key"` //日志加密key
}
