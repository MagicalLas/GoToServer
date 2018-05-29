package main;

import (
	"net/http"
	"fmt"
)

func main(){
	picture := [100][100]int{}
	print("Server Start\n")
	fmt.Println(picture)
	
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello World"))
    })
	http.ListenAndServe(":5000", nil)
	
	print("Server End\n")
}