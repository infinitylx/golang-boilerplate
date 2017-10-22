package main

import (
	"./config"
	"./server"
	"flag"
	"fmt"
	"os"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	server.Init()
}
