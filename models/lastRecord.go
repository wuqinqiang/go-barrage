package models

import (
	"fmt"
	"time"
)

type LastRecord struct {
	Id          int
	SendTime    time.Time
	FromId      int
	ToId        int
	Content     string
	Type        int
	SendName    string
	ContentType int
}

func CreateLastRecord(send_name string, record ChatRecord) (err error) {
	stamOut, err := Db.Prepare("select * from last_records where (from_id =? and to_id=?) or (from_id =? and to_id=?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	last := LastRecord{}

	defer stamOut.Close()
	stamOut.QueryRow(record.FromId, record.ToId, record.ToId, record.FromId).Scan(&last.Id, &last.SendTime, &last.FromId, &last.ToId, &last.Content, &last.Type, &last.SendName, &last.ContentType)

	if last != (LastRecord{}) { //删除掉老的一条 插入新的一条 很沙雕对不对？
		stamdelete, err := Db.Prepare("delete  from last_records  where (from_id =? and to_id=?) or (from_id =? and to_id=?) ")
		if err !=nil{
			fmt.Println(err.Error())
		}
		defer stamdelete.Close()
		stamdelete.Exec(record.FromId, record.ToId, record.ToId, record.FromId)
	}

	statement := "insert into last_records (send_time,from_id,to_id,content,type,send_name,content_type) value(?,?,?,?,?,?,?)"
	stmin, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmin.Close()
	_, err = stmin.Exec(record.SendTime, record.FromId, record.ToId, record.Content, record.Type, send_name, record.ContentType)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
