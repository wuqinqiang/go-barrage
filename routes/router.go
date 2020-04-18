package routes

import "github.com/gorilla/mux"

func NewRoute() *mux.Router {
	//创建路由示例
	r := mux.NewRouter().StrictSlash(true)
	for _, route := range webRouters {
		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).Handler(route.HandlerFunc)
	}

	return r
}
