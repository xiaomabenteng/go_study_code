package src

type TopicModel struct {
	TopicID int `json:"id"` //该model转换为json输出时，自动转换为tag的字段名
	TopicTitle string `json:"title"`
}

func CreateTopic(id int,title string)  TopicModel{
	return TopicModel{id,title}
}

type TopicQuery struct { //将querystring 参数绑定到模型
	UserName string `json:"username" form:"username"`
	Page int `json:"page" form:"page" binding:"required"` // form 是参数绑定，binding 是设定必须绑定该参数
	PageSize int `json:"pagesize" form:"pagesize"`
}
