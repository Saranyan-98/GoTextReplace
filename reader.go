package mapper

import (
	"bufio"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"
)

type Tags struct {
	Names []string
}

// YAML file and source file path should be assigned to the Struct, then Mapper method needs to be called to execute the task
type TextReplace struct {
	Tags           Tags
	Filename       string
	FileObj        *os.File
	YAMLfile       string
	OutputPath     string
	OutputFileName string
	Keys           map[string]interface{}
}

// Read the file and get the Tags
func (t *TextReplace) Reader() (Tags, *os.File, error) {

	file, err := os.Open(t.Filename)
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

// Get YAML Key-Values from file
func (t *TextReplace) GetYAMLValues() (map[string]interface{}, error) {

	var data map[string]interface{}

	ymlFile, err := os.ReadFile(t.YAMLfile)
	CheckError(err)

	err = yaml.Unmarshal(ymlFile, &data)
	CheckError(err)

	return data, nil
}
