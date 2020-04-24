package handlers

import (
	"encoding/json"
	"github.com/wuqinqiang/chitchat/models"
	"io"
	"net/http"
)

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
	user, _:= sess.User()
	friend_id := r.PostFormValue("user_id")
	if err := models.AddApplication(user, friend_id, 1); err != nil {
		danger(err.Error())
		return
	}
	io.WriteString(w,string(friend_id))
}
