package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/wuqinqiang/chitchat/models"
	"html/template"
	"io"
	"net/http"
)

type ChatInfo struct {
	CurrentId   int
	Records     [] models.LastRecord
	Friends     [] models.Friend
	UnReadCount int  //未读消息
	UnHandleCount    int  //未处理请求
}
func unescaped (str string) template.HTML { return template.HTML(str) }


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
			Friends:     user.GetUserFriends(),
			CurrentId:   user.Id, //当前登录用户id
			UnReadCount: user.SumUnRead(),
			UnHandleCount:    user.SumUnHandle(),
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

//历史记录
func UserMessages(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger(err.Error())
			return
		}
		to_id := r.PostFormValue("to_id")
		messages, err := models.GetUserMessagesAll(sess.UserId, to_id)
		if err != nil {
			danger(err.Error())
			return
		}
		b, err := json.Marshal(messages)
		io.WriteString(w, string(b))
	}
}
