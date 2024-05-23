package main

import (
	"log"
	"os"

	"ascii-art-fs/programs"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		log.Fatalln("Usage: go run . [STRING] [BANNER]\nEX: go run . \"something\" standard")
	}
	if len(os.Args[1:]) <= 2 {
		programs.AsciiArtFs()
	}dfdj
}
