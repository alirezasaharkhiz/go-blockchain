package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 5000, "TCP port for blockchain web server")
	flag.Parse()
	app := NewServer(uint16(*port))
	app.Run()
}
