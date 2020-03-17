package main

import (
	c "test.lee/common/cfg"
	"test.lee/user/model/util"
)

func main() {
	// todo grpc service
	c.InitConfig("./cfg/config.yml")

	util.InitMysql()

}
