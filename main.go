package main

import (
	"net/http"
	"log"
	"awesomeProject/controllers"
)


func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/dog/", controllers.Dog)
	http.HandleFunc("/me/", controllers.MeRoute)
	http.HandleFunc("/handle/", controllers.UrlValue)
	http.HandleFunc("/jsontest/", controllers.MarshalTest)
	http.HandleFunc("/encodetest/", controllers.EncodeTest)
	http.HandleFunc("/unmarshaltest/", controllers.UnMarshalTest)
	http.HandleFunc("/decodetest/", controllers.DecodeTest)
	http.HandleFunc("/signup/", controllers.Signup)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
