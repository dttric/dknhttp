package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	// "github.com/julienschmidt/httprouter"
)

type responseData map[string]interface{}

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "СТРАНИЦА БЕЗ HTML")
}

func loadTemplate(tpl string) *template.Template {
  return template.Must(template.ParseFiles(fmt.Sprintf(tpl)))
}

func main() {
	cfg, _ := os.ReadFile("./config.json")
	fmt.Println(string(cfg))
	// portNum := ":" + string(cfg.port)

	// router := httprouter.New()

	// router.GET(cfg.path, func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// 	_ = loadTemplate(cfg.path[1:]).Execute(response, responseData{})
	// })

  // _ = http.ListenAndServe(portNum, router)
	// log.Println("Запущен на порте", portNum)
  // fmt.Println("Для разноса сервера CTRL+C")
}