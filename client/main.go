package main

import (
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error", err)
	}

	a := Item{"First", "A first item"}
	b := Item{"Second", "A second item"}
	c := Item{"Third", "A third item"}

	client.Call("RpcCrudStore.AddItem", a, &reply)
	client.Call("RpcCrudStore.AddItem", b, &reply)
	client.Call("RpcCrudStore.AddItem", c, &reply)
	client.Call("RpcCrudStore.EditItem", Item{"Second", "A new second item"}, &reply)
	client.Call("RpcCrudStore.DeleteItem", c, &reply)
}
