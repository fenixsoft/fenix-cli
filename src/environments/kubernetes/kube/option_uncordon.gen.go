// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var uncordonOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "--selector", Alias: "-l", Description: "Selector (label query) to filter on"},
}