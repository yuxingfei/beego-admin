package conf

// ueditor配置数据
type Ueditor struct {
	ImageActionName         string `mapstructure:"image_action_name" json:"image_action_name" yaml:"image_action_name"`
	ImageFieldName          string `mapstructure:"image_field_name" json:"image_field_name" yaml:"image_field_name"`
	ImageMaxSize            int    `mapstructure:"image_max_size" json:"image_max_size" yaml:"image_max_size"`
	ImageAllowFiles         string `mapstructure:"image_allow_files" json:"image_allow_files" yaml:"image_allow_files"`
	ImageCompressEnable     bool   `mapstructure:"image_compress_enable" json:"image_compress_enable" yaml:"image_compress_enable"`
	ImageCompressBorder     int    `mapstructure:"image_compress_border" json:"image_compress_border" yaml:"image_compress_border"`
	ImageInsertAlign        string `mapstructure:"image_insert_align" json:"image_insert_align" yaml:"image_insert_align"`
	ImageUrlPrefix          string `mapstructure:"image_url_prefix" json:"image_url_prefix" yaml:"image_url_prefix"`
	ImagePathFormat         string `mapstructure:"image_path_format" json:"image_path_format" yaml:"image_path_format"`
	ScrawlActionName        string `mapstructure:"scrawl_action_name" json:"scrawl_action_name" yaml:"scrawl_action_name"`
	ScrawlFieldName         string `mapstructure:"scrawl_field_name" json:"scrawl_field_name" yaml:"scrawl_field_name"`
	ScrawlPathFormat        string `mapstructure:"scrawl_path_format" json:"scrawl_path_format" yaml:"scrawl_path_format"`
	ScrawlMaxSize           int    `mapstructure:"scrawl_max_size" json:"scrawl_max_size" yaml:"scrawl_max_size"`
	ScrawlUrlPrefix         string `mapstructure:"scrawl_url_prefix" json:"scrawl_url_prefix" yaml:"scrawl_url_prefix"`
	ScrawlInsertAlign       string `mapstructure:"scrawl_insert_align" json:"scrawl_insert_align" yaml:"scrawl_insert_align"`
	ScrawlAllowFiles        string `mapstructure:"scrawl_allow_files" json:"scrawl_allow_files" yaml:"scrawl_allow_files"`
	SnapscreenActionName    string `mapstructure:"snapscreen_action_name" json:"snapscreen_action_name" yaml:"snapscreen_action_name"`
	SnapscreenPathFormat    string `mapstructure:"snapscreen_path_format" json:"snapscreen_path_format" yaml:"snapscreen_path_format"`
	SnapscreenUrlPrefix     string `mapstructure:"snapscreen_url_prefix" json:"snapscreen_url_prefix" yaml:"snapscreen_url_prefix"`
	SnapscreenInsertAlign   string `mapstructure:"snapscreen_insert_align" json:"snapscreen_insert_align" yaml:"snapscreen_insert_align"`
	CatcherLocalDomain      string `mapstructure:"catcher_local_domain" json:"catcher_local_domain" yaml:"catcher_local_domain"`
	CatcherActionName       string `mapstructure:"catcher_action_name" json:"catcher_action_name" yaml:"catcher_action_name"`
	CatcherFieldName        string `mapstructure:"catcher_field_name" json:"catcher_field_name" yaml:"catcher_field_name"`
	CatcherPathFormat       string `mapstructure:"catcher_path_format" json:"catcher_path_format" yaml:"catcher_path_format"`
	CatcherUrlPrefix        string `mapstructure:"catcher_url_prefix" json:"catcher_url_prefix" yaml:"catcher_url_prefix"`
	CatcherMaxSize          int    `mapstructure:"catcher_max_size" json:"catcher_max_size" yaml:"catcher_max_size"`
	CatcherAllowFiles       string `mapstructure:"catcher_allow_files" json:"catcher_allow_files" yaml:"catcher_allow_files"`
	VideoActionName         string `mapstructure:"video_action_name" json:"video_action_name" yaml:"video_action_name"`
	VideoFieldName          string `mapstructure:"video_field_name" json:"video_field_name" yaml:"video_field_name"`
	VideoPathFormat         string `mapstructure:"video_path_format" json:"video_path_format" yaml:"video_path_format"`
	VideoUrlPrefix          string `mapstructure:"video_url_prefix" json:"video_url_prefix" yaml:"video_url_prefix"`
	VideoMaxSize            int    `mapstructure:"video_max_size" json:"video_max_size" yaml:"video_max_size"`
	VideoAllowFiles         string `mapstructure:"video_allow_files" json:"video_allow_files" yaml:"video_allow_files"`
	FileActionName          string `mapstructure:"file_action_name" json:"file_action_name" yaml:"file_action_name"`
	FileFieldName           string `mapstructure:"file_field_name" json:"file_field_name" yaml:"file_field_name"`
	FilePathFormat          string `mapstructure:"file_path_format" json:"file_path_format" yaml:"file_path_format"`
	FileUrlPrefix           string `mapstructure:"file_url_prefix" json:"file_url_prefix" yaml:"file_url_prefix"`
	FileMaxSize             int    `mapstructure:"file_max_size" json:"file_max_size" yaml:"file_max_size"`
	FileAllowFiles          string `mapstructure:"file_allow_files" json:"file_allow_files" yaml:"file_allow_files"`
	ImageManagerActionName  string `mapstructure:"image_manager_action_name" json:"image_manager_action_name" yaml:"image_manager_action_name"`
	ImageManagerListPath    string `mapstructure:"image_manager_list_path" json:"image_manager_list_path" yaml:"image_manager_list_path"`
	ImageManagerListSize    string `mapstructure:"image_manager_list_size" json:"image_manager_list_size" yaml:"image_manager_list_size"`
	ImageManagerUrlPrefix   string `mapstructure:"image_manager_url_prefix" json:"image_manager_url_prefix" yaml:"image_manager_url_prefix"`
	ImageManagerInsertAlign string `mapstructure:"image_manager_insert_align" json:"image_manager_insert_align" yaml:"image_manager_insert_align"`
	ImageManagerAllowFiles  string `mapstructure:"image_manager_allow_files" json:"image_manager_allow_files" yaml:"image_manager_allow_files"`
	FileManagerActionName   string `mapstructure:"file_manager_action_name" json:"file_manager_action_name" yaml:"file_manager_action_name"`
	FileManagerListPath     string `mapstructure:"file_manager_list_path" json:"file_manager_list_path" yaml:"file_manager_list_path"`
	FileManagerUrlPrefix    string `mapstructure:"file_manager_url_prefix" json:"file_manager_url_prefix" yaml:"file_manager_url_prefix"`
	FileManagerListSize     int    `mapstructure:"file_manager_list_size" json:"file_manager_list_size" yaml:"file_manager_list_size"`
	FileManagerAllowFiles   string `mapstructure:"file_manager_allow_files" json:"file_manager_allow_files" yaml:"file_manager_allow_files"`
}
