package main

import (
	redis "github.com/figoxu/go-redis-server"
)

func main() {
	server, err := redis.NewServer(redis.DefaultConfig())
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
