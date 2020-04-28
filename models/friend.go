package models

type Friend struct {
	Id            int
	UserId        int
	FriendId      int
	FriendName    string
	UnreadMessage int
}
