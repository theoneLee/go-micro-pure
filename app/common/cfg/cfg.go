package cfg

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
)

// Config ：外部获取配置字段：cfg.Config.DbInfo.DbUser
var Config *Conf

// Conf 全局配置结构体
type Conf struct {
	DbInfo    *DbInfo // todo 新增配置需要新增字段。
	MongoInfo *MongoInfo
}

type MongoInfo struct {
	url string
}

//DbInfo :mysql所需内容
type DbInfo struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

//InitConfig :
func InitConfig(path string) {
	Config = &Conf{}
	// 加载配置文件
	err := config.Load(file.NewSource(
		file.WithPath(path),
		//file.WithPath("./config/config.json"),
	))
	if err != nil {
		panic(err)
	}
	//var dbInfo DbInfo
	if err := config.Get("mysql", "database1").Scan(&Config.DbInfo); err != nil {
		panic(err)
	}
	//fmt.Println(Config.DbInfo)
	//Config.DbInfo = &dbInfo

	if err := config.Get("mongo", "database1").Scan(&Config.MongoInfo); err != nil {
		panic(err)
	}
	// todo 新增配置需要为新增字段进行解析。

	//fmt.Println(host.Name, host.Address, host.Port)
}
