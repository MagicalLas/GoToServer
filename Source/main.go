package main;

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
//    "syscall"
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
func Method(command string, s chan int){
    if command == "end"{
        s<-0
    }
    if command == "state"{
        fmt.Println("State Ok....")
    }
}
func CLI_io(s chan int){
    var command string
    for {
        fmt.Print(">")
        fmt.Scanln(&command)
        Method(command, s)
    }
}
func main() {

    Start()
    s := make(chan int,0)
    go CLI_io(s)
    for <-s != 0{
    }
    fmt.Println("Sysyem End")
}
