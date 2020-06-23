package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type App struct {
	Address   string
	Static    string
	Log       string
	StaticPath string
}

type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

type Oss struct {
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	BucketUrl       string
	IsSaveLocal       bool
}

type Configuration struct {
	App App
	Db  Database
	Oss Oss
}

var config *Configuration
var once sync.Once

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {
	//全局值加载一次 底层是通过互斥锁加原子性实现的，可以自行查看代码
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return config
}
