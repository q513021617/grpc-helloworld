package main

import (
	"grpc-go-helloworld/client"
	"grpc-go-helloworld/server"
	"time"
)

func main() {
	go func() {
		server.Run()
	}()
	time.Sleep(3 * time.Millisecond)
	client.Run()
}
