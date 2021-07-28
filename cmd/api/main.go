package main

import (
	"github.com/alexperezortuno/go-auth-with-jwt-redis-postgres/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
