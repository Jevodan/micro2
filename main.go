package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const (
	URL_POST string = "http://localhost:8081/createUser"
	URL_GET  string = "http://localhost:8081/getInfoUser"
)

type User struct {
	Name     string `json:"name"`
	LastName string `json:"last"`
	Sex      string `json:"sex"`
}

func main() {

	http.HandleFunc("/", getHandler)
	http.HandleFunc("/create", postHandler)
	http.ListenAndServe(":8080", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка при чтении тела запроса", http.StatusBadRequest)
		return
	}
	//defer r.Body.Close()

	_, err = http.Post(URL_POST, "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Println("Ошибка при отправке POST-запроса:", err)
		return
	}

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	id := queryParams.Get("id")
	params := "?id=" + id
	resp, err := http.Get(URL_GET + params)

	if err != nil {
		fmt.Println("Ошибка при отправке GET-запроса:", err)
		return
	}

	responseBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	fmt.Println("Ответ сервера:", string(responseBytes))
	fmt.Fprintf(w, "Ответ сервера: %s", string(responseBytes))

}
