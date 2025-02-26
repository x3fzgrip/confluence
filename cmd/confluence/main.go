package main

import (
	"fmt"

	"github.com/x3fzgrip/confluence/pkg/confluence/reader"
)

const (
	configStructName = "Config"
	path             = "config.go"
)

func main() {
	_, err := reader.ReadConfig(path, configStructName)
	if err != nil {
		fmt.Println(err)
	}
}
