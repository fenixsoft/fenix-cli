package prompt

import (
	"fmt"
	"strings"
)

// Filter is the type to filter the prompt.Suggestion array.
type Filter func([]Suggest, string, bool) []Suggest

// FilterHasPrefix checks whether the string completions.Text begins with sub.
func FilterHasPrefix(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, strings.HasPrefix)
}

func FilterDescHasPrefix(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestionsInDesc(completions, sub, ignoreCase, strings.HasPrefix)
}

// FilterHasSuffix checks whether the completion.Text ends with sub.
func FilterHasSuffix(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, strings.HasSuffix)
}

func FilterDescHasSuffix(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestionsInDesc(completions, sub, ignoreCase, strings.HasSuffix)
}

// FilterContains checks whether the completion.Text contains sub.
func FilterContains(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, strings.Contains)
}

func FilterDescContains(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestionsInDesc(completions, sub, ignoreCase, strings.Contains)
}

// FilterFuzzy checks whether the completion.Text fuzzy matches sub.
// Fuzzy searching for "dog" is equivalent to "*d*o*g*". This search term
// would match, for example, "Good food is gone"
//                               ^  ^      ^
func FilterFuzzy(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, fuzzyMatch)
}

func FilterDescFuzzy(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestionsInDesc(completions, sub, ignoreCase, fuzzyMatch)
}

func fuzzyMatch(s, sub string) bool {
	sChars := []rune(s)
	sIdx := 0

	// https://staticcheck.io/docs/checks#S1029
	for _, c := range sub {
		found := false
		for ; sIdx < len(sChars); sIdx++ {
			if sChars[sIdx] == c {
				found = true
				sIdx++
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func filterSuggestions(suggestions []Suggest, sub string, ignoreCase bool, function func(string, string) bool) []Suggest {
	if sub == "" {
		return suggestions
	}

	var subC string
	if ignoreCase {
		subC = strings.ToUpper(sub)
	} else {
		subC = sub
	}

	ret := make([]Suggest, 0, len(suggestions))
	for i := range suggestions {
		c := suggestions[i].Text
		if ignoreCase {
			c = strings.ToUpper(c)
		}
		if suggestions[i].Alias == sub { // MUST NOT ignore case
			ret = append(ret, Suggest{
				Text:        suggestions[i].Text,
				Alias:       suggestions[i].Alias,
				Description: fmt.Sprintf("ShortCut[%v] | %v", suggestions[i].Alias, suggestions[i].Description),
			})
		} else if function(c, subC) {
			ret = append(ret, suggestions[i])
		}
	}
	return ret
}

func filterSuggestionsInDesc(suggestions []Suggest, sub string, ignoreCase bool, function func(string, string) bool) []Suggest {
	if sub == "" {
		return suggestions
	}
	if ignoreCase {
		sub = strings.ToUpper(sub)
	}

	ret := make([]Suggest, 0, len(suggestions))
	for i := range suggestions {
		c := suggestions[i].Description
		if ignoreCase {
			c = strings.ToUpper(c)
		}
		if function(c, sub) {
			ret = append(ret, suggestions[i])
		}
	}
	return ret
}
