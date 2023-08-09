package src

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func GetYAMLValues(filename string) (map[string]interface{}, error) {

	var data map[string]interface{}

	ymlFile, err := os.ReadFile(filename)
	CheckError(err)

	err = yaml.Unmarshal(ymlFile, &data)
	CheckError(err)

	return data, nil
}

func Validate(Tags Tags, Keys map[string]interface{}) error {

	for i := range Tags.Names {
		_, ok := Keys[Tags.Names[i]]
		if !ok {
			return fmt.Errorf("%s does not exist in YAML file", Tags.Names[i])
		}

	}
	return nil
}

func Replace(keys map[string]interface{}, filename string) error {

	file, err := os.Open(filename)
	CheckError(err)

	fileContent, err := io.ReadAll(file)
	CheckError(err)

	fileString := string(fileContent)

	for key, value := range keys {
		newContent := strings.ReplaceAll(fileString, fmt.Sprintf("{{%s}}", key), value.(string))
		fileString = newContent
	}

	err = os.WriteFile("new.txt", []byte(fileString), 0664)

	CheckError(err)

	return nil
}
func (m *Mapper) Mapper() error {

	var err error
	m.Tags, m.FileObj, err = Reader(m.Filename)
	CheckError(err)

	m.Keys, err = GetYAMLValues(m.YAMLfile)
	CheckError(err)

	err = Validate(m.Tags, m.Keys)
	CheckError(err)

	err = Replace(m.Keys, m.Filename)
	CheckError(err)

	return nil
}
