package src

import (
	"bufio"
	"os"
	"regexp"
)

type Tags struct {
	Names []string
}
type Mapper struct {
	Tags     Tags
	Filename string
	FileObj  *os.File
	YAMLfile string
	Keys     map[string]interface{}
}

func Reader(filename string) (Tags, *os.File, error) {

	file, err := os.Open(filename)
	if err != nil {
		return Tags{}, file, err
	}

	scanner := bufio.NewScanner(file)

	var Tags Tags
	tagRegex := regexp.MustCompile(`\{\{(.+?)\}\}`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := tagRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			Tags.Names = append(Tags.Names, match[1])
		}
	}

	err = scanner.Err()
	if err != nil {
		return Tags, file, err
	}

	return Tags, file, nil
}
