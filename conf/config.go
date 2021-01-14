package conf

type Server struct {
	Base       Base       `mapstructure:"base" json:"base" yaml:"base"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis      Mysql      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Attachment Attachment `mapstructure:"attachment" json:"attachment" yaml:"attachment"`
	Login      Login      `mapstructure:"login" json:"login" yaml:"login"`
	Other      Other      `mapstructure:"other" json:"other" yaml:"other"`
	Ueditor    Ueditor    `mapstructure:"ueditor" json:"ueditor" yaml:"ueditor"`
}
