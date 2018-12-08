package main

import(
	"stroom/router"
	"stroom/auth"
)

func main() {
	client := new(auth.AuthenticateClient)
	router.Init(client)
}