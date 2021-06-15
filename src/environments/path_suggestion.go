package environments

import (
	"github.com/c-bata/go-prompt"
	"path/filepath"
	"strings"
)

func GetPathSuggestion(path string) []prompt.Suggest {
	if !strings.HasSuffix(path, "*") {
		path = path + "*"
	}
	files, _ := filepath.Glob(path)
	var ret []prompt.Suggest
	for i, file := range files {
		if i > 16 {
			return ret
		} else {
			ret = append(ret, prompt.Suggest{Text: file})
		}
	}
	return ret
}
