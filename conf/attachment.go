package conf

// 上传附件配置
type Attachment struct {
	ThumbPath    string `mapstructure:"thumb_path" json:"thumb_path" yaml:"thumb_path"`          //缩略图路径
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                            //上传目录配置（相对于根目录）
	Url          string `mapstructure:"url" json:"url" yaml:"url"`                               //url（相对于web目录）
	ValidateSize string `mapstructure:"validate_size" json:"validate_size" yaml:"validate_size"` //默认不超过50mb
	ValidateExt  string `mapstructure:"validate_ext" json:"validate_ext" yaml:"validate_ext"`    //url（相对于web目录）
}
