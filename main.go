package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var user string
var pass string

const STATIC_URL string = "/static/"
const STATIC_ROOT string = "static/"

func main() {
	fmt.Println("Listening...")
	http.HandleFunc("/", Home)
	http.HandleFunc("/about/", About)
<<<<<<< HEAD
	http.HandleFunc("/contact/", Contact)
	http.HandleFunc("/service/", Service)
	lookup := http.HandlerFunc(lookup)
	http.Handle("/signin/", Signin(lookup))
=======

	emaillookup := http.HandlerFunc(Lookup)
	http.Handle("/signin/", Signin(Auth(emaillookup)))
	showresult := http.HandlerFunc(Display)
	http.Handle("/display/", Auth(showresult))
	http.Handle("/lookup/", Auth(emaillookup))
	membernews := http.HandlerFunc(News)
	http.Handle("/membernews/", Auth(membernews))
>>>>>>> origin/master
	http.HandleFunc("/register/", Register)
	http.HandleFunc("/verify/", Verify)
	http.HandleFunc(STATIC_URL, StaticHandler)
	http.ListenAndServe(GetPort(), nil)

}

type Context struct {
	Title  string
	Static string
	User   string
}

func Home(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Welcome to Siaria"}
	render(w, "home", context)
}

func About(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "About Company"}
	render(w, "about", context)
}

func Contact(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Contact us"}
	render(w, "contact", context)
}

func Service(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Service"}
	render(w, "service", context)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = STATIC_URL
	tmpl_list := []string{"templates/base.html",
		fmt.Sprintf("templates/%s.html", tmpl)}
	t, err := template.ParseFiles(tmpl_list...)
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, context)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Info: No port detected in the environment, defaulting to :" + port)

	return ":" + port
}
