package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Note    string `json:"note"`
}

var httpClient = &http.Client{}
var basic_scan = bufio.NewScanner(os.Stdin)

func (nn *Note) UploadNote() {
	json, err := json.Marshal(nn)
	if err != nil {
		log.Fatal(err)
	}

	bb := bytes.Buffer{}
	bb.Write(json)

	req, err := http.NewRequest("POST", "http://127.0.0.1:4000/save_note", &bb)
	if err != nil {
		log.Fatal(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Вы ввели:")
		fmt.Printf("Имя - %s, Фамилия - %s, Текст заметки - %s\n", nn.Name, nn.Surname, nn.Note)
	} else {
		log.Fatal(err)
	}
}

func input() Note {
	fmt.Println("Введите, пожалуйста, ваше имя")
	basic_scan.Scan()
	name := basic_scan.Text()

	fmt.Println("Введите, пожалуйста, вашу фамилию")
	basic_scan.Scan()
	surname := basic_scan.Text()

	fmt.Println("Введите, пожалуйста, текст заметки")
	basic_scan.Scan()
	note := basic_scan.Text()

	newNote := Note{name, surname, note}
	return newNote
}

func WatchAllNotes() {
	var noteList []Note

	res, err := http.Get("http://127.0.0.1:4000/watch_notes")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if json.Unmarshal(body, &noteList) != nil {
		log.Fatal(err)
		return
	}

	for index, note := range noteList {
		fmt.Println()
		fmt.Println("Запись № ", index+1)
		fmt.Println("Имя: ", note.Name)
		fmt.Println("Фамилия: ", note.Surname)
		fmt.Println("Запись: ", note.Note)
		fmt.Println()
	}
}

func main() {
	newNote := input()
	newNote.UploadNote()
	for {
		fmt.Println()
		fmt.Println("Что следует сделать дальше?")
		fmt.Println("c -- Создать новую заметку \t l -- вывести все заметки \t q -- завершить программу")
		fmt.Println()
		basic_scan.Scan()
		choice := basic_scan.Text()
		switch choice {
		case "c":
			newNote := input()
			newNote.UploadNote()
			break
		case "l":
			WatchAllNotes()
			break
		case "q":
			os.Exit(0)
			break
		}
	}
}
