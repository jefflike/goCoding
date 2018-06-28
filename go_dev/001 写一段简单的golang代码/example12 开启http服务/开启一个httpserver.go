package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(
		writer http.ResponseWriter,
		request *http.Request) { // 使用指针简单来说，需要引用传递就使用指针，再简单点就是值可以改变的我们就是用指针
			fmt.Fprintf(writer,
				"<h1>hello %s</h1>", request.FormValue("name"))
	})

	http.ListenAndServe(":8000", nil)
}
