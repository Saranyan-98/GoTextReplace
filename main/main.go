package main

import (
	"fmt"

	"github.com/Saranyan-98/mapper/src"
)

func main() {

	getMap := src.Mapper{}

	getMap.Filename = "test.txt"
	getMap.YAMLfile = "tags.yaml"
	err := getMap.Mapper()
	if err != nil {
		fmt.Println(err)
	}
}
