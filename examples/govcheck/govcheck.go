package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// "github.com/fkcyber/soup"

type PayloadForGov struct {
	txtUsername  string `json:"txtUsername"`
	txtPassword  string `json:"txtPasasword"`
	txtKey       string `json:"txtKey"`
	ddlLanguages string `json:"ddlLanguages"`
	ddlSkins     string `json:"ddlSkins"`
	loginParam   string `json:"loginParam"`
	Token        string `json:"Token"`
	AuthId       string `json:"AuthId"`
	offset       string `json:"offset"`
}

func check(url string, email string, password string) string {
	fmt.Println("Running govcheck")
	headers := map[string]string{
		"Host":             "mail.cuk.gov.mk",
		"Accept":           "*/*",
		"X-Requested-With": "XMLHttpRequest",
		"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.60 Safari/537.36",
		"Content-Type":     "application/x-www-form-urlencoded",
		"Origin":           "http://mail.cuk.gov.mk",
		"Referer":          "http://mail.cuk.gov.mk/Mondo/lang/sys/login.aspx",
		"Accept-Encoding":  "gzip, deflate, br",
		"Accept-Language":  "en-GB,en-US;q=0.9,en;q=0.8",
		"Connection":       "keep-alive",
	}
	data := PayloadForGov{
		txtUsername:  email,
		txtPassword:  password,
		txtKey:       "",
		ddlLanguages: "en",
		ddlSkins:     "Default",
		loginParam:   "SubmitLogin",
		Token:        "",
		AuthId:       "",
		offset:       "-60",
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error:", err)
		return "had error"
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return "idk yet"
}

func main() {
	var url, email, password string

	fmt.Print("Enter URL: ")
	fmt.Scanln(&url)
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&password)

	check(url, email, password)
}
