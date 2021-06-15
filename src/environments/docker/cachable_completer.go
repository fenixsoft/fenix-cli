package docker

import (
	"context"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/registry"
	"encoding/json"
	"github.com/c-bata/go-prompt"
	"github.com/fenixsoft/fenix-cli/lib/throttle"
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//DockerHubResult : Wrap DockerHub API call
type DockerHubResult struct {
	PageCount        *int                    `json:"num_pages,omitempty"`
	ResultCount      *int                    `json:"num_results,omitempty"`
	ItemCountPerPage *int                    `json:"page_size,omitempty"`
	CurrentPage      *int                    `json:"page,omitempty"`
	Query            *string                 `json:"query,omitempty"`
	Items            []registry.SearchResult `json:"results,omitempty"`
}

func imageFromHubAPI(count int) []registry.SearchResult {
	client := retryablehttp.NewClient()
	client.HTTPClient = &http.Client{
		Timeout: 1 * time.Second,
	}
	client.RetryWaitMin = client.HTTPClient.Timeout
	client.RetryWaitMax = client.HTTPClient.Timeout
	client.RetryMax = 3
	client.Logger = nil
	url := url.URL{
		Scheme:   "https",
		Host:     "registry.hub.docker.com",
		Path:     "/v2/repositories/library",
		RawQuery: "page=1&page_size=" + strconv.Itoa(count),
	}
	apiURL := url.String()
	response, err := client.Get(apiURL)
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	searchResult := &DockerHubResult{}
	decoder.Decode(searchResult)
	if searchResult.Items == nil || len(searchResult.Items) <= 0 {
		return nil
	}

	return searchResult.Items
}

func imageFromContext(imageName string, count int) []registry.SearchResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ctxResponse, err := dockerClient.ImageSearch(ctx, imageName, types.ImageSearchOptions{Limit: count})
	if err != nil {
		return nil
	}

	if ctxResponse == nil || len(ctxResponse) <= 0 {
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

//var memoryCache = cache.New(5*time.Minute, 10*time.Minute)
var imageKeyword = ""
var lastQueryResult []prompt.Suggest
var imageQuery = throttle.ThrottleFunc(200*time.Millisecond, false, func() {
	lastQueryResult = imageFetchCompleter(imageKeyword, 10)
})
