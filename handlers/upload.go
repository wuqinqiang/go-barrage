package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/wuqinqiang/chitchat/config"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var dateTime string = time.Now().UTC().Format(http.TimeFormat)

func init() {
	os.MkdirAll("./public/resource", os.ModePerm)
}

//上传文件
func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data, head, err := r.FormFile("img")
	if err != nil {
		danger(err.Error())
		fmt.Println(22)
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

	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), suffix)
	//创建文件
	filePath := "resource/" + fileName
	datfile, err := os.Create("./public/" + filePath)

	if err != nil {
		danger(err.Error())
		return
	}
	//将源文件拷贝到新文件当中
	_, err = io.Copy(datfile, data)

	if err != nil {
		return
	}
	res, _ := json.Marshal(config.StaticPath + filePath)
	//如果配置是本地 那就本地文件
	if config.FileLocal == true {
		io.WriteString(w, string(res))
		return
	}
	config := config.LoadConfig()
	//否则上传至oss
	client, err := oss.New(config.Oss.BucketUrl, config.Oss.AccessKeyID, config.Oss.AccessKeySecret)
	if err != nil {
		danger(err.Error())
	}

	bucket, err := client.Bucket(config.Oss.Bucket)
	if err != nil {
		danger(err.Error())
	}
	err = bucket.PutObjectFromFile("room/"+fileName, "./public/"+filePath)
	if err != nil {
		danger(err.Error())
	}

	url, err := bucket.SignURL("room/"+fileName, oss.HTTPGet, 600*24*100)
	if err != nil {
		danger(err.Error())
		return
	}
	ossRes, _ := json.Marshal(url)
	io.WriteString(w, string(ossRes))

}
