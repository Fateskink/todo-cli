package main

import (
	"flag"
	"fmt"
)

func main() {
	userName := flag.String("u", "", "User Name")
	password := flag.String("p", "", "Password")

	var port int
	flag.IntVar(&port, "port", 0, "DB Port")

	flag.Parse()
	fmt.Printf("Port: %d\n", port)
	fmt.Printf("username: %s\npassword: %s\n", *userName, *password)
}
