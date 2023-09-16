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
func (m *Mapper) GetYAMLValues() (map[string]interface{}, error) {

	var data map[string]interface{}

	ymlFile, err := os.ReadFile(m.Filename)
	CheckError(err)

	err = yaml.Unmarshal(ymlFile, &data)
	CheckError(err)

	return data, nil
}

// Validate to check whether all the keys present in the file is present in the YAML file
func (m *Mapper) Validate() error {

	for i := range m.Tags.Names {
		_, ok := m.Keys[m.Tags.Names[i]]
		if !ok {
			return fmt.Errorf("%s does not exist in YAML file", m.Tags.Names[i])
		}

	}
	return nil
}

// Replace the Tags with corresponding Values from the YAML
func (m *Mapper) Replace() error {

	file, err := os.Open(m.Filename)
	CheckError(err)

	fileContent, err := io.ReadAll(file)
	CheckError(err)

	fileString := string(fileContent)

	for key, value := range m.Keys {

		newContent := strings.ReplaceAll(fileString, fmt.Sprintf("{{%s}}", key), InterfaceToString(value))
		fileString = newContent
	}

	if m.OutputFileName == "" || m.OutputFileExtension == "" {
		err = os.WriteFile(m.Filename, []byte(fileString), 0664)
		CheckError(err)
		return nil
	}

	err = os.WriteFile(fmt.Sprintf("%s.%s", m.OutputFileName, m.OutputFileExtension), []byte(fileString), 0664)
	CheckError(err)

	return nil
}

// Final Mapper Function which executes the logics in sync
func (m *Mapper) Mapper() error {

	var err error
	m.Tags, m.FileObj, err = m.Reader()
	CheckError(err)

	m.Keys, err = m.GetYAMLValues()
	CheckError(err)

	err = m.Validate()
	CheckError(err)

	err = m.Replace()
	CheckError(err)

	return nil
}
