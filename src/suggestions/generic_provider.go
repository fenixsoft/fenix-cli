package suggestions

import (
	"github.com/c-bata/go-prompt"
	"path/filepath"
	"strings"
)

const (
	Path = "path"
)

// The providePathSuggestion provide filesystem completion
func providePathSuggestion(args ...string) []prompt.Suggest {
	l := len(args)
	if l == 0 {
		return []prompt.Suggest{}
	}
	path := args[l-1]
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
