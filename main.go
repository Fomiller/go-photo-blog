package main

import (
	"net/http"
	"strings"
	"text/template"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	c := getCookie(res, req)
	xs := strings.Split(c.Value, "|")

	tpl.ExecuteTemplate(res, "index.gohtml", xs)

}

func getCookie(res http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String() + "|" + "sunset.jpg" + "|" + "meadow.jpg" + "|" + "beach.jpg",
		}
		http.SetCookie(res, c)
	}
	return c
}

func appendValues(res http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// Values
	p1 := "sunset.jpg"
	p2 := "meadow.jpg"
	p3 := "beach.jpg"
	// Append
	s := c.Value
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	// Set
	c.Value = s
	http.SetCookie(res, c)
	return c
}
