package email

import (
	"github.com/go-gomail/gomail"
	"github.com/scott-x/email/model"
	"github.com/scott-x/email/util"
	"log"
	"strings"
)


// 全局变量，因为发件人账号、密码，需要在发送时才指定
// 注意，由于是小写，外面的包无法使用

var (
	m *gomail.Message
	config *model.EmailParam
	err error
	serverHost, fromEmail, fromPasswd string
	serverPort int

)
func init() {
	config, err = util.ParseConfig()
	if err!=nil {
		log.Printf("init configuration error: %s\n",err)
		panic(err)
	}
	toers := []string{}

	serverHost = config.ServerHost
	serverPort = config.ServerPort
	fromEmail = config.FromEmail
	fromPasswd = config.FromPasswd
	m = gomail.NewMessage()

	if len(config.Toers) == 0 {
		return
	}

	for _, tmp := range strings.Split(config.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)

	//抄送列表
	if len(config.CCers) != 0 {
		for _, tmp := range strings.Split(config.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", fromEmail, "")
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