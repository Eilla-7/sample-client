package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	var username, password string

	apiUrl := "http://localhost:8080"

	fmt.Println("Welcome to our app")
	fmt.Println("Login")
	fmt.Println("username: ")
	fmt.Scanln(&username)
	fmt.Println("password: ")
	fmt.Scanln(&password)

	loginData, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	resp, err := http.Post(apiUrl+"/auth", "application/json", bytes.NewBuffer(loginData))
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("Authentication failed")
		return
	}

	var authRes map[string]string
	json.NewDecoder(resp.Body).Decode(&authRes)
	token := authRes["token"]

	req, _ := http.NewRequest("GET", apiUrl+"/query", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	queryResp, err := client.Do(req)
	if err != nil {
		fmt.Println("Query request failed")
		return
	}
	defer queryResp.Body.Close()

	body, _ := io.ReadAll(queryResp.Body)
	fmt.Println("Response from Service:", string(body))
}
