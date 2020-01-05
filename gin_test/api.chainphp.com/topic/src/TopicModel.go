package src

import "time"
type Topics struct { //单个Topic实体
	TopicID int `json:"id" gorm:"PRIMARY_KEY"`
	TopicTitle string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIp string `json:"ip" binding:"ipv4"`
	TopicScore int `json:"score" binding:"omitempty,gt=5"`
	TopicUrl string `json:"url" binding:"omitempty,topicurl"`
	TopicDate time.Time `json:"url" binding:"required"`
}

type TopicModel struct {
	TopicID int `json:"id" gorm:"PRIMARY_KEY"` //该model转换为json输出时，自动转换为tag的字段名
	TopicTitle string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP string `json:"ip" binding:"ipv4"`
	TopicScore int `json:"score" binding:"omitempty,gt=5"`
	TopicUrl string `json:"url" binding:"omitempty,topicurl"`
	TopicDate time.Time `json:"url" binding:"required"`
}
type TopicsModel struct {
	TopicList []TopicModel `json:"topics" binding:"gt=0,lt=3,topicValidate,dive"`
	TopicListSize int `json:"size"`

}
type TopicClass struct {
	ClassId int `gorm:"PRIMARY_KEY"`
	ClassName string
	ClassRemark string
	ClassType string `gorm:"Column:classtype"` //设置数据库真实字段名
}

func CreateTopic(id int,title string)  TopicModel{
	return TopicModel{}
	//return TopicModel{id,title}
}

type TopicQuery struct { //将querystring 参数绑定到模型
	UserName string `json:"username" form:"username"`
	Page int `json:"page" form:"page" binding:"required"` // form 是参数绑定，binding 是设定一定规则绑定该参数
	PageSize int `json:"pagesize" form:"pagesize"`
}
