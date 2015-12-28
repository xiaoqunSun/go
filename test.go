/*package main

import "fmt"

import "encoding/json"

func main() {
	fmt.Println("Hello, 世界")
	const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday  
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
	var i string;
	i = "xxx"
	fmt.Println(i)
	i = 
`ab 
c`
	fmt.Println(i)	
	
	fmt.Println(json.Marshal)
}
*/

package main

import (
	"flag"
	"github.com/funny/link"
	"io"
)

func main() {
	var addr string

	flag.StringVar(&addr, "addr", ":10010", "echo server address")
	flag.Parse()

	server, err := link.Serve("tcp", addr, link.Bytes(link.Uint16BE))
	if err != nil {
		panic(err)
	}
	println("server start:", server.Listener().Addr().String())
	for {
		session, err := server.Accept()
		if err != nil {
			break
		}
		go io.Copy(session.Conn(), session.Conn())
	}
}
