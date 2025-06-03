package extras

import (

	// "example.com/pb"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	// "strings"

	"github.com/sfreiberg/gotwilio"

	"github.com/anthdm/hollywood/actor"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func LogThis(ctx *actor.Context, Error string, Message string) {
	// var (
	// 	listenAddr = ""
	// )
	// isLocal, err := strconv.ParseBool(GetEnv("isLocal"))
	// if err != nil {
	// 	log.Fatal("ERROR: unable to convert isLocal to boolean", err)
	// }
	// if isLocal {
	// 	listenAddr = GetEnv("loggerReceive")
	// } else {
	// 	listenAddr = GetEnv("loggerReceiveDocker")
	// }
	// var (
	// 	loggerPID = actor.NewPID(listenAddr, "logger")
	// )
	// ctx.Send(loggerPID, &pb.Auth{Error: Error, Message: Message, Time: time.Now().String()})
}

func LogThisWithActor(e *actor.Engine, Error string, Message string) {
	// var (
	// 	listenAddr = ""
	// )
	// isLocal, err := strconv.ParseBool(GetEnv("isLocal"))
	// if err != nil {
	// 	log.Fatal("ERROR: unable to convert isLocal to boolean", err)
	// }
	// if isLocal {
	// 	listenAddr = GetEnv("loggerReceive")
	// } else {
	// 	listenAddr = GetEnv("loggerReceiveDocker")
	// }
	// var (
	// 	loggerPID = actor.NewPID(listenAddr, "logger")
	// )
	// e.Send(loggerPID, &pb.Auth{Error: Error, Message: Message, Time: time.Now().String()})
}

const EMAIL = ""
const PASSWORD = ""

const accountSid = ""
const authToken = ""
const sender = ""

func GetEnv(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func GetAuthReceiver() (*actor.PID, error) {
	var (
		listenAddr = ""
	)
	isLocal, err := strconv.ParseBool(GetEnv("isLocal"))
	if err != nil {
		return nil, err
	}
	if isLocal {
		listenAddr = GetEnv("authReceive")
	} else {
		listenAddr = GetEnv("authReceiveDocker")
	}
	var (
		loggerPID = actor.NewPID(listenAddr, "auth")
	)

	return loggerPID, nil
}

func GetInternalReceiver() (*actor.PID, error) {
	var (
		listenAddr = ""
	)
	isLocal, err := strconv.ParseBool(GetEnv("isLocal"))
	if err != nil {
		return nil, err
	}
	if isLocal {
		listenAddr = GetEnv("internalReceive")
	} else {
		listenAddr = GetEnv("internalReceiveDocker")
	}
	var (
		loggerPID = actor.NewPID(listenAddr, "internal-comm")
	)

	return loggerPID, nil
}

func GetChatReceiver() (*actor.PID, error) {
	var (
		listenAddr = ""
	)
	isLocal, err := strconv.ParseBool(GetEnv("isLocal"))
	if err != nil {
		return nil, err
	}
	if isLocal {
		listenAddr = GetEnv("socketReceive")
	} else {
		listenAddr = GetEnv("socketReceiveDocker")
	}
	var (
		loggerPID = actor.NewPID(listenAddr, "socket")
	)

	return loggerPID, nil
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ConvertDashesToUnderscores(input string) string {
	return strings.ReplaceAll(input, "-", "_")
}

func GetTypeForColumn(datatype string) string {
	if datatype == "Short Text" {
		return "VARCHAR(100)"
	} else if datatype == "Paragraph" {
		return "TEXT"
	} else if datatype == "Multiple choice" {
		return "TEXT"
	} else if datatype == "Yes/No" {
		return "INT"
	} else if datatype == "Checkbox" {
		return "TEXT"
	} else if datatype == "File upload" {
		return "TEXT"
	} else if datatype == "Multiple choice grid" {
		return "TEXT"
	} else if datatype == "Date" {
		return "DATE"
	} else if datatype == "Time" {
		return "TIME"
	} else if datatype == "Phone number" {
		return "VARCHAR(100)"
	} else if datatype == "Address" {
		return "VARCHAR(500)"
	} else if datatype == "Location" {
		return "VARCHAR(500)"
	} else if datatype == "Document" {
		return "TEXT"
	} else if datatype == "End screen" {
		return "INT"
	}
	return ""
}

// get secretkey
func GetSecretKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		// log.Error("empty json body")
		return nil
	}

	return jsonBody
}

func GenerateSixDigitCode() int {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(900000) + 100000
	return code
}

func ContactAdminByEmail(admin_email string, email string, message string) {
	subject := "Admin Contact | TeamUp"
	body := `
		<html>
		<body>
			<h2>Account Information</h2>
			<p>You have a query from a user.</p>
			<p>Email: ` + email + `</p>
			<p>Message: ` + message + `</p>
		</body>
		</html>`

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", EMAIL, PASSWORD, "mail.weteck.io")
	addr := "mail.weteck.io:587"

	// Compose the email message
	msg := []byte("To: " + admin_email + "\r\n" +
		"From: TeamUp <hello@weteck.co>\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")
	// Send the email
	err := smtp.SendMail(addr, auth, EMAIL, []string{admin_email}, msg)
	if err != nil {
		fmt.Println("Error in sending")
	}

	fmt.Println("sent successfully")
}

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000 // Minimum 4-digit number
	max := 9999 // Maximum 4-digit number
	otp := rand.Intn(max-min+1) + min
	return fmt.Sprintf("%04d", otp) // Format the OTP with leading zeros if necessary
}

func ForgotPassEmail(email string, message string) {
	subject := "Reset Password | TeamUp"
	body := `
		<html>
		<body>
			<h2>Forgot Password</h2>
			<p>Here are your details</p>
			<p>Email: ` + email + `</p>
			<p>Your Verify Pin to reset password is: ` + message + `</p>
		</body>
		</html>`

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", EMAIL, PASSWORD, "mail.weteck.io")
	addr := "mail.weteck.io:587"

	// Compose the email message
	msg := []byte("To: " + email + "\r\n" +
		"From: TeamUp <hello@weteck.co>\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail(addr, auth, EMAIL, []string{email}, msg)
	if err != nil {
		fmt.Println("error in sending")
	}

	fmt.Println("sent successfully")
}
func OtpEmail(email string, message string) {
	subject := "Account Verification | TeamUp"
	body := `
		<html>
		<body>
			<h2>Please verify your TeamUp account</h2>
			<p>Here are your details</p>
			<p>Email: ` + email + `</p>
			<p>Your Verify OTP is: ` + message + `</p>
			<p>Don't share your OTP with anyone.</p>
		</body>
		</html>`

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", EMAIL, PASSWORD, "mail.weteck.io")
	addr := "mail.weteck.io:587"

	// Compose the email message
	msg := []byte("To: " + email + "\r\n" +
		"From: TeamUp <hello@weteck.co>\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail(addr, auth, EMAIL, []string{email}, msg)
	if err != nil {
		fmt.Println("error in sending")
	}

	fmt.Println("sent successfully")
}
func WelcomeEmail(email string, message string) {
	subject := "Welcome to TeamUp! "
	body := fmt.Sprintf(`
	<html>
	<body>
		<pre>%s</pre>
	</body>
	</html>`, message)

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", EMAIL, PASSWORD, "mail.weteck.io")
	addr := "mail.weteck.io:587"

	// Compose the email message
	msg := []byte("To: " + email + "\r\n" +
		"From: TeamUp <hello@weteck.co>\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail(addr, auth, EMAIL, []string{email}, msg)
	if err != nil {
		fmt.Println("error in sending")
	}

	fmt.Println("sent successfully")
}

func SendMailUsingweteck(loginLink string, myusername string, mypassword string, memberid string, invitation bool, company string) {
	if !invitation {
		subject := "Login Information | TeamUp"
		body := `
			<html>
			<body>
				<h2>Login Information</h2>
				<p>We’re thrilled to announce that ` + company + ` has partnered with TeamUp, the ultimate platform for transforming team communication and collaboration. As a part of our team, you now have access to an innovative tool that will keep everyone connected and engaged. Don’t wait—download the app today and start experiencing the benefits of enhanced communication and teamwork!.</p>
				<p><a href="` + loginLink + `">` + loginLink + `</a></p>
				<p>Username: ` + myusername + `</p>
				<p>Password: ` + mypassword + `</p>
				<p>Download the app:
Play Store: https://play.google.com/store/apps/details?id=com.TeamUp.TeamUp
App Store: https://apps.apple.com/us/app/TeamUp-hospitality/id6608983142

Visit our Website: https://TeamUp.io

Cheers,
The TeamUp Team</p>
			</body>
			</html>`

		// Connect to the SMTP server
		auth := smtp.PlainAuth("", EMAIL, PASSWORD, "mail.weteck.io")
		addr := "mail.weteck.io:587"

		// Compose the email message
		msg := []byte("To: " + myusername + "\r\n" +
			"From: TeamUp <hello@weteck.co>\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body + "\r\n")

		// Send the email
		err := smtp.SendMail(addr, auth, EMAIL, []string{myusername}, msg)
		if err != nil {
			fmt.Println("Error in sending")
		}
		fmt.Println("sent successfully : " + myusername)
	} else {

		subject := "Invitation Link | TeamUp"
		body := `
			<html>
			<body>
				<h2>Invitation Link</h2>
				<p>We’re thrilled to announce that` + company + ` has partnered with TeamUp, the ultimate platform for transforming team communication and collaboration. As a part of our team, you now have access to an innovative tool that will keep everyone connected and engaged. Don’t wait—download the app today and start experiencing the benefits of enhanced communication and teamwork!.</p>
				<p><a href="` + loginLink + `">` + loginLink + `</a></p>
				<p>Email: ` + myusername + `</p>
				<p>Download the app:
Play Store: https://play.google.com/store/apps/details?id=com.TeamUp.TeamUp
App Store: https://apps.apple.com/us/app/TeamUp-hospitality/id6608983142

Visit our Website: https://TeamUp.io

Cheers,
The TeamUp Team</p>
			</body>
			</html>`

		// Connect to the SMTP server
		auth := smtp.PlainAuth("", EMAIL, PASSWORD, "mail.weteck.io")
		addr := "mail.weteck.io:587"

		// Compose the email message
		msg := []byte("To: " + myusername + "\r\n" +
			"From: TeamUp <hello@weteck.co>\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body + "\r\n")

		// Send the email
		err := smtp.SendMail(addr, auth, EMAIL, []string{myusername}, msg)
		if err != nil {
			log.Println("Failed to send email: %v\n", err)
		} else {
			fmt.Printf("Email sent successfully to: %s\n", myusername)
		}
	}
}

func SendBugReportMail(receiver string, title string, description string, user string, file string) {

	subject := "Bug Report | TeamUp"
	body := `
		<html>
		<body>
			<h2>New Bug Report</h2>
			<p>New bug has been reported in system. The user and bug details are: </p>
			<p>User Email: ` + user + `</p>
			<p>Bug Title: ` + title + `</p>
			<p>Bug Description: ` + description + `</p>
			<p>Bug Attachment: ` + file + `</p>
		</body>
		</html>`

	// Connect to the SMTP server
	auth := smtp.PlainAuth("", EMAIL, PASSWORD, "mail.weteck.io")
	addr := "mail.weteck.io:587"

	// Compose the email message
	msg := []byte("To: " + receiver + "\r\n" +
		"From: TeamUp <hello@weteck.co>\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email
	err := smtp.SendMail(addr, auth, EMAIL, []string{receiver}, msg)
	if err != nil {
		fmt.Println("Error in sending email")
	}

	fmt.Println("sent successfully")
}
func SendSMS(loginLink string, myphone string) error {

	client := gotwilio.NewTwilioClient(accountSid, authToken)

	_, _, err := client.SendSMS(sender, myphone, loginLink, "", "")
	if err != nil {
		return err
	}

	return nil
}
func SendMail(loginLink string, myusername string, mypassword string, invitation bool) {
	if !invitation {
		subject := "Login Information | TeamUp"
		body := `
			<html>
			<body>
				<h2>Login Information</h2>
				<p>You've been invited to TeamUp. Please use following Credentials to Login.</p>
				<p><a href="` + loginLink + `">` + loginLink + `</a></p>
				<p>Username: ` + myusername + `</p>
				<p>Password: ` + mypassword + `</p>
			</body>
			</html>`

		// Connect to the SMTP server
		auth := smtp.PlainAuth("", EMAIL, PASSWORD, "smtp.gmail.com")
		addr := "smtp.gmail.com:587"

		// Compose the email message
		msg := []byte("To: " + myusername + "\r\n" +
			"From: arhum\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body + "\r\n")

		// Send the email
		err := smtp.SendMail(addr, auth, EMAIL, []string{myusername}, msg)
		if err != nil {
			err.Error()
		}

	} else {

		subject := "Invitation Link | TeamUp"
		body := `
			<html>
			<body>
				<h2>Invitation Link</h2>
				<p>You've been invited to TeamUp. Please use following Link to proceed.</p>
				<p><a href="` + loginLink + `">` + loginLink + `</a></p>
				<p>Email: ` + myusername + `</p>
			</body>
			</html>`

		// Connect to the SMTP server
		auth := smtp.PlainAuth("", EMAIL, PASSWORD, "smtp.gmail.com")
		addr := "smtp.gmail.com:587"

		// Compose the email message
		msg := []byte("To: " + myusername + "\r\n" +
			"From: arhum\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body + "\r\n")

		// Send the email
		err := smtp.SendMail(addr, auth, EMAIL, []string{myusername}, msg)
		if err != nil {
			err.Error()
		}

	}
}
