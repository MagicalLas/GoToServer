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
    var arr = [5][5]int{}

    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/A",func (w http.ResponseWriter, r *http.Request, ps httprouter.Params){
        I := ps.ByName("name")
        X := I[0:2]
        Y := I[2:4]
        arr[X][Y]=85
    })
    router.GET("/A/:X",func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
        fmt.Fprint(w, "<h1>Las World!</h1>\n")
        fmt.Fprint(w, "",arr)
    })
    log.Fatal(http.ListenAndServe(":8080", router))
}
func Start(){
    fmt.Println("Start Successfully")
    fmt.Println("Las Server Controller Start~!")
}
func Method(command string, s chan int){
    if command == "end"{
        s<-0
    }else if command == "state"{
        fmt.Println("State Ok....")
    }else if command == "version" || command == "ver"{
        fmt.Println("Server Version is 0.3\r\nAuther is Wonho Ha")
    }else if command == "start"{
        fmt.Println("Server Start")
        go Server()
    } else {
        fmt.Println("wrong command")
        fmt.Println("help command will help you")
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
