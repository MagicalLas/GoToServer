package main;

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
//    "syscall"
    "time"
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
    fmt.Println("Las Server Controller Start~!")
}
func Method(command string, s chan int){
    if command == "end"{
        s<-0
    }
    if command == "state"{
        fmt.Println("State Ok....")
    }
    if command == "version" || command == "ver"{
        fmt.Println("Server Version is 0.3\r\nAuther is Wonho Ha")
    }
    if command == "start"{
        fmt.Println("Server Start")
    }
}
func CLI_io(state chan int){
    var command string
    for {
        fmt.Print(">")
        fmt.Scanln(&command)
        Method(command, state)
    }
}
func main() {

    Start()
    state := make(chan int,0)

    go CLI_io(state)
    for <-state != 0{
        time.Sleep(time.Second)
    }
    fmt.Println("Sysyem End")
}
