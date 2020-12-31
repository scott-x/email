# email
Send email with golang

### SetUp
1. download package
```bash
go get -u -v github.com/scott-x/email
```
2. configuration

Edit `email.json` in your project root directory, here is an example for QQ
```json
{
	"server_host":"smtp.exmail.qq.com",
	"server_port":465,
	"from_email":"",
	"from_passwd":"",
	"to_ers":"",
	"cc_ers":""
}
```
- `server_host`: server host
- `server_port`: server port
- `from_email`: sender
- `from_passwd`: sender email
- `to_ers`: receiver 
- `cc_ers`: the email that were CC'd

### Example

```go
package main

import (
    "github.com/scott-x/email"
)

func main() {
    test()
}

func test() {
    subject := "Release final files / WMT USA / KYE / C202018_APP"
    body := `Welcme to use email<br>
            <h3>This is title</h3>
             Hello <a href = "http://github.com/scott-x">github</a><br>`
    //send email
    email.SendEmail(subject, body)
}
```