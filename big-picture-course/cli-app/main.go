package main

import (
	"flag"
	"strings"
	"fmt"
	"bufio"
	"log"
	"os"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analysed")
	level := flag.String("level", "ERROR", "Log level to search for. OPtions are DEBUG, INFO, ERROR, and Critical")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}