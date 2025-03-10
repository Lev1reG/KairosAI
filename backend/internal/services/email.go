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

	verificationLink := fmt.Sprintf("%s/auth/verify-email?token=%s", cfg.FE_URL, token)

	subject := "Verify Your Email"
	plainTextContent := "Click the link to verify your email: " + verificationLink
	htmlContent := fmt.Sprintf("<p>Click <a href='%s'>here</a> to verify your email.</p>", verificationLink)

	return SendEmail(toEmail, subject, plainTextContent, htmlContent)
}

func SendResetPasswordEmail(toEmail, token string) error {
	cfg := config.LoadConfig()

	resetPasswordLink := fmt.Sprintf("%s/auth/reset-password?token=%s", cfg.FE_URL, token)

	subject := "Request to Reset Password"
	plainTextContent := "Click the link to reset your password: " + resetPasswordLink
	htmlContent := fmt.Sprintf("<p>Click <a href='%s'>here</a> to reset your password.</p>", resetPasswordLink)

	return SendEmail(toEmail, subject, plainTextContent, htmlContent)
}
