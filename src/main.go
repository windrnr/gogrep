package main

import (
	"fmt"
	"minigrep/src/api"

	"os"
)

func main() {
	var config, err = api.Build(os.Args)
	if err != nil {
		fmt.Printf("Problem parsing arguments: %v\n", err)
		return
	}
	config.Run(*config)
}
