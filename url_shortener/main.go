package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var mappings = make(map[string]string)
var hostName string

func main() {
	if val := os.Getenv("hostname"); val == "" {
		panic("Need hostname env value")
	}
	hostName = os.Getenv("hostname")
	rtr := mux.NewRouter()
	rtr.HandleFunc("/{.}", redirectHandler)
	rtr.HandleFunc("/", createShortURLHandler)
	rtr.HandleFunc("/create/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", rtr))
}

func createShortURLHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("create_short_url.html")
	t.Execute(w, nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	if val, ok := mappings[path[len(path)-1]]; ok {
		http.Redirect(w, r, val, http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles("link_doesnt_exist.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Resource Not Found"))
	}
	t.Execute(w, nil)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("url")
	if body == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//need some logic to test if has http/https and if not, add it to the stored url
	if body[0:4] != "http" {
		body = fmt.Sprintf("http://%v", body)
	}

	hashedURL := Hashstr(body)

	//check for collision
	i := 5
	for {
		existing := mappings[hashedURL[:i]]
		if existing == "" {
			mappings[hashedURL[:i]] = body
			break
		}
		i++
	}

	t, _ := template.ParseFiles("generated_short_url.html")
	t.Execute(w, fmt.Sprintf("http://%v/%v", hostName, hashedURL[:i]))
}

func Hashstr(Txt string) string {
	h := sha1.New()
	h.Write([]byte(Txt))
	bs := h.Sum(nil)
	sh := string(fmt.Sprintf("%x\n", bs))
	return sh
}
