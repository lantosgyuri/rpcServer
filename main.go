package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	inMemoryCrud := NewInMemoryCrudStore()
	err := rpc.Register(inMemoryCrud)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("Listening on Port %d", 4040)
	http.Serve(listener, nil)

}
