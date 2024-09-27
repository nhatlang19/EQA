package mail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"text/template"
)

type MailService interface {
	SendEmail(toEmail string, subject string, templatePath string, mailInfo any) error
	SendEmailWithAttachments(toEmail string, subject string, templatePath string, mailInfo any, fileName string, fileBuffer []byte) error
}

type MailServiceImpl struct {
}

func (s MailServiceImpl) SendEmail(toEmail string, subject string, templatePath string, mailInfo any) error {
	GMAIL_MAIL := os.Getenv("GMAIL_MAIL")
	GMAIL_PASSWORD := os.Getenv("GMAIL_PASSWORD")

	to := []string{
		toEmail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", GMAIL_MAIL, GMAIL_PASSWORD, smtpHost)

	t, _ := template.ParseFiles(templatePath)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	err := t.Execute(&body, mailInfo)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, GMAIL_MAIL, to, body.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (s MailServiceImpl) SendEmailWithAttachments(toEmail string, subject string, templatePath string, mailInfo any, fileName string, fileBuffer []byte) error {
	GMAIL_MAIL := os.Getenv("GMAIL_MAIL")
	GMAIL_PASSWORD := os.Getenv("GMAIL_PASSWORD")

	to := []string{
		toEmail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", GMAIL_MAIL, GMAIL_PASSWORD, smtpHost)

	body, err := toBytesAttachment(subject, toEmail, templatePath, fileName, fileBuffer, mailInfo)
	if err != nil {
		return err
	}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, GMAIL_MAIL, to, body)
	if err != nil {
		return err
	}
	return nil
}

func toBytesAttachment(subject, to, templatePath, fileName string, fileBuffer []byte, mailInfo any) ([]byte, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	// Header
	buf.WriteString(fmt.Sprintf("Subject: %s\n", subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", to))
	buf.WriteString("MIME-Version: 1.0\n")

	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))

	// Body with template
	buf.WriteString(fmt.Sprintf("\n--%s\n", boundary))
	buf.WriteString("Content-Type: text/html; charset=UTF-8\n\n")
	err = t.Execute(buf, mailInfo)
	if err != nil {
		return nil, err
	}

	// Attachment
	buf.WriteString(fmt.Sprintf("\n--%s\n", boundary))
	buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(fileBuffer)))
	buf.WriteString("Content-Transfer-Encoding: base64\n")
	buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\n\n", fileName))

	// Encode file by base64
	b := make([]byte, base64.StdEncoding.EncodedLen(len(fileBuffer)))
	base64.StdEncoding.Encode(b, fileBuffer)
	buf.Write(b)

	// MIME boundary
	buf.WriteString(fmt.Sprintf("\n--%s--", boundary))

	return buf.Bytes(), nil
}

func MailServiceInit() *MailServiceImpl {
	return &MailServiceImpl{}
}
