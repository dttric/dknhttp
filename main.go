package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type responseData map[string]interface{}

func loadTemplate(tpl string) *template.Template {
	return template.Must(template.ParseFiles(fmt.Sprintf("%s", tpl)))
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args
	if contains(args, "-h") || contains(args, "-r") {
		if args[1] == "-h" {
			fmt.Println("-h - помощь\n-r <./путь/до/файла.html> <порт> - запуск сервера")
		} else if args[1] == "-r" {
			portNum := ":" + string(args[3])
			router := httprouter.New()

			router.GET("/", func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
				_ = loadTemplate(args[2]).Execute(response, responseData{})
			})

			log.Println("Запущен на порте", portNum)
			fmt.Println("Для разноса сервера CTRL+C")
			_ = http.ListenAndServe(portNum, router)
		}
	} else {
		fmt.Println("Введите -h!")
	}
}
