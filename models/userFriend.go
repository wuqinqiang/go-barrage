package models

import "fmt"

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

	_, err = parpe.Exec(user_id, friend_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
