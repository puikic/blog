package util

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

var (
	ProjectRootPath = path.Dir(getOnCurrentPath()+"/../") + "/"
)

func getOnCurrentPath() string {
	_, filename, _, _ := runtime.Caller(0) //0表示当前本行代码在什么位置
	return path.Dir(filename)              //返回文件所在目录
}

func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	configPath := ProjectRootPath + "config/"
	config.AddConfigPath(configPath) //文件所在目录
	config.SetConfigName(file)       //文件名
	config.SetConfigType("yaml")     //文件类型
	configFile := configPath + file + ".yaml"
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("找不到配置文件：%s", configFile))
		} else {
			panic(fmt.Errorf("解析配置文件%s出错:%s", configFile, err))
		}
	}
	return config
}
