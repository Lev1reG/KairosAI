package services

import (
	"fmt"

	"github.com/Lev1reG/kairosai-backend/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toEmail, subject, plainContent, htmlContent string) error {
  cfg := config.LoadConfig()

  from := mail.NewEmail("KairosAI", cfg.SMTPUsername)
  to := mail.NewEmail("user", toEmail)
  message := mail.NewSingleEmail(from, subject, to, plainContent, htmlContent)

  client := sendgrid.NewSendClient(cfg.SendGridAPIKey)
  _, err := client.Send(message)
  if err != nil {
    return err
  }

  return nil
}

func SendVerificationEmail(toEmail, token string) error {
  cfg := config.LoadConfig()
    
  verificationLink := fmt.Sprintf("%s/api/auth/verify-email?token=%s", cfg.APP_URL, token)

	subject := "Verify Your Email"
	plainTextContent := "Click the link to verify your email: " + verificationLink
	htmlContent := fmt.Sprintf("<p>Click <a href='%s'>here</a> to verify your email.</p>", verificationLink)

	return SendEmail(toEmail, subject, plainTextContent, htmlContent)
}
