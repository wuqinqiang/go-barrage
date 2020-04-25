package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/wuqinqiang/chitchat/models"
	"io"
	"net/http"
	"strconv"
)

type AppHandle struct {
	app_type    int //1申请 2同意
	application models.Application
}

//根据name获取用户
func FindUser(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	err = r.ParseForm()
	if err != nil {
		danger(err.Error())
		return
	}
	name := r.PostFormValue("name")
	user, err := models.UserName(name)
	if err != nil {
		danger(err.Error())
		io.WriteString(w, "")
	}
	res, err := json.Marshal(user)
	if err != nil {
		danger(err.Error())
		return
	}
	io.WriteString(w, string(res))
}

//获取所有的用户
func Users(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	users, err := models.Users(sess.UserId)
	if err != nil {
		danger(err.Error())
		return
	}
	res, err := json.Marshal(users)
	if err != nil {
		danger(err.Error())
		return
	}
	io.WriteString(w, string(res))
}

//加好友
func CreateFriend(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	err = r.ParseForm()
	if err != nil {
		danger(err.Error())
		return
	}
	user, _ := sess.User()
	friend_id := r.PostFormValue("user_id")
	to_id, _ := strconv.Atoi(friend_id)
	state, err := models.AddApplication(user, to_id, 1)
	fmt.Println(state)
	if err != nil {
		danger(err.Error())
		return
	}

	//如果提交成功 并且对方在线的话 发送通知消息给他
	if state == 1 && user_clients[to_id] != nil {
		app := models.GetApplicationByUser(user.Id, to_id, 1)
		err = user_clients[to_id].WriteJSON(app)
		if err != nil {
			danger(err.Error())
			return
		}
	}

	res, err := json.Marshal(state)
	io.WriteString(w, string(res))
}

//获取申请列表
func GetApplications(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	apps := models.GetUserApplications(sess.UserId, 1)
	res, err := json.Marshal(apps)
	io.WriteString(w, string(res))
}

//同意拒绝好友申请
func HandleApplication(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}
	param := r.PostFormValue("from_id")
	param1 := r.PostFormValue("status")
	from_id, _ := strconv.Atoi(param)
	status, _ := strconv.Atoi(param1)
	app := models.HandleApp(from_id, sess.UserId, 1, status)
	res, _ := json.Marshal(app)
	io.WriteString(w, string(res))
}
