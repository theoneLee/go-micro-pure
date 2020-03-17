package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/util/log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql驱动
	c "test.lee/common/cfg"
)

var db *gorm.DB

// GetDb search库 配置的DB_NAME
func GetDb() *gorm.DB {
	return db
}

// InitMysql initializes the database instance
func InitMysql() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Config.DbInfo.DbUser,
		c.Config.DbInfo.DbPassword,
		c.Config.DbInfo.DbHost,
		c.Config.DbInfo.DbName))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	db.SingularTable(true)

	db.DB().SetConnMaxLifetime(time.Second * 50)
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)
	log.Logf(c.Config.DbInfo.DbName, " 连接成功\n\n")
}
