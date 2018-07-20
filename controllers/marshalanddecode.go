package controllers

import (
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"fmt"
)

type person struct {
	Fname string
	Lname string
	Desc  []string
}


func MarshalTest(w http.ResponseWriter, r *http.Request) {
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

func UnMarshalTest (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Read body first
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal and send response back
	var data person
	err = json.Unmarshal(b, &data)
	if err != nil{
		log.Println(err)
		return
	}
	output, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(output)
}

func EncodeTest(w http.ResponseWriter, r *http.Request) {
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

func DecodeTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var data person
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	// Send back data in response
	// json.NewEncoder(w).Encode(data)
	fmt.Println(data)
}
