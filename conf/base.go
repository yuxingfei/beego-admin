package conf

//后台的基本信息设置
type Base struct {
	Name            string `mapstructure:"name" json:"name" yaml:"name"`                                     //后台名称
	ShortName       string `mapstructure:"short_name" json:"short_name" yaml:"short_name"`                   //后台简称
	Author          string `mapstructure:"author" json:"author" yaml:"author"`                               //后台作者
	Version         string `mapstructure:"version" json:"version" yaml:"version"`                            //后台版本
	Link            string `mapstructure:"link" json:"link" yaml:"link"`                                     //footer链接地址
	PasswordWarning string `mapstructure:"password_warning" json:"password_warning" yaml:"password_warning"` //默认密码警告
	ShowNotice      string `mapstructure:"show_notice" json:"show_notice" yaml:"show_notice"`                //是否显示提示信息
	NoticeContent   string `mapstructure:"notice_content" json:"notice_content" yaml:"notice_content"`       //提示信息内容
}
