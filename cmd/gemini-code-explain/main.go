package main

import (
	"flag"
	"log"

	"github.com/emicklei/gemini-code-explain/golang"
)

// gemini-code-explain -gopkg github.com/emicklei/dot@v1.6.2
func main() {
	gopkg := flag.String("gopkg", "", "go package w/o version")
	flag.Parse()
	if *gopkg != "" {
		if err := golang.Explain(*gopkg); err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
}
