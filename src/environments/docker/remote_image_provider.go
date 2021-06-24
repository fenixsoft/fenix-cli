package docker

import (
	"context"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/registry"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/lib/throttle"
	"time"
)

const RemoteImage = "remote_image"

func provideRemoteImageSuggestion(args ...string) []prompt.Suggest {
	l := len(args)
	if l == 0 || len(args[l-1]) <= 2 {
		return []prompt.Suggest{}
	} else {
		imageKeyword = args[l-1]
		imageQuery.Trigger()
		return lastQueryResult
	}
}

//HubResult : Wrap DockerHub API call
type HubResult struct {
	PageCount        *int                    `json:"num_pages,omitempty"`
	ResultCount      *int                    `json:"num_results,omitempty"`
	ItemCountPerPage *int                    `json:"page_size,omitempty"`
	CurrentPage      *int                    `json:"page,omitempty"`
	Query            *string                 `json:"query,omitempty"`
	Items            []registry.SearchResult `json:"results,omitempty"`
}

func imageFromContext(imageName string, count int) []registry.SearchResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ctxResponse, err := dockerClient.ImageSearch(ctx, imageName, types.ImageSearchOptions{Limit: count})
	if err != nil || ctxResponse == nil || len(ctxResponse) <= 0 {
		return nil
	}
	return ctxResponse
}

func imageFetchCompleter(imageName string, count int) []prompt.Suggest {
	var searchResult []registry.SearchResult
	if imageName != "" {
		searchResult = imageFromContext(imageName, count)
	} else {
		searchResult = nil
	}

	if searchResult == nil || len(searchResult) <= 0 {
		return []prompt.Suggest{}
	}

	var suggestions []prompt.Suggest
	for _, s := range searchResult {
		description := "Not Official"
		if s.IsOfficial {
			description = "Official"
		}
		suggestions = append(suggestions, prompt.Suggest{Text: s.Name, Description: "(" + description + ") " + s.Description})
	}
	return suggestions
}

var imageKeyword = ""
var lastQueryResult []prompt.Suggest
var imageQuery = throttle.ThrottleFunc(200*time.Millisecond, false, func() {
	lastQueryResult = imageFetchCompleter(imageKeyword, 10)
})
