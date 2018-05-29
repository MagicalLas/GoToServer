package main;

import (
	"net/http"
)

func main(){
	arr := [100][100]int{}

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello World"))
    })
	http.ListenAndServe(":5000", nil)
	
	arr[1][2]=3
	
	print("Server Start")
	
}