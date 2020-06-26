package e

var MsgFlags = map[int]string{
	SUCCESS:        "操作成功",
	ERROR:          "操作失败",
	INVALID_PARAMS: "请求参数错误",
	FORBIDDEN:      "禁止访问",
	NOLOGIN:        "未登录",
	USERERROR:      "用户不存在或者密码错误",
	NOCATEGORY:     "不存在该分类",
	NOARTICLE:      "不存在该文章",
	EXISTSARTICLE:  "该分类下存在文章",
	NOTALK:         "不存在该说说",
}
