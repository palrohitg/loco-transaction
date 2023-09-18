package main

import (
	"github.com/fatih/structs"
	"loco-transaction/server"
	"os"
	"strconv"
)
import (
	"log"
)

const DefaultPort = "8003"
const Port = "PORT"

func main() {
	structs.DefaultTagName = "json"
	routerR := server.Setup()
	port := DefaultPort
	if p := os.Getenv(Port); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}
	err := routerR.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not run server, error: %v", err)
	}
}
