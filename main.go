package main

import (
	"log"

	"github.com/jaakidup/project/transport"
)

func main() {
	log.Println("OK Computer!")

	webServer := transport.NewWebServer("HTTP:UserService", ":8081")
	webServer.Serve()

}
