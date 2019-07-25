package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(write io.Writer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}



func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

// go run main.go
func main() {
	//Greet(os.Stdout, "Carrie")

	// 运行程序并访问 http://localhost:5000。你会看到你的 greeting 函数被使用了。
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
