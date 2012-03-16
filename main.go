package main

import (
	"fmt"
	"log"
	"net/http"
	"util/json"
)

type Config struct {
	Host string
	Port string
	Root string
}

var config = &Config{"saturn", "8080", "/home/strings/downloads"}

func init() {
	err := json.Write(config, "/home/strings/.httpd.json")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		host    = fmt.Sprintf("%s:%s", config.Host, config.Port)
		handler = http.FileServer(http.Dir(config.Root))
	)
	fmt.Printf("staring http://%s:%s/\n", config.Host, config.Port)
	err := http.ListenAndServe(host, handler)
	if err != nil {
		log.Fatal(err)
	}
}
