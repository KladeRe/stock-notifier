package main

import(
	"log"
	"net/smtp"
)

func sendEmail(message []byte) {

	envError := godotenv.Load()
    if envError != nil {
        fmt.Printf("Couldn't load .env file")
    }
    sender := os.Getenv("SENDER")
	receiver := os.Getenv("RECEIVER")
	password := os.Getenv("PASSWORD")


	auth := smtp.PlainAuth("", sender, password, "smtp.gmail.com")

	to := []string{receiver}

	msg := []byte("To: " + receiver + "\r\n" + "Subject: Hello from my new SMTP server\r\n" + "\r\n" + message)

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender, to, msg)

	if err != nil {

		log.Fatal(err)
		
	}
}