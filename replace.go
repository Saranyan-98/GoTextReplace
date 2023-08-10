package mapper

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Generic Function to convert Interface to String
func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// Generic Function to Check the errors
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// Get YAML Key-Values from file
func GetYAMLValues(filename string) (map[string]interface{}, error) {

	var data map[string]interface{}

	ymlFile, err := os.ReadFile(filename)
	CheckError(err)

	err = yaml.Unmarshal(ymlFile, &data)
	CheckError(err)

	return data, nil
}

// Validate to check whether all the keys present in the file is present in the YAML file
func Validate(Tags Tags, Keys map[string]interface{}) error {

	for i := range Tags.Names {
		_, ok := Keys[Tags.Names[i]]
		if !ok {
			return fmt.Errorf("%s does not exist in YAML file", Tags.Names[i])
		}

	}
	return nil
}

// Replace the Tags with corresponding Values from the YAML
func Replace(keys map[string]interface{}, filename string) error {

	file, err := os.Open(filename)
	CheckError(err)

	fileContent, err := io.ReadAll(file)
	CheckError(err)

	fileString := string(fileContent)

	for key, value := range keys {

		newContent := strings.ReplaceAll(fileString, fmt.Sprintf("{{%s}}", key), InterfaceToString(value))
		fileString = newContent
	}

	err = os.WriteFile("new.txt", []byte(fileString), 0664)
	CheckError(err)

	return nil
}

// Final Mapper Function which executes the logics in sync
func (m *Mapper) Mapper() error {

	var err error
	m.Tags, m.FileObj, err = Reader(m.Filename, "handlebar")
	CheckError(err)

	m.Keys, err = GetYAMLValues(m.YAMLfile)
	CheckError(err)

	err = Validate(m.Tags, m.Keys)
	CheckError(err)

	err = Replace(m.Keys, m.Filename)
	CheckError(err)

	return nil
}
