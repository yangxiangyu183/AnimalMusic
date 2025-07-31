package inits

import (
	"Music_api/appconfig"
	"Music_service/global"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitAll() {
	NacosInit()
	InitZap()
	MysqlInit()
	GetViperData()
	InitGrpc()
}

func InitGrpc() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//config.UserClient = __.NewUserClient(conn)
	zap.S().Info("grpc is connected")
}

func InitZap() {
	zapConfig := zap.NewDevelopmentConfig() //初始化开发环境配置
	zapConfig.OutputPaths = []string{
		"stdout",
		"D:\\GOWORK\\library\\AnimalMusic\\Music_service\\exam.log",
	}
	logger, _ := zapConfig.Build()
	zap.ReplaceGlobals(logger)
}
func MysqlInit() {
	var err error
	data := appconfig.Conf.Mysql
	fmt.Println(data.Username, "		1111111111111111111111111")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", data.Username, data.Password, data.Host, data.Port, data.Database)
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql is fail!!!")
	}
	fmt.Println("mysql is success!!!")

}

func GetViperData() {
	v := viper.New()
	v.SetConfigFile("D:\\GOWORK\\library\\AnimalMusic\\config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
	//nacosConfig:=config.Nac
	//err = v.Unmarshal()

	if err != nil {
		panic("配置文件解析失败")
	}
	fmt.Println("配置文件读取成功")
}

func NacosInit() {
	fmt.Println(appconfig.NacosConfig.NamespaceID, 2222222222222)
	clientConfig := constant.ClientConfig{
		NamespaceId:         appconfig.NacosConfig.NamespaceID, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      appconfig.NacosConfig.IpAddr,
			ContextPath: "/nacos",
			Port:        uint64(appconfig.NacosConfig.Port),
			Scheme:      "http",
		},
	}
	configClient, _ := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: appconfig.NacosConfig.DataId,
		Group:  appconfig.NacosConfig.Group})

	app := appconfig.Conf
	json.Unmarshal([]byte(content), &app)

	appconfig.Conf = app
	fmt.Println(appconfig.Conf)
}
