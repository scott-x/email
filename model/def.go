package model

type EmailParam struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerHost string `json:"server_host`
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	ServerPort int `json:"server_port"`
	// FromEmail　发件人邮箱地址
	FromEmail string `json:"from_email"`
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string `json:"from_passwd"`
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	Toers string `json:"to_ers"`
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCers string `json:"cc_ers"`
}
