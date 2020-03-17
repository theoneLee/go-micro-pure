package model

import (
	"github.com/micro/go-micro/v2/util/log"
	"test.lee/user/model/util"
)

//Notice : 红点通知.todo 当回复评论时，向消息队列发送插入到User的NoticeList
type Notice struct {
	ID         string `gorm:"column:id;primary_key" json:"id"`
	Category   string `json:"category"` //todo 消息类别
	Title      string `json:"title"`
	Content    string `json:"content"`
	FromUserId string `json:"from_user_id"` //消息发送方id

	UserId string `json:"user_id" gorm:"column:user_id"` //外键
}

//TableName 数据库中的表名
func (Notice) TableName() string {
	return "user_notice"
}

// User
type User struct {
	ID         string   `gorm:"column:id;primary_key" json:"id"`
	UserName   string   `gorm:"column:user_name" json:"user_name"`
	Password   string   `json:"-" gorm:"column:password"`
	NoticeList []Notice `json:"notice_list" gorm:"foreignkey:user_id"`
}

//TableName 数据库中的表名
func (User) TableName() string {
	return "user_info"
}

//todo 获取用户消息顺带获取notice；save notice

type UserDao interface {
	SaveUser(user User)
	SaveNotice(notice Notice)
	GetUser(id string) *User
}

type UserSQLDao struct {
}

func (u UserSQLDao) SaveUser(user User) {
	err := util.GetDb().Debug().Create(user).Error
	if err != nil {
		panic(err)
	}
}

func (u UserSQLDao) SaveNotice(notice Notice) {
	err := util.GetDb().Debug().Create(notice).Error
	if err != nil {
		panic(err)
	}
}

func (u UserSQLDao) GetUser(id string) *User {
	var user, user2 User
	// preload 查询一对多
	err := util.GetDb().Debug().Preload("NoticeList").First(&user, "id=?", id).Error
	if err != nil {
		panic(err)
	}
	log.Log("preload:", user)

	//related 查询一对多
	err = util.GetDb().Debug().First(&user2, "id=?", id).Error
	if err != nil {
		panic(err)
	}
	err = util.GetDb().Debug().Model(&user2).Related(&user2.NoticeList).Find(&user2.NoticeList).Error
	if err != nil {
		panic(err)
	}
	log.Log("relate:", user2)
	return &user
}
