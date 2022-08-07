package main

import ("net/http"
	   "fmt"
	   "os")


func main() {
	fmt.Println("Rodando aplicação Server Go")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80",nil)
}

func Hello(w http.ResponseWriter, req *http.Request) {

	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello %s. Your have %s years old. ", name, age)
}