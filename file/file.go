package file

import (
	"gopkg.in/yaml.v3"
	"os"

	"github.com/2hangpeng/lib/e"
	jsoniter "github.com/json-iterator/go"
)

func LoadJSON(filePath string, target interface{}) error {
	if content, err := os.ReadFile(filePath); err != nil {
		return e.NewError(e.ErrCodeSystemError, "read file fail")
	} else if err = jsoniter.Unmarshal(content, &target); err != nil {
		return e.NewError(e.ErrCodeSystemError, "unmarshal file fail")
	}
	return nil
}

// LoadYAML reads a YAML file and unmarshals it into the target interface.
func LoadYAML(filePath string, target interface{}) error {
	if content, err := os.ReadFile(filePath); err != nil {
		return e.NewError(e.ErrCodeSystemError, "read file fail")
	} else if err := yaml.Unmarshal(content, target); err != nil {
		return e.NewError(e.ErrCodeSystemError, "unmarshal file fail")
	}
	return nil
}
