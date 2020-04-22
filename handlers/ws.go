package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/wuqinqiang/chitchat/models"
	"net/http"
	"time"
)

type Msg struct {
	Message     string    `json:"message"`
	UserName    string    `json:"user_name"`
	Type        int       `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	ContentType int       `json:"content_type"`
	To          int       `json:"to"`
}

var clients = make(map[*websocket.Conn]bool) //ws客户端

var messageChannel = make(chan interface{}) //消息通道存储

var upgrader = websocket.Upgrader{
	//HandshakeTimeout: 5,
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//当前预先任何连接
	CheckOrigin: func(r *http.Request) bool { return true },
}

//读取发送的消息
func Reader(conn *websocket.Conn, sess models.Session, r *http.Request) {
	for {
		var msg Msg
		//读取信息
		_, p, err := conn.ReadMessage()

		if err != nil {
			delete(clients, conn) //删除掉这个没用的客户端连接
			danger(err.Error())
			break
		}

		err = json.Unmarshal([]byte(string(p)), &msg)
		if err != nil {
			danger(err.Error())
			break
		}
		user, err := sess.User()
		//记录发送消息

		if msg.Type == 5 { //是弹幕发送
			if _, err := user.CreateMessage(RemoteIP(r), msg.Message, msg.Type); err != nil {
				danger(err.Error())
				break
			}
			msg.UserName = user.Name
			msg.CreatedAt = time.Now()

			messageChannel <- msg
		} else { //单聊或者群聊消息
			chat, err := user.CreateChatMessage(msg.Message, msg.To, msg.Type, msg.ContentType)
			if err != nil {
				danger(err.Error())
				break
			}
			if err := models.CreateLastRecord(user.Name,chat); err != nil {
				danger(err.Error())
				break
			}
			//找到次客户端发送消息
		}
	}
}

//获取消息
func SendClientMessage() {
	for {
		msg := <-messageChannel
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
				danger(err.Error())
				break
			}
		}
	}
}

func WsContent(w http.ResponseWriter, r *http.Request) {

	//没有登录的人不让连接
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}

	//将此连接升级为ws
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		warning(err.Error())
		return
	}
	//注册一个新的客户端
	clients[ws] = true
	fmt.Println("连接成功")
	go Reader(ws, sess, r)
	go SendClientMessage()
}
