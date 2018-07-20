package controllers

import (
	"net/http"
	"fmt"
	"io"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is home page")
}

func Dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is dog")
}

func MeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bobby")
}

func UrlValue(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("test")
	io.WriteString(w, v)
}

