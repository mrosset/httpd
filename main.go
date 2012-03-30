package main

import (
	"fmt"
	"github.com/str1ngs/util/json"
	"log"
	"net/http"
)

type Config struct {
	Host string
	Port string
	Root string
}

var config = new(Config)

func init() {
	err := json.Read(config, "/home/strings/.httpd.json")
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
