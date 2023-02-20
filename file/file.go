package util

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

func LoadJSON(filePath string, target interface{}) error {
	if content, err := os.ReadFile(filePath); err != nil {
		return fmt.Errorf("read json file fail, err: %s, path: %s", err.Error(), filePath)
	} else if err = jsoniter.Unmarshal(content, &target); err != nil {
		return fmt.Errorf("unmarshal json file fail, err: %s, path: %s", err.Error(), filePath)
	}
	return nil
}
