package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":9000"

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data := `{"Name": "NooBee", "Class": "B"}`
		dataByte := []byte(data)

		var user User

		err := json.Unmarshal(dataByte, &user)
		if err != nil {
			fmt.Println("Tambah data gagal ! \n", err.Error())
			return
		}

		fmt.Println(user)

		fmt.Println("Index dipanggil ...", user.Name)
		// rw.Write([]byte("<b>Halaman Index</b>"))
		json.NewEncoder(rw).Encode(user)
	})

	http.HandleFunc("/about", about)

	fmt.Println("server running at port", port)
	http.ListenAndServe(port, nil)
}

type User struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func about(rw http.ResponseWriter, r *http.Request) {
	users := []User{
		User{Name: "NooB", Class: "A"},
		User{Name: "Java", Class: "B"},
		User{Name: "Docker", Class: "C"},
	}
	logg := log.Default()
	logg.Printf("about called with method %v ...\n", r.Method)

	fmt.Println("bentuk awal :", users)

	usersJSON, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bentuk akhir :", string(usersJSON))

	rw.Write(usersJSON)
}
