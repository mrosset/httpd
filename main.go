package main

import (
	"fmt"
	"github.com/str1ngs/util/file"
	"github.com/str1ngs/util/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	Host string
	Root string
}

var (
	config = new(Config)
)

func init() {
	err := json.Read(config, filepath.Join(os.Getenv("HOME"), ".httpd.json"))
	if err != nil {
		log.Fatal(err)
	}
}

func Root(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	if path == "./" {
		switch {
		case file.Exists("./index.html"):
			path = "./index.html"
		case file.Exists("./index.php"):
			path = "./index.php"
		}
	}
	fmt.Println(r.RemoteAddr, path)
	if filepath.Ext(path) == ".php" {
		phpHandle(w, r, path)
		return
	}
	http.ServeFile(w, r, path)
}

func phpHandle(w http.ResponseWriter, r *http.Request, path string) {
	fd, err := os.Open(path)
	if err != nil {
		log.Println(err)
	}
	php := exec.Command("php")
	php.Stdin = fd
	php.Stdout = w
	err = php.Run()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	var (
		host = fmt.Sprintf("%s", config.Host)
	)
	fmt.Println(os.Getwd())
	fmt.Printf("staring http://%s/\n", host)
	http.HandleFunc("/", Root)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Println(err)
	}
}
