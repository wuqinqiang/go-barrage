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
//status 1表示请求成功 2已经是好友关系 3已提交过申请
func AddApplication(user User, to_id int, add_type int) (status int, err error) {

	status = 1
	var friend UserFriend
	//查询是否已是好友关系
	Db.QueryRow("select * from user_friends where (user_id=? and friend_id=?) or (user_id=? and friend_id=?)", user.Id, to_id, to_id, user.Id).
		Scan(&friend.Id, &friend.UserId, &friend.FriendId, &friend.FriendName, &friend.UnreadMessage);
	if friend.Id != 0 {
		status = 2
		return
	}
	//查询是否已经提交过好友申请,如果是不作处理
	application := GetApplicationByUser(user.Id, to_id, add_type)
	if application.Id != 0 {
		status = 3
		return
	}

	//什么都没有增加一条申请记录
	parerSql := "insert into applications (from_id,to_id,user_id,type,status,remark,created_at,updated_at) values(?,?,?,?,?,?,?,?)"
	statExce, err := Db.Prepare(parerSql)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	remark := fmt.Sprintf("%s%s", user.Name, "请求添加您为好友")
	defer statExce.Close()
	if add_type == 1 {
		statExce.Exec(user.Id, to_id, to_id, add_type, 0, remark, time.Now(), time.Now())
	}
	return
}

//通过用户查看是否存在未处理的审批
func GetApplicationByUser(user_id int, to_id int, add_type int) (application Application) {
	Db.QueryRow("select * from applications where from_id=? and to_id=? and type=? and status=?", user_id, to_id, add_type, 0).
		Scan(&application.Id, &application.FromId, &application.ToId, &application.UserId, &application.Type,
			&application.Status, &application.Remark, &application.CreatedAt, &application.UpdatedAt)
	return
}

//所有未处理的列表
func GetUserApplications(user_id int, add_type int) (apps [] Application) {
	rows, err := Db.Query("select * from applications where user_id=? and type=? and status=? order by created_at", user_id, add_type, 0)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		app := Application{}
		if err = rows.Scan(&app.Id, &app.FromId, &app.ToId,
			&app.UserId, &app.Type, &app.Status, &app.Remark, &app.CreatedAt, &app.UpdatedAt); err != nil {
			fmt.Println(err.Error())
			return
		}
		apps = append(apps, app)
	}
	return
}

//同意拒绝
func HandleApp(from_id int, to_id int, add_type int, handle_status int) (app Application) {
	tx,_:=Db.Begin()
	state := "update applications set status=? where from_id=? and to_id=? and type=? and status=?"
	parpe, err := tx.Prepare(state)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer parpe.Close()
	parpe.Exec(handle_status, from_id, to_id, add_type, 0)

	//同意申请的话就添加好友关系
	if handle_status == 1 {
		if err := AddFriends(from_id, to_id,tx); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	Db.QueryRow("select * from applications where from_id=? and to_id=? and status >? order by updated_at", from_id, to_id, 0).
		Scan(&app.Id, &app.FromId, &app.ToId,
			&app.UserId, &app.Type, &app.Status, &app.Remark, &app.CreatedAt, &app.UpdatedAt)
	return
}
