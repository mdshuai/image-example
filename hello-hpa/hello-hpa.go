package main

import (
	"fmt"
	"math"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	x := 0.0001
	for i := 0; i < 1000000; i++ {
		x = math.Sqrt(x)
	}
	fmt.Fprintln(w, "Hello hpa!")
}

func listenAndServe(port string) {
	fmt.Printf("hello-hpa example serving on %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	go listenAndServe(port)

	port = os.Getenv("SECOND_PORT")
	if len(port) == 0 {
		port = "8888"
	}
	go listenAndServe(port)

	select {}
}
