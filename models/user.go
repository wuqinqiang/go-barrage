package models

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into sessions (uuid,email,user_id,created_at)values(?,?,?,?)"
	stmin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmin.Close()
	//创建uuid
	uuid := createUUID()

	stmin.Exec(uuid, user.Email, user.Id, time.Now())
	stmtout, err := Db.Prepare("select id ,uuid,email,user_id,created_at from sessions where uuid=?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return

}

//Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("select id,uuid,email,user_id,created_at FROM sessions where user_id=?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

//create a new user
func (user *User) Create() (err error) {
	statement := "insert into users(uuid,name,email,password,created_at)values(?,?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()
	uuid := createUUID()
	stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())
	stmtout, err := Db.Prepare("select id ,uuid,created_at from users where uuid=?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

func (user *User) Delete() (err error) {
	statement := "delete from users where id=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

func (user *User) Update() (err error) {
	statement := "update users set name=?,email=?,where id=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Name, user.Email, user.Id)
	return
}

func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

func Users() (users []User, err error) {
	rows, err := Db.Query("select id,uuid,name,email,password,created_at from users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name,&user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id,uuid,name,email,password,created_at from users where email=?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id,uuid,name,email,password,created_at FROM users WHERE uuid=?", uuid).
		Scan(&user.Id, &user.Email, &user.Name, &user.Uuid, &user.CreatedAt)
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, topic, user.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, topic, user_id, created_at from threads where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := createUUID()
	stmtin.Exec(uuid, body, user.Id, conv.Id, time.Now())

	stmtout, err := Db.Prepare("select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmtout.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

//发送消息
func (user *User) CreateMessage(ip string, content string,types int) (message Message, err error) {
	statement := "insert into messages(uuid,message,user_name,type,ip,user_id,created_at) values(?,?,?,?,?,?,?)"
	stmtin, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error());
		return
	}
	uuid := createUUID()
	defer stmtin.Close()

	stmtin.Exec(uuid, content, user.Name, types,ip, user.Id, time.Now())
	stmtout, err := Db.Prepare("select * from messages where uuid=?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&message.Id, &message.Uuid, &message.Message,&message.UserName, &message.Type,&message.IP, &message.UserId, &message.CreatedAt)
	return
}
