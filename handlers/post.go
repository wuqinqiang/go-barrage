package handlers

import (
	"fmt"
	"github.com/wuqinqiang/chitchat/models"
	"net/http"
)

func PostThread(write http.ResponseWriter, request *http.Request) {
	sess, err := session(write, request)
	if err != nil {
		http.Redirect(write, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			fmt.Println("cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("cannot find user session")
		}
		uuid := request.PostFormValue("uuid")
		body := request.PostFormValue("body")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			error_message(write, request, "cannot find thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			fmt.Println("cannot create new post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(write, request, url, 302);
	}

}
