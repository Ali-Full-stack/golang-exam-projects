package email

import (
	"fmt"
	"gateway/internal/models"
	"log"

	"gopkg.in/gomail.v2"
)

func SendEmail(email, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "juraboevizzatillo5@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Registration Status")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "juraboevizzatillo5@gmail.com", "tnkx qlnz anuk exdn")

	if err := d.DialAndSend(m); err != nil {
		log.Println("failed to send an email:", err)
		return err
	}

	return nil
}

func RegisterClient(id, name string)string{
	body := fmt.Sprintf(`
    <html>
    <body>
        <p>Hello <strong>%v</strong>,</p>
        <p>I Hope you are doing Well ..</p>
        <p>You are successfully registered to <strong>Secure Delivery Service  Website</strong> </p>
        <p>This is your special  ID: <strong>%v</strong></p>
        <p>You need this ID in order to ordering products</p>
        <p><strong>PLEASE DO NOT SHARE WITH ANYONE</strong></p>
        <p>Thanks and have a nice day </p>
        <p>from <strong>Ali Team</strong>.........</p>
    </body>
</html>
    `, name, id)
	return body
}

func SendNewAdmin(id, name string)string{
	body := fmt.Sprintf(`
    <html>
    <body>
		<h1>Congratulations New Admin <strong>%v</strong> </h1>
        <p>I Hope you are doing Well ..</p>
        <p>You are successfully Joined our <strong>Ali Team </strong> as admin </p>
        <p>This is your special  ID: <strong>%v</strong></p>
        <p>You need this ID in order to maintain the website gracefully </p>
        <p><strong>PLEASE DO NOT SHARE WITH ANYONE</strong></p>
        <p>Thanks and have a nice day </p>
        <p>from <strong>Ali Team</strong>.........</p>
    </body>
</html>
    `, name,id)
	return body
}

func RegisterDriver(id, name string)string{
	body := fmt.Sprintf(`
    <html>
    <body>
		<h1>Congratulations New Driver <strong>%v</strong> </h1>
        <p>I Hope you are doing Well ..</p>
        <p>You are successfully Joined our <strong>Ali Team </strong> as Driver </p>
        <p>This is your special  ID: <strong>%v</strong></p>
        <p>You need this ID in order to get an access to application </p>
        <p><strong>PLEASE DO NOT SHARE WITH ANYONE</strong></p>
        <p>Thanks and have a nice day </p>
        <p>from <strong>Ali Team</strong>.........</p>
    </body>
</html>
    `, name,id)
	return body
}

func SendOrderConfirmationToClient()string{
        return fmt.Sprintf(`
    <html>
    <body>
        <h1>Thank you for your order from ALI ONLINE SHOPPING</h1>
    
        <p>We're delighted to confirm your recent order with ALI ONLINE SHOPPING.</p>
    
        <You can track the status of your order by logging into your account on our website.</p>
    
        <p>If you have any questions or need further assistance, please don't hesitate to contact our customer service team at alishopping@gmail.com.</p>
    
        <p>Thank you again for choosing ALI ONLINE SHOPPING. We hope you're satisfied with your purchase!</p>
    
        <p>Best regards,<br>
        ALI team....</p>
    </body>
    </html>
    `)
}
func SendDriverConfirmationOrder(cl models.Clientcontact,name, id string, charge float32) string {
    body := fmt.Sprintf(`
<html>
<body>
    <h1>Congratulations, <strong>%s</strong>!</h1>
    <p>We're excited to have you on board as a driver with our team.</p>
    <p>This is your special ID: <strong>%s</strong></p>
    <p>You'll need this ID to access the application, so please keep it safe and <strong>do not share it with anyone</strong>.</p>
    <hr>
    <h2>Client Information</h2>
    <p>
        Phone: %s<br>
        Email: %s<br>
        City: %s<br>
        Region: %s<br>
        Home Address: %s
    </p>
    <hr>
    <h2>Charge for this Delivery</h2>
    <p>You will receive $%.2f for this delivery.</p>
    <p>Thank you for your hard work, and have a great day!</p>
    <p>- The Ali Team</p>
</body>
</html>
`, name, id, cl.Phone, cl.Email, cl.City, cl.Region, cl.HomeAddress, charge)

    return body
}