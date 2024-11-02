package main

import (
	"flag"
	"log"

	"github.com/emicklei/gemini-code-explain/golang"
)

// gemini-code-explain -gopkg github.com/emicklei/dot@v1.6.2
func main() {
	gopkg := flag.String("gopkg", "", "go package w/o version")
	promptFile := flag.String("prompt", "", "use another prompt for explanation")
	flag.Parse()
	if *gopkg != "" {
		if err := golang.Explain(*gopkg, *promptFile); err != nil {
			log.Fatal("ERROR: ", err)
		}
	}
}
