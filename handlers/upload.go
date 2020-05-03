package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func init() {
	os.MkdirAll("./public/resource", os.ModePerm)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data, head, err := r.FormFile("img")
	if err != nil {
		danger(err.Error())
		return
	}
	suffix := ".png"
	srcFilename := head.Filename
	split := strings.Split(srcFilename, ".")
	if len(split) > 1 {
		suffix = "." + split[len(split)-1]
	}

	fileType := r.FormValue("filetype")
	if len(fileType) > 0 {
		suffix = fileType
	}

	fileName:=fmt.Sprintf("%d%s",time.Now().Unix(),suffix)
	//创建文件
	filePath:="resource/"+fileName
	datfile,err:=os.Create("./public/"+filePath)

	if err !=nil{
		danger(err.Error())
		return
	}
	//将源文件拷贝到新文件当中
	_,err=io.Copy(datfile,data)

	if err !=nil{
		danger(err.Error())
		return
	}
	res,_:=json.Marshal(filePath)
	io.WriteString(w,string(res))
}
