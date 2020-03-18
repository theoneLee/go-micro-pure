package model

import "test.lee/user/model/util"

//News :
type News struct {
	ID          string    `gorm:"column:id;primary_key" json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CommentList []Comment `json:"comment_list" gorm:"foreignkey:news_id"` //一对多

	Status int `json:"status"` // -1为待审核。1为已拒绝。2为已通过
}

//TableName 数据库中的表名
func (News) TableName() string {
	return "user_notice"
}

type Comment struct {
	ID     string `gorm:"column:id;primary_key" json:"id"`
	NewsId string `json:"news_id" gorm:"column:news_id"` //外键

	UserId  string `json:"user_id"` //发言用户id
	Content string `json:"content"`

	ReplyList []Reply `json:"reply_list" gorm:"foreignkey:comment_id"`
}

//TableName 数据库中的表名
func (Comment) TableName() string {
	return "user_notice"
}

type Reply struct {
	ID        string `gorm:"column:id;primary_key" json:"id"`
	CommentId string `json:"comment_id" gorm:"column:comment_id"` //外键

	Content  string `json:"content"`
	UserId   string `json:"user_id"`    //发言用户id
	AtUserId string `json:"at_user_id"` //被回复的userId
}

//TableName 数据库中的表名
func (Reply) TableName() string {
	return "user_notice"
}

// todo 添加新闻 审核新闻（这个审核是一个服务监听消息队列一个topic，然后将其status更改为已审核。） 添加/删除评论 添加评论回复
type NewsDAO interface {
	SaveNews(news News) error
	UpdateNews(news News) error
	SaveComment(comment Comment) error
	DeleteComment(comment Comment) error

	SaveReply(reply Reply) error

	FindNews(newsID string) *News
	FindComment(commentID string) *Comment
	FindCommentEager(commentID string) *Comment // 顺带拉取reply list

	FindNewsList() []News //只返回新闻列表
	FindCommentList(newsID string) []Comment
}

type NewsSQLDAO struct {
}

func (n NewsSQLDAO) FindCommentEager(commentID string) *Comment {
	//panic("implement me")
	var c Comment
	err := util.GetDb().Debug().Preload("ReplyList").First(&c, "id=?", commentID).Error
	if err != nil {
		panic(err)
	}
	return &c
}

func (n NewsSQLDAO) SaveNews(news News) error {
	return util.GetDb().Debug().Create(news).Error
}

func (n NewsSQLDAO) UpdateNews(news News) error {
	//panic("implement me")
	return util.GetDb().Debug().Model(&News{}).Where("id=?", news.ID).Updates(news).Error
}

func (n NewsSQLDAO) SaveComment(comment Comment) error {
	//若要一并保存comment.ReplyList,使用save
	return util.GetDb().Debug().Create(comment).Error
}

func (n NewsSQLDAO) DeleteComment(comment Comment) error {
	//panic("implement me")
	return util.GetDb().Debug().Model(&Comment{}).Delete(comment, "id =?", comment.ID).Error
}

func (n NewsSQLDAO) SaveReply(reply Reply) error {
	//panic("implement me")
	return util.GetDb().Debug().Create(reply).Error
}

func (n NewsSQLDAO) FindNews(newsID string) *News {
	//panic("implement me")
	var news News
	err := util.GetDb().Where("id=?", newsID).First(&news).Error
	if err != nil {
		panic(err)
	}
	return &news
}

func (n NewsSQLDAO) FindComment(commentID string) *Comment {
	//panic("implement me")
	var comment Comment
	err := util.GetDb().Where("id=?", commentID).First(&comment).Error
	if err != nil {
		panic(err)
	}
	return &comment
}

func (n NewsSQLDAO) FindNewsList() []News {
	//panic("implement me")
	var list []News
	err := util.GetDb().Debug().Preload("CommentList").Find(&list).Error
	if err != nil {
		panic(err)
	}
	return list
}

func (n NewsSQLDAO) FindCommentList(newsID string) []Comment {
	//panic("implement me")
	var list []Comment
	err := util.GetDb().Debug().Preload("ReplyList").Find(&list, "news_id").Error
	if err != nil {
		panic(err)
	}
	return list
}
