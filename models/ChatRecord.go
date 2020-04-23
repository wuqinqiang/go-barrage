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

func GetUserMessagesAll(user_id int, to_id string) (chats []ChatRecord,err error) {
	rows, err := Db.Query("select * from chat_records where (from_id=? and to_id=?) or (from_id=? and to_id=?)", user_id, to_id, to_id, user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		record := ChatRecord{}
		if err = rows.Scan(&record); err != nil {
			fmt.Println(err.Error())
			return
		}
		chats = append(chats, record)
	}
	return
}
