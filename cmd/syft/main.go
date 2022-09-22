package main

import (
	"log"
	"github.com/anchore/syft/cmd/syft/cli"
)

func main() {
	//fmt.Println("HELLO")
	//os.Exit(0)

	cli, err := cli.New()
	if err != nil {
		log.Fatalf("error during command construction: %v", err)
	}

	if err := cli.Execute(); err != nil {
		log.Fatalf("error during command execution: %v", err)
	}
}
