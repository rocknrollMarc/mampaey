package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	fmt.Println("Listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

// Get the Port from the environment so we  can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is  nothing in the environment.
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment varable detected, defaulting to " + port)
	}
	return ":" + port
}
