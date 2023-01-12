package handlers

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

var (
	l = logrus.New()
)

func hloggerTrace(msg string, handlerName string) {
	l.WithField("handler", handlerName).Trace(msg)
}
func hloggerError(err error, handlerName string) {
	l.WithField("handler", handlerName).Error(err)
}

func InitHandlersLogger(logLevel int) {
	l.Level = logrus.Level(logLevel)
}

func Index(w http.ResponseWriter, r *http.Request) {
	hloggerTrace("Take request", "index")
	templ, err := template.New("index").ParseFiles("front/index.html")
	if err != nil {
		hloggerError(err, "index")
	}
	err = templ.ExecuteTemplate(w, "index", nil)
	if err != nil {
		hloggerError(err, "index")
	}
}

func Auth(w http.ResponseWriter, r *http.Request) {
	hloggerTrace("Take request", "auth")
	templ, err := template.New("auth").ParseFiles("front/auth.html")
	if err != nil {
		hloggerError(err, "auth")
	}
	err = templ.ExecuteTemplate(w, "auth", nil)
	if err != nil {
		hloggerError(err, "auth")
	}
}

func Reset(w http.ResponseWriter, r *http.Request) {
	hloggerTrace("Take request", "reset")
	templ, err := template.New("reset").ParseFiles("front/reset.html")
	if err != nil {
		hloggerError(err, "reset")
	}
	err = templ.ExecuteTemplate(w, "reset", nil)
	if err != nil {
		hloggerError(err, "reset")
	}
}

func Signup(w http.ResponseWriter, r *http.Request) {
	hloggerTrace("Take request", "signup")
	templ, err := template.New("signup").ParseFiles("front/signup.html")
	if err != nil {
		hloggerError(err, "signup")
	}
	err = templ.ExecuteTemplate(w, "signup", nil)
	if err != nil {
		hloggerError(err, "signup")
	}
}

func SingleBlog(w http.ResponseWriter, r *http.Request) {
	hloggerTrace("Take request", "singleBlog")
	templ, err := template.New("singleBlog").ParseFiles("front/single-blog.html")
	if err != nil {
		hloggerError(err, "singleBlog")
	}
	err = templ.ExecuteTemplate(w, "singleBlog", nil)
	if err != nil {
		hloggerError(err, "singleBlog")
	}
}
