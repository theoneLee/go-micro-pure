package model

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
