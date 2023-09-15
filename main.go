package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	URL_POST string = "http://localhost:8081/server"
)

func main() {

	http.HandleFunc("/kurier", postHandler)
	http.ListenAndServe(":8080", nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {

	user := User{}
	requestBody, err := json.Marshal(user)

	resp, err := http.Post(URL_POST, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Ошибка при отправке POST-запроса:", err)
		return
	}

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}
	fmt.Println("Ответ сервера:", string(responseBytes))
}
