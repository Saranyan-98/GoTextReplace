package main

import (
	"fmt"

	"github.com/Saranyan-98/mapper"
)

func main() {

	getMap := mapper.Mapper{}

	getMap.Filename = "sample.txt"
	getMap.YAMLfile = "tags.yaml"
	err := getMap.Mapper()
	if err != nil {
		fmt.Println(err)
	}
}
