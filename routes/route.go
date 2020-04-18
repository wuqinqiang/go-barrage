package routes

import (
	"github.com/wuqinqiang/chitchat/handlers"
	"net/http"
)

//定义一个WebRoute 结构体用于存放单个路由

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//声明WebRouters 切片存放所有的web路由
type WebRouters []WebRoute

//定义所有的web路由
var webRouters = WebRouters{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"ws",
		"GET",
		"/ws",
		handlers.WsContent,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	{
		"newThread",
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	{
		"createThread",
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},
	{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},
	{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
	{
		"chat",
		"GET",
		"/chat",
		handlers.ChatRoom,
	},
	{
		"postChat",
		"POST",
		"/chat/post",
		handlers.SendMessage,
	},
	{
		"messageAll",
		"GET",
		"/chat/messages",
		handlers.MessageAll,
	},
}
