package main

import (
	"encoding/json"
	"fmt"
	"io"
	storage "micro2/sql"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/server", postHandler)
	http.ListenAndServe(":8081", nil)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	response := string(body)
	var user User
	if err := json.Unmarshal([]byte(response), &user); err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	db := storage.InitDb()
	name, err := db.GetName(user.Id)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("пришли данные", response)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(name))
	//fmt.Fprintf(w, "Данные, полученные на сервере: %s", response)
}
