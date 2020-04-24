package models

import (
	"fmt"
	"time"
)

type Application struct {
	Id        int
	FromId    int
	ToId      int
	UserId    int
	Type      int
	Status    int
	Remark    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//加好友或者加群请求 add_type =1 表示好友请求 2 表示加群请求
func AddApplication(user User, to_id string, add_type int) (err error) {
	parerSql := "insert into applications (from_id,to_id,user_id,type,status,remark,created_at,updated_at) values(?,?,?,?,?,?,?,?)"
	statExce, err := Db.Prepare(parerSql)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	remark := fmt.Sprintf("%s%s",user.Name,"请求添加您为好友")
	defer statExce.Close()
	if add_type == 1 {
		statExce.Exec(user.Id, to_id, to_id, add_type, 0, remark, time.Now(), time.Now())
	}
	return
}
