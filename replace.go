package GoTextReplace

import (
	"fmt"
	"io"
	"os"
	"strings"
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

// Validate to check whether all the keys present in the file is present in the YAML file
func (t *TextReplace) Validate() error {

	for i := range t.Tags.Names {
		_, ok := t.Keys[t.Tags.Names[i]]
		if !ok {
			return fmt.Errorf("%s does not exist in YAML file", t.Tags.Names[i])
		}

	}
	return nil
}

// Replace the Tags with corresponding Values from the YAML
func (t *TextReplace) Replace() error {

	file, err := os.Open(t.Filename)
	CheckError(err)

	fileContent, err := io.ReadAll(file)
	CheckError(err)

	fileString := string(fileContent)

	for key, value := range t.Keys {

		newContent := strings.ReplaceAll(fileString, fmt.Sprintf("{{%s}}", key), InterfaceToString(value))
		fileString = newContent
	}

	if t.OutputFileName == "" || t.OutputPath == "" {
		err = os.WriteFile(t.Filename, []byte(fileString), 0664)
		CheckError(err)
		return nil
	}

	_, err = os.ReadDir(t.OutputPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(t.OutputPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return err
		}
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", t.OutputPath, t.OutputFileName), []byte(fileString), 0664)
	CheckError(err)

	return nil
}

// Final Mapper Function which executes the logics in sync
func (t *TextReplace) Run() error {

	var err error
	t.Tags, t.FileObj, err = t.Reader()
	CheckError(err)

	t.Keys, err = t.GetYAMLValues()
	CheckError(err)

	err = t.Validate()
	CheckError(err)

	err = t.Replace()
	CheckError(err)

	return nil
}
