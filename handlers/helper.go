package handlers

import (
	"errors"
	"fmt"
	"github.com/wuqinqiang/chitchat/models"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var logger *log.Logger

//通过Cookie 判断用户是否已登录
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}

	}
	return
}

//解析 HTML 模板 (应对需要传入多个模板文件的情况，避免重复编写模板)
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files [] string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

//生成响应 HTML
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	//时间转换
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)

	//之前
	//templates := template.Must(template.ParseFiles(files...))
	//templates.ExecuteTemplate(w, "layout", data)
}

func init() {
	file, err := os.OpenFile("logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

//不命名为error 避免和error 类型重名
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNNIN ")
	logger.Println(args...)
}

func error_message(write http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(write, request, strings.Join(url, ""), 302)

}

func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}

//返回版本号
func Version() string {
	return "0.1"
}

//获取ip地址

func RemoteIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}


