package usecases

// 与交付层对接的公共操作
// http状态只有为200(成功)/400(客户端错误)/500(服务端错误)
const (
	//成功
	StatusOK = 200
	//客户端错误
	StatusClientError = 400
	//服务端失败
	StatusServerError = 500

	//参数相关错误
	//参数解析错误 [来自客户端的参数解析出错]
	ErrorParameterParse = 4001
	//必要的参数缺失
	ErrorParameterDefect = 4002

	//获取错误
	//文件上传错误
	ErrorUpload = 5001
	//文件删除出错
	ErrorDelete = 5002
	//读取远程出错
	ErrorReadRemote = 5003
	//解析远程数据出错
	ErrorParseRemote = 5004
)

// Response 是交付层的基本回应
type Response struct {
	Code    int         `json:"code"`    //请求状态代码
	Message interface{} `json:"message"` //请求结果提示
	Data    interface{} `json:"data"`    //请求结果与错误原因
}

// List 会返回给交付层一个列表回应
type List struct {
	Code    int         `json:"code"`    //请求状态代码
	Count   int         `json:"count"`   //数据量
	Message interface{} `json:"message"` //请求结果提示
	Data    interface{} `json:"data"`    //请求结果
}
