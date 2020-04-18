package handlers

import (
	"fmt"
	"github.com/wuqinqiang/chitchat/models"
	"net/http"
)

// GET /threads/new
// 创建群组页面
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "auth.navbar", "new.thread")
	}

}

//POST //threads/new
//创建群组的页面
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			fmt.Println("cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			fmt.Println("create thread error")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

//GET /thread/read
//通过 ID 渲染指定群页面
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		error_message(writer,request,"cannot read thread")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(writer, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
