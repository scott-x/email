package util

import (
	"encoding/json"
	"github.com/go-gomail/gomail"
	"github.com/scott-x/email/model"
	"strings"
	"io/ioutil"
)

// 全局变量，因为发件人账号、密码，需要在发送时才指定
// 注意，由于是小写，外面的包无法使用
var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message

func ParseConfig() (*model.EmailParam,error){
	bs, err:=ioutil.ReadFile("config.json")
	if err!=nil {
		return nil,err
	}
	ep := &model.EmailParam{}
	err=json.Unmarshal(bs,ep)
	if err!=nil {
		return nil, err
	}
	return ep, nil
}

func InitEmail(ep *model.EmailParam) {

	toers := []string{}

	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromPasswd = ep.FromPasswd

	m = gomail.NewMessage()

	if len(ep.Toers) == 0 {
		return
	}

	for _, tmp := range strings.Split(ep.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)

	//抄送列表
	if len(ep.CCers) != 0 {
		for _, tmp := range strings.Split(ep.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", fromEmail, "")
}
func getToken() {

}
// SendEmail body支持html格式字符串
func SendEmail(subject, body string) {

	// 主题
	m.SetHeader("Subject", subject)

	// 正文
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
