package models

import "time"

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
