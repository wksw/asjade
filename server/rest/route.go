package rest

// Route rest请求路由定义
type Route struct {
	Name string
	// 请求方式, GET, PUT, POST, DELETE
	Method string
	// 请求路由
	Path string
	// 处理方法
	Handle func(ctx *Context)
	// 描述
	Desc string
	// 其他一些源数据
	Metadata map[string]interface{}
}

// Handler 处理者需要实现的方法
type Handler interface {
	Routes() []*Route
}
