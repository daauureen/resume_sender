package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/jordan-wright/email"
)

type Report struct {
	CVURL     string `json:"cv_url"`
	Hash      string `json:"hash"`
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Timestamp string `json:"timestamp"`
}

func main() {
	cvURL := flag.String("cv-url", "", "CV URL")
	emailAddr := flag.String("email", "", "Your email")
	smtpLogin := flag.String("smtp-login", "", "SMTP login")
	smtpPassword := flag.String("smtp-password", "", "SMTP app password")
	smtpServer := flag.String("smtp-server", "smtp.gmail.com", "SMTP server")
	smtpPort := flag.String("smtp-port", "587", "SMTP port")

	flag.Parse()

	if *cvURL == "" || *emailAddr == "" || *smtpLogin == "" || *smtpPassword == "" {
		fmt.Println("все пункты должны быть указаны")
		flag.Usage()
		os.Exit(1)
	}

	// хеш
	hashBytes := sha256.Sum256([]byte(*cvURL))
	hash := hex.EncodeToString(hashBytes[:])

	// user_id = первые 8 символов + 4 случайных
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	randomPart := make([]rune, 4)
	for i := range randomPart {
		randomPart[i] = letters[rand.Intn(len(letters))]
	}
	userID := fmt.Sprintf("%s-%s", hash[:8], string(randomPart))

	// json отчёт
	report := Report{
		CVURL:     *cvURL,
		Hash:      hash,
		UserID:    userID,
		Email:     *emailAddr,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	jsonFileName := fmt.Sprintf("report_%s.json", userID)
	file, err := os.Create(jsonFileName)
	if err != nil {
		log.Fatal("ошибка при создании json файла:", err)
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(report)
	if err != nil {
		log.Fatal("oшибка при записи json:", err)
	}

	fmt.Println("json отчёт создан:", jsonFileName)

	// отправка email
	e := email.NewEmail()
	e.From = *emailAddr
	e.To = []string{"szhaisan@wtotem.com"}
	e.Subject = "ТЗ – " + userID
	e.Text = []byte("автоматическая отправка отчёта")

	_, err = e.AttachFile(jsonFileName)
	if err != nil {
		log.Fatal("не удалось прикрепить json файл ", err)
	}

	_, err = e.AttachFile("source_code.zip")
	if err != nil {
		log.Fatal("не удалось прикрепить zip файл ", err)
	}

	err = e.Send(*smtpServer+":"+*smtpPort, smtp.PlainAuth("", *smtpLogin, *smtpPassword, *smtpServer))
	if err != nil {
		log.Fatal("ошибка при отправке письма ", err)
	}

	fmt.Println("письмо отправлено, молодец!")
}
