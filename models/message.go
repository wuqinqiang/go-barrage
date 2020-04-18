package models

import (
	"time"
)

type Message struct {
	Id        int
	Message   string
	Uuid      string
	UserName  string
	IP        string
	Type      int
	UserId    int
	CreatedAt time.Time
}

func Messages() (messages [] Message, err error) {
	rows, err := Db.Query("select id,message,uuid,user_name,ip,type,user_id,created_at from messages order by created_at desc ")
	if err != nil {
		return
	}
	for rows.Next() {
		msg := Message{}
		if err = rows.Scan(
			&msg.Id, &msg.Message, &msg.Uuid,
			&msg.UserName,&msg.IP,&msg.Type, &msg.UserId, &msg.CreatedAt); err != nil {
			return
		}
		messages = append(messages, msg)
	}
	rows.Close()
	return
}
