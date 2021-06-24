package environments

import "github.com/fenixsoft/fenix-cli/src/suggestions"

type RuntimeCompleter struct {
	*suggestions.GenericCompleter
	*Runtime
}
