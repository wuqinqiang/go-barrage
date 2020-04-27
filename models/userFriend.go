package models

import (
	"database/sql"
	"fmt"
)

type UserFriend struct {
	Id            int
	UserId        int
	FriendId      int
	FriendName    string
	UnreadMessage int
}

//增加未读消息

func AddUnreadMessage(user_id int, friend_id int) (err error) {
	statemt := "update user_friends set unread_message=unread_message+1 where user_id=? and friend_id=? "
	parpe, err := Db.Prepare(statemt)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer parpe.Close()

	_, err = parpe.Exec(friend_id, user_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

//添加好友关系 双向关系 所以需要插入两条
func AddFriends(from_id int, user_id int, tx *sql.Tx) (err error) {
	statemt := "insert into user_friends (user_id,friend_id,friend_name,unread_message) values(?,?,?,?)"
	statemt2 := "insert into user_friends (user_id,friend_id,friend_name,unread_message) values(?,?,?,?)"

	fromUser, _ := UserByID(from_id)
	toUser, _ := UserByID(user_id)

	parpe1, _ := tx.Prepare(statemt)
	parpe2, _ := tx.Prepare(statemt2)
	defer parpe1.Close()
	defer parpe2.Close()

	fmt.Println(fromUser)
	_, err = parpe1.Exec(from_id, user_id, toUser.Name, 0)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	_, err = parpe2.Exec(user_id, from_id, fromUser.Name, 0)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
	}
	tx.Commit()
	return

}
