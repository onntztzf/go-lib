package file

import (
	"os"

	"github.com/gh-zhangpeng/go-lib/e"
	jsoniter "github.com/json-iterator/go"
)

func LoadJSON(filePath string, target interface{}) error {
	if content, err := os.ReadFile(filePath); err != nil {
		return e.NewErrorWithMsg("read json file fail")
	} else if err = jsoniter.Unmarshal(content, &target); err != nil {
		return e.NewErrorWithMsg("unmarshal json file fail")
	}
	return nil
}
