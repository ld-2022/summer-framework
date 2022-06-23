package main

import (
	"embed"
	"fmt"
	"github.com/summer-framework/core"
)

//go:embed resources
var resFs embed.FS

func main() {

	application := core.Application{
		Host:  "0.0.0.0",
		Port:  8082,
		ResFs: resFs,
	}
	err := application.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}
