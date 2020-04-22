package handlers

import (
	"fmt"
	"github.com/wuqinqiang/chitchat/models"
	"html/template"
	"net/http"
)

type ChatInfo struct {
	CurrentId int
	Records   [] models.LastRecord
	Friends   [] models.Friend
}

func ChatIndex(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.User()
		if err != nil {
			danger(err.Error())
			return
		}
		chatInfo := ChatInfo{
			Records:   user.GetReCordFriends(),
			Friends:   user.GetUserFriends(),
			CurrentId: user.Id, //当前登录用户id
		}
		tem, err := template.ParseFiles("views/qq.html")

		if err != nil {
			fmt.Println("访问未定义页面", err)
			return
		}
		// 利用给定数据渲染模板，并将结果写入w
		tem.Execute(w, chatInfo)
	}

}
