package conf

import (
	"beego-admin/global"
	"beego-admin/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

//加载配置文件
func init() {
	var config string
	if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
		config = utils.ConfigFile
		fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
	} else {
		config = configEnv
		fmt.Printf("您正在使用BA_CONFIG环境变量,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.BA_CONFIG); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&global.BA_CONFIG); err != nil {
		panic(err)
	}

}
