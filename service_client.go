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
	var id int

	fmt.Print("Введите id юзера:(В БД 5 пользователей) ")
	_, err := fmt.Scanln(&id)

	if err != nil || id > 5 || id < 1 {
		fmt.Println("Ошибка ввода")
		return
	}

	user := map[string]int{"id": id}
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
