package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Note    string `json:"note"`
}

var ListOfNotes = []Note{}

func SaveNote(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	n := Note{}

	if json.Unmarshal(body, &n) != nil {
		log.Fatal(err)
	}
	ListOfNotes = append(ListOfNotes, n)

	fmt.Println("Имя:", n.Name)
	fmt.Println("Фамилия:", n.Surname)
	fmt.Println("Текст заметки:", n.Note)
}

func WatchAllNotes(rw http.ResponseWriter, req *http.Request) {
	resp, err := json.Marshal(ListOfNotes)
	if err != nil {
		log.Fatal(err)
	}
	rw.Write(resp)
}

func main() {
	http.HandleFunc("/save_note", SaveNote)
	http.HandleFunc("/watch_notes", WatchAllNotes)
	log.Fatal(http.ListenAndServe("127.0.0.1:4000", nil))
}
