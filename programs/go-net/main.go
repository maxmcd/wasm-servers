package main

import (
	"fmt"

	wnet "github.com/maxmcd/wasabi/pkg/net"
)

func main() {
	// fmt.Println("hi")
	l := wnet.ListenTcp("127.0.0.1:8000")
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		b := make([]byte, 10)
		if ln, err := c.Read(b); err != nil {
			panic(err)
		} else {
			fmt.Println(ln, b)
		}
		if ln, err := c.Write([]byte(`HTTP/1.1 200 OK
Content-Length: 6
Content-Type: text/plain

wasabi`)); err != nil {
			panic(err)
		} else {
			fmt.Println(ln)
		}
	}
	// wnet.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("%s %s\n", r.Method, r.URL.String())
	// 	w.Write([]byte("Hello"))
	// }))
}
