package main

import (
	"fmt"
	"net/smtp"
)

func Email(data EmailResponseInput) {
  // smtp server configuration.
  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  // Sender data.
  from := data.SenderEmail
  password := data.Password

  fmt.Println("SENDER EMAIL: ", data.SenderEmail)
  
  // Authentication.
  auth := smtp.PlainAuth("", from, password, smtpHost)
  
  

  for i := range data.Recipients{
    // Receiver email address.
    to := []string{
      data.Recipients[i].Email,
    }

    // Message.
    message := []byte(fmt.Sprintf("To: %s\r\n" + "Subject: %s\r\n" + "\r\n" + "Hi %s,\n\n%s\n \nSincerely,\n%s", data.Recipients[i].Email, data.Subject, data.Recipients[i].Name, data.EmailBody, data.SenderName))

    // Sending email.
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

    if err != nil {
      fmt.Println("ERROR: ", err)
      return
    }

  }


  fmt.Println("Email Sent Successfully!")
}
