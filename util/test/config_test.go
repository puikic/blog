package test

import (
	"blog/util"
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestConfig(t *testing.T) {
	dbViper := util.CreateConfig("mysql")
	dbViper.WatchConfig() //动态感知配置变化
	//读取配置的第一种方式
	if !dbViper.IsSet("blog.port") { //判断此项是否已配置
		t.Fail()
	}
	port := dbViper.GetInt("blog.port")
	fmt.Println("port:", port)
	// time.Sleep(10 * time.Second) //10s内修改一下配置文件，测试viper能否感知读取最新值
	// port = dbViper.GetInt("blog.port")
	// fmt.Println("port:", port)

	//读取配置的第二种方式
	logViper := util.CreateConfig("log")
	logViper.WatchConfig()
	type LogConfig struct {
		Level string `mapstruture:"level"` // tag
		File  string `mapstructure:"file"`
	}
	var logConfig LogConfig
	if err := logViper.Unmarshal(&logConfig); err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		fmt.Println(logConfig.Level)
		fmt.Println(logConfig.File)
	}
}

// TestCaller() -> bf()  ->  af()     56 -> 51 -> 46
func af() (string, int) {
	//参数换成0、1、2试试
	_, filename, line, _ := runtime.Caller(0) //0
	return filename, line
}

func bf() (string, int) {
	return af() //1
}
func TestCaller(t *testing.T) {
	filename, line := bf() //2
	fmt.Println(filename, line)
	fmt.Println(path.Dir(filename) + "/../../")
	fmt.Println(path.Dir(path.Dir(filename) + "/../../"))
}

// go test -v .\util\test\ -run=^TestConfig$ -count=1
// go test -v .\util\test\ -run=^TestCaller$ -count=1
