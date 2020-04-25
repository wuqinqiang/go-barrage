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
	//单聊
	{
		"chatIndex",
		"GET",
		"/chat/index",
		handlers.ChatIndex,
	},
	//单聊
	{
		"chatMessage",
		"POST",
		"/chat/chatAll",
		handlers.UserMessages,
	},
	//获取所有用户
	{
		"getUsers",
		"GET",
		"/users",
		handlers.Users,
	},
	//查找指定用户
	{
		"findUserByName",
		"POST",
		"/user",
		handlers.FindUser,
	},
	//加好友
	{
		"AddUserFriend",
		"POST",
		"/friend/crete",
		handlers.CreateFriend,
	},
	//获取审核列表
	{
		"getApplications",
		"GET",
		"/user/apps",
		handlers.GetApplications,
	},
	//获取审核列表
	{
		"handleApp",
		"POST",
		"/user/handleApp",
		handlers.HandleApplication,
	},
}
