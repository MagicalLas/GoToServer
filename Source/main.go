package main;

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "time"
    "strconv"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "<h1>Las World!</h1>\n")
}

func Server(){
    var arr = [30][30]int{}
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/A",func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
        fmt.Fprint(w, "<h1>Las World!</h1>\n")
        fmt.Fprint(w, "",arr)
    })
    router.GET("/A/:X",func (w http.ResponseWriter, r *http.Request, ps httprouter.Params){
        I := ps.ByName("X")
        X, _ := strconv.Atoi(I[0:2])
        Y, _ := strconv.Atoi(I[2:4])
        V, _ := strconv.Atoi(I[4:])
        arr[X][Y]=V
        fmt.Fprint(w, "X :", X, "\nY :", Y, "\nV :", V)
    })
    log.Fatal(http.ListenAndServe(":"+strconv.Itoa(8080), router))
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
        fmt.Println("Server Version is 1.0\r\nAuther is Wonho Ha")
    }else if command == "start"{
        fmt.Println("Server Start")
        fmt.Println("Port is 8080")
        go Server()
    }else if command == "help"{
        help()
    }else {
        fmt.Println("->wrong command")
        fmt.Println("->help command will help you")
    }
}

func help(){
    fmt.Println("       Go To Server        ")
    fmt.Println(">This Program is Wonho's Server Controller")
    fmt.Println(">Used Golang, Source is https://github.com/Las-Wonho/GoToServer")
    fmt.Println(">Server Port was 8080")
    fmt.Println(">>Command list<<")
    fmt.Println("    ->help")
    
    fmt.Println("    ->state")
    
    fmt.Println("    ->start")
    
    fmt.Println("    ->version or ver")
    
    fmt.Println("    ->end")    
}

func CLI_io(state chan int){
    var command string
    for {
        fmt.Print(">")
        fmt.Scanln(&command)
        Method(command, state)
    }
}
func wait(state chan int){
    for <-state != 0{
        time.Sleep(time.Second)
    }
}
func main() {
    Start()
    state := make(chan int,0)
    go CLI_io(state)
    wait(state)
    fmt.Println("System End")
}
