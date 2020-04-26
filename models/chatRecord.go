package models

import (
	"fmt"
	"time"
)

type ChatRecord struct {
	Id          int
	Uuid        string
	SendTime    time.Time
	Content     string
	FromId      int
	ToId        int
	Type        int
	ContentType int
}

func GetUserMessagesAll(user_id int, to_id string) (chats []ChatRecord, err error) {
	rows, err := Db.Query("select * from chat_records where (from_id=? and to_id=?) or (from_id=? and to_id=?)", user_id, to_id, to_id, user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		record := ChatRecord{}
		if err = rows.Scan(&record.Id, &record.Uuid, &record.SendTime, &record.Content, &record.FromId, &record.ToId, &record.Type, &record.ContentType); err != nil {
			fmt.Println(err.Error())
			return
		}
		chats = append(chats, record)
	}

	//还需要把未读消息清空
	statemet := "update user_friends set unread_message=0 where  user_id=? and friend_id=?"
	paper, err := Db.Prepare(statemet)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer paper.Close()
	paper.Exec(user_id, to_id)
	return
}
