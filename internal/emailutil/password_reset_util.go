package emailutil

import (
	"Blog/bootstrap"
	"Blog/domain"
	"errors"
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendPasswordResetEmail(accessToken string, user *domain.User, env *bootstrap.Env) error{
	
	verificationLink := fmt.Sprintf("%s/reset-password?token=%s", "http://localhost:8080", accessToken)
	m := gomail.NewMessage()
	m.SetHeader("From", env.Email)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Password Reset")
	m.SetBody("text/plain", fmt.Sprintf("Click the following link to reset your password: %s", verificationLink))

	d := gomail.NewDialer("smtp.gmail.com", 587, env.Email, env.Password)
	err := d.DialAndSend(m)
	if err != nil{
		return errors.New("error sending reset password email")
	}
	return nil
	
}
