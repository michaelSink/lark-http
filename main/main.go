package main

import (
	"lark"
)

func main() {

	server := new(lark.Server)
	server.Address = "localhost:42069"

	server.ListenAndServe()

	return
}
