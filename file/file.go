package file

import (
	"os"

	"github.com/2hangpeng/go-lib/e"
	jsoniter "github.com/json-iterator/go"
)

func LoadJSON(filePath string, target interface{}) error {
	if content, err := os.ReadFile(filePath); err != nil {
		return e.SystemError.ReplaceMsg("read json file fail")
	} else if err = jsoniter.Unmarshal(content, &target); err != nil {
		return e.SystemError.ReplaceMsg("unmarshal json file fail")
	}
	return nil
}
