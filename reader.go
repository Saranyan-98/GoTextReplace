package mapper

import (
	"bufio"
	"os"
	"regexp"
)

type Tags struct {
	Names []string
}

// YAML file and source file path should be assigned to the Struct, then Mapper method needs to be called to execute the task
type Mapper struct {
	Tags                Tags
	Filename            string
	FileObj             *os.File
	YAMLfile            string
	OutputFileName      string
	OutputFileExtension string
	Keys                map[string]interface{}
}

// Read the file and get the Tags
func Reader(filename string, tagType string) (Tags, *os.File, error) {

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
