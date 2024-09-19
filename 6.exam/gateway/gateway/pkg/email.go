package pkg

import (
	"fmt"
	"log"


	"gopkg.in/gomail.v2"
)

func SendEmail(email, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "saidazimqilichov@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Registration Status")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "saidazimqilichov@gmail.com", "fler dgia ntyz mjoq")

	if err := d.DialAndSend(m); err != nil {
		log.Println("failed to send an email:", err)
		return err
	}

	return nil
}

func RegisterClient(id, name string) string {
	body := fmt.Sprintf(`
    <html>
    <body>
        <p>Hello <strong>%v</strong>,</p>
        <p>I Hope you are doing Well ..</p>
        <p>You are successfully registered to <strong>FINANCE</strong> </p>
        <p>This is your special  ID: <strong>%v</strong></p>
        <p><strong>PLEASE DO NOT SHARE WITH ANYONE</strong></p>
        <p>Thanks and have a nice day </p>
        <p>from <strong>FINANCE Team</strong>.........</p>
    </body>
</html>
    `, name, id)
	return body
}

func UserCode(code int, name string) string {
	body := fmt.Sprintf(`
    <html>
    <body>
        <p>Hello <strong>%v</strong>,</p>
        <p>I Hope you are doing Well ..</p>
        <p>This is your One Time Password : <strong>%v</strong></p>
        <p>Thanks and have a nice day </p>
        <p>from <strong>Strong Team</strong>.........</p>
    </body>
</html>
    `, name, code)
	return body
}