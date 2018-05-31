package main;

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
//    "time"
)
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "<h1>Las World!</h1>\n")
}

func Server(){
    router := httprouter.New()
    router.GET("/", Index)
    log.Fatal(http.ListenAndServe(":8080", router))
}
func Start(){
    fmt.Println("Start Successfully")
    fmt.Println("Las Server Start~!")
}
func main() {

    Start()
    
    for {
        fmt.Print(">")
        
    }
}
