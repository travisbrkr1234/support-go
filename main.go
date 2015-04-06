package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("listening...")
	r := NewLoggedRouter()
	err := http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		panic(err)
	}
}
