package models

import (
	"database/sql"
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

func CreateLastRecord(send_name string, record ChatRecord,tx *sql.Tx) (err error) {

	stamOut, err := Db.Prepare("select * from last_records where (from_id =? and to_id=?) or (from_id =? and to_id=?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	last := LastRecord{}

	defer stamOut.Close()
	stamOut.QueryRow(record.FromId, record.ToId, record.ToId, record.FromId).Scan(&last.Id, &last.SendTime, &last.FromId, &last.ToId, &last.Content, &last.Type, &last.SendName, &last.ContentType)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	if last != (LastRecord{}) { //删除掉老的一条 插入新的一条 很沙雕对不对？
		stamdelete, err := tx.Prepare("delete  from last_records  where (from_id =? and to_id=?) or (from_id =? and to_id=?) ")
		defer stamdelete.Close()
		_,err=stamdelete.Exec(record.FromId, record.ToId, record.ToId, record.FromId)
		if err !=nil{
			tx.Rollback()
		}
	}

	statement := "insert into last_records (send_time,from_id,to_id,content,type,send_name,content_type) value(?,?,?,?,?,?,?)"
	stmin, err := tx.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	defer stmin.Close()
	_, err = stmin.Exec(record.SendTime, record.FromId, record.ToId, record.Content, record.Type, send_name, record.ContentType)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	tx.Commit()
	return
}
