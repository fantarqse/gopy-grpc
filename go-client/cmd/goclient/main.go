package main

import (
	"log"

	goclientcmd "go-client/cmd"
)

func main() {
	if err := goclientcmd.Run(); err != nil {
		log.Fatalf("failed to start go service: %s", err.Error())
	}
}
