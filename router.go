package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Handle func(http.ResponseWriter, *http.Request)

// 定义参数类型
type param struct {
	key   string
	value string
}

// 定义参数组
type params []param

// 参数获取值，通过参数名
func (ps params) ByName(name string) string {
	for i := range ps {
		if ps[i].key == name {
			return ps[i].value
		}
	}
	return ""
}

// 定义路由实体
type Router struct {
}

// 初始化方法
func New() *Router {
	return &Router{}
}

// 处理 路由类型
func (r *Router) Handle(method string, path string, route *Router) {
	method = strings.ToUpper(method)
	if path[0] != '/' {
		panic(fmt.Sprintf("path must begin with '/',in path :'%s'", path))
	}
}

func (r *Router) Get(path string, route *Router) {
	r.Handle("GET", path, route)
}

func (r *Router) POST(path string, route *Router) {
	r.Handle("POST", path, route)
}
