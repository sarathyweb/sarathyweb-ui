package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	domain   = "cloudmigo.com"
	keyPath  = "/etc/letsencrypt/live/" + domain + "/privkey.pem"
	certPath = "/etc/letsencrypt/live/" + domain + "/fullchain.pem"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		tmpl.Execute(w, nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("server is running")
	log.Fatal(http.ListenAndServe(":8051", nil))
	//go http.ListenAndServe(":80", http.HandlerFunc(redirectToHttps))

	//log.Fatal(http.ListenAndServeTLS(":443", certPath, keyPath, nil))
}

func redirectToHttps(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:443" will only work if you are accessing the server from your local machine.
	http.Redirect(w, r, "https://"+domain+r.RequestURI, http.StatusMovedPermanently)
}
