package main

import (
	"fmt"
	"github.com/str1ngs/util/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	Host string
	Port string
	Root string
}

var (
	config = new(Config)
	home   = os.Getenv("HOME")
)

func init() {
	err := json.Read(config, filepath.Join(home, ".httpd.json"))
	if err != nil {
		log.Fatal(err)
	}
}

func Root(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	switch path {
	case "./":
		path = "./index.php"
	}
	fmt.Println(path)
	if filepath.Ext(path) == ".php" {
		phpHandle(w, r, path)
		return
	}
	http.ServeFile(w, r, path)
}

func phpHandle(w http.ResponseWriter, r *http.Request, path string) {
	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	php := exec.Command("php")
	php.Stdin = fd
	php.Stdout = w
	err = php.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var (
		host = fmt.Sprintf("%s", config.Host)
		root = config.Root
	)
	if root == "" {
		root = "."
	}
	err := os.Chdir(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getwd())
	fmt.Printf("staring http://%s/\n", host)
	http.HandleFunc("/", Root)
	err = http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal(err)
	}
}
