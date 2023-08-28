package service

import "github.com/gin-gonic/gin"

type ChatHistory struct {
	MessageList []Message `json:"message_list"` // 用户列表
	StatusCode  string    `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`   // 返回状态描述
}

// Message
type Message struct {
	Content    string `json:"content"`     // 消息内容
	CreateTime string `json:"create_time"` // 消息发送时间 yyyy-MM-dd HH:MM:ss
	ID         int64  `json:"id"`          // 消息id
}

/*
example:
参数名	   位置	   类型	   必填	 说明
token      query  string  是    用户鉴权token
to_user_id query  string  是    对方用户id


{
    "status_code": "string",
    "status_msg": "string",
    "message_list": [
        {
            "id": 0,
            "content": "string",
            "create_time": "string"
        }
    ]
}
*/
func GetChatHistory(c *gin.Context) {

}
