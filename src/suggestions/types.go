package suggestions

import (
	"github.com/c-bata/go-prompt"
)

// The Completer interface is a collection of functions for the user to perform
// intelligent completion when entering the command line
type Completer interface {
	Complete(doc prompt.Document) []prompt.Suggest
}

// The Provider is data source for different types of Suggestions
type Provider func(...string) []prompt.Suggest

// The SuggestionFilter determines how to map user input information to Suggestion
type SuggestionFilter func([]string, map[string][]prompt.Suggest) []prompt.Suggest

// The Repository is a set of Provider
type Repository struct {
	privateProviders map[string]Provider
}

// Adding a Provider to Repository
func (r *Repository) Add(id string, provider Provider) {
	r.privateProviders[id] = provider
}

// Providing suggestion by a specified Provider
func (r *Repository) Provide(id string, args ...string) []prompt.Suggest {
	if v, ok := r.privateProviders[id]; ok {
		return v(args...)
	} else if v, ok := sharedProviders[id]; ok {
		return v(args...)
	} else {
		return []prompt.Suggest{}
	}
}

// The sharedProviders are registered in different environments with universality
// and can be handed over to providers used by other environments
var sharedProviders map[string]Provider

func init() {
	sharedProviders = make(map[string]Provider)
	RegisterSharedProvider(Path, providePathSuggestion)
	RegisterSharedProvider(Output, BuildFixedSelectionProvider("yaml", "json", "table", "short"))
	RegisterSharedProvider(Loglevel, BuildFixedSelectionProvider("trace", "debug", "info", "warn", "error", "fatal"))
}

func RegisterSharedProvider(id string, provider Provider) {
	sharedProviders[id] = provider
}

const (
	Option   = "option"
	Argument = "argument"
)
