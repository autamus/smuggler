package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/autamus/smuggler/config"
	"github.com/autamus/smuggler/web"
)

func main() {
	http.HandleFunc("/upload", web.Upload)
	http.HandleFunc("/", web.NotFound)
	fmt.Println("Smuggler Server Started!")
	err := http.ListenAndServe(":"+strconv.FormatInt(int64(config.Global.General.Port), 10), nil)
	if err != nil {
		log.Fatal(err)
	}
}
