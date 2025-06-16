package mailconfig

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func SendPrebookMail(to_mail_id string) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("MAILTRAP_TOKEN")
	templateUuid := os.Getenv("MAILTRAP_TEMPLATE_UUID") //prebook confirmation template uuid
	url := os.Getenv("MAILTRAP_API_URL")
	from_email_id := os.Getenv("FROM_EMAIL_ID") //from email id

	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
		"from":{
			"email":"%s",
			"name":"Thank You from YesFix!"},
		"to":[{
			"email":"%s"}],
		"template_uuid":"%s",
		"template_variables":{
			"user_name":"YesFix"}}`, from_email_id, to_mail_id, templateUuid))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
}

// func ConfigureAndSendMail(target string) {
// 	// This function is a placeholder for mail configuration.

// 	// For example, you can set up SMTP server details, authentication, etc.
// 	// This is where you would initialize your mail client.
// 	username := "smtp@mailtrap.io"
// 	password := "c95b350d6b1a7c123eb6480837244281"
// 	smtpHost := "live.smtp.mailtrap.io"

// 	auth := smtp.PlainAuth("", username, password, smtpHost)
// 	from := "hello@demomailtrap.co"
// 	to := []string{target}

// 	message := []byte("To: " + to[0] + "\r\n" +
// 		"From: " + from + "\r\n" +
// 		"Subject: Thank you from YesFix!\r\n" +
// 		"This is in confirmation of you pre-booking for services from YesFix!\n\nThank you for choosing us for you home and office needs!\n\nTeam YesFix")

// 	smtpUrl := smtpHost + ":587"

// 	err := smtp.SendMail(smtpUrl, auth, from, to, message)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
