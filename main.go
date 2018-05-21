package main

import (
	"server"
)

func main() {
	// server.Start()
	server.StartWithConfiguration("./config/sample.yml")
}
