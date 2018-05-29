package main;

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "<h1>Las World!</h1>\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "%v", ps.ByName("name"))
}
func Server(){
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
    log.Fatal(http.ListenAndServe(":8080", router))
}
func main() {
	go Server()
    var input string
	for {
		fmt.Print("> ")
		fmt.Scanln(&input)
		if input == "end"{
			break
		}
	}
}