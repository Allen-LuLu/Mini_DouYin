package errmsg

// 该文件定义一些与客户端交互的错误信息

const(
	/* USER ERR*/

	USERALREADYEXITS = "用户已存在"
	GENTOKENFAILED = "无法生成token,请重试"
	STORETOKENFAILED = "存储token失败"
	QUERYTOKEN = "查询token错误"
	CHECKUSERFAILED = "请输入正确的用户名密码"
	USERNOTEXIST = "该用户不存在"
	AUTHFAILED = "权限认证失败"
)