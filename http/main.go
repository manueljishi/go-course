package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct {}

func main(){
	res, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	// fmt.Println(resp)
	lw := logWriter{}
	io.Copy(lw, res.Body)
}

//Custom Write method to implement the writer interface
func (logWriter) Write(bs []byte)(int, error){
	fmt.Println(string(bs))
	return len(bs), nil
}