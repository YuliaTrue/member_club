package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"time"
)

var Members []Member

type Member struct {
	Name  string
	Email string
	Date  string
}

func GetListOfMembers() []Member {
	return Members
}

func AddMember(name, email string) {
	fmt.Println(email)
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if ok := emailRegex.MatchString(email); !ok {
		return
	}
	e, m, d := time.Now().Date()
	date := fmt.Sprintf("%v %v %v", d, m, e)
	Members = append(Members,
		Member{Name: name,
			Email: email,
			Date:  date,
		})

	fmt.Println(name, email, date)
}

func status(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func index(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("./static/index.html")
	if err != nil {
		w.WriteHeader(http.StatusOK)
	}
	tmpl.Execute(w, GetListOfMembers())
}

func newMember(w http.ResponseWriter, r *http.Request) {
	AddMember(r.PostFormValue("name"), r.PostFormValue("email"))
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func Handler(port string) {
	http.HandleFunc("/status", status)
	http.HandleFunc("/", index)
	http.HandleFunc("/new_member", newMember)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	Handler("8000")
}
