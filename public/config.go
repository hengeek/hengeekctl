package public

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// Configuration 项目配置
type Configuration struct {
	TcSecretID  string `yaml:"tc_secretId"`
	TcSecretKey string `yaml:"tc_secretKey"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		data, err := ioutil.ReadFile("config.yml")
		// 此处如果使用go-bindata将本地配置文件转成二进制，那么使用如下代码读取配置文件
		// data, err := Asset("config.yml")
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatal(err)
		}

		// 如果环境变量有配置，读取环境变量
		tc_secretId := os.Getenv("tc_secretId")
		if tc_secretId != "" {
			config.TcSecretID = tc_secretId
		}
		tc_secretKey := os.Getenv("tc_secretKey")
		if tc_secretKey != "" {
			config.TcSecretKey = tc_secretKey
		}
	})

	return config
}
