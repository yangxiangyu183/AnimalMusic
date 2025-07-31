package appconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

type ViperData struct {
	Mysql struct {
		Username string
		Password string
		Host     string
		Port     int
		Database string
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	AliYun struct {
		AccessKeyId     string
		AccessKeySecret string
	}
	ALiPay struct {
		PrivateKey string
		AppId      string
		NotifyURL  string
		ReturnURL  string
	}
}

type Nacos struct {
	NamespaceID string
	DataId      string
	Group       string
	IpAddr      string
	Port        int
}

var (
	Conf        ViperData
	NacosConfig Nacos
)

func GetViperData() {
	viper.SetConfigFile("D:\\GOWORK\\library\\AnimalMusic\\config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
	err = viper.Unmarshal(&Conf)

	if err != nil {
		panic("配置文件解析失败1")
	}
	err = viper.Unmarshal(&NacosConfig)

	if err != nil {
		panic("配置文件解析失败2")
	}
	fmt.Println("配置文件读取成功")
}
