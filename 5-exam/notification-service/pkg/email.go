package pkg

import (
	"fmt"
	"log"

	"notify-service/protos"

	"gopkg.in/gomail.v2"
)

func RegistrationEmail(id string, user *protos.UserInfo) error {
	if err :=SendEmail(user.Email, RegisterClient(id, user.Username)); err != nil {
		return fmt.Errorf("failed to send registration email to user:%v",err)
	}
	return nil
}
func BookingConfirmationEmail(req *protos.BookingEmail)error{
	if err :=SendEmail(req.Email, BookingConfirmation(req)); err != nil {
		return fmt.Errorf("failed to send booking confirmation email to user :%v",err) 
	}
	return nil
}


func SendEmail(email, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "juraboevizzatillo5@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Registration Status")
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "juraboevizzatillo5@gmail.com", "aahx qvex xiyf ggjy")

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
        <p>You are successfully registered to <strong>Hotels</strong> </p>
        <p>This is your special  ID: <strong>%v</strong></p>
        <p><strong>PLEASE DO NOT SHARE WITH ANYONE</strong></p>
        <p>Thanks and have a nice day </p>
        <p>from <strong>Ali Team</strong>.........</p>
    </body>
</html>
    `, name, id)
	return body
}
func BookingConfirmation(req *protos.BookingEmail)string{
		body := fmt.Sprintf(`
		<html>
	<body>
		<p>Hello <strong>%v</strong>,</p>
		<p>I hope you are doing well.</p>
		<p>Booking has been confirmed successfully.</p>
		<p>-------------Details------------</p>
		<p><strong>Booking ID:</strong> %v</p>
		<p><strong>Room Type:</strong> %v</p>
		<p><strong>Total Days:</strong> %v</p>
		<p><strong>Check-In Date:</strong> %v</p>
		<p><strong>Check-Out Date:</strong> %v</p>
		<p><strong>Total Amount:</strong> %v</p>
		<p>Thanks and have a nice day.</p>
		<p>From <strong>Ali Team</strong>.........</p>
	</body>
</html>
		`, req.Username, req.BookingId, req.RoomType, req.TotalDays, req.CheckInDate, req.CheckOutDate, req.TotalAmount )
		return body
}

