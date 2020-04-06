package main

import (
	"github.com/jaakidup/reactor-core/transport"
)

func main() {
	webServer := transport.NewWebServer("HTTP:UserService", ":8081")
	webServer.Serve()
}
