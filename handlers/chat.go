package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func ChatIndex(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	}else {

		tem,err := template.ParseFiles("views/qq.html")

		if err != nil{
			fmt.Println("读取文件失败,err",err)
			return
		}
		// 利用给定数据渲染模板，并将结果写入w
		tem.Execute(w,nil)
	}

}
