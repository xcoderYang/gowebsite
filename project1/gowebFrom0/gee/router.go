package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router{
	return &router{
		roots: make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

/**
/p/:lang/doc => ["p", ":lang", "doc"]
 */

func parsePattern(pattern string)[]string{
	vs:=strings.Split(pattern, "/")
	parts:=make([]string, 0)

	for _, item:=range vs{
		if item!= ""{
			parts = append(parts, item)
			if item[0] == '*'{
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc){
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_,ok := r.roots[method]
	if !ok{
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string){

	// 解析路径 a/b/c=>["a","b","c"]
	searchParts := parsePattern(path)
	params := make(map[string]string)

	// 通过访问方式判断 GET,POST等
	root, ok := r.roots[method]

	if !ok{
		return nil, nil
	}

	// 直接从根节点找
	n := root.search(searchParts, 0)

	if n!= nil{
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			// 如果路由以 : 开头，例如 :path
			if part[0] == ':'{
				// 将当前层级路由赋值给 path
				params[part[1:]] = searchParts[index]
			}
			// 如果路由以 * 开头 *path
			if part[0] == '*' && len(part) > 1{
				// 将路由的余下部分都赋值给 path
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}


func (r *router) handle(c *Context){
	n, params := r.getRoute(c.Method, c.Path)
	if n!= nil{
		key := c.Method + "-"+n.pattern
		c.Params = params
		c.handlers = append(c.handlers, r.handlers[key])
	}else{
		c.handlers = append(c.handlers, func(c *Context){
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}