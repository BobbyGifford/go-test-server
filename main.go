package main

import (
	"net/http"
	"log"
	"io"
	"fmt"
	"encoding/json"
)

type person struct {
	Fname string
	Lname string
	Desc  []string
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is home page")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this is dog")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bobby")
}

func urlValue(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("test")
	io.WriteString(w, v)
}

func marshalTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"hunter",
		"smelly",
		[]string{"nicey", "nicey"},
	}

	json, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encodeTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	p1 := person{
		"newton",
		"gooset",
		[]string{
			"encode", "also", "works",
		},
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

func decodeTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var data person
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Println(data)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/handle/", urlValue)
	http.HandleFunc("/jsontest/", marshalTest)
	http.HandleFunc("/encodetest/", encodeTest)
	http.HandleFunc("/decodetest/", decodeTest)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
