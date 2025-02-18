package command

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Jh123x/pokedex/internal/consts"
	"github.com/Jh123x/pokedex/internal/pokecache"
)

type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type AreaResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Area `json:"results"`
}

type URLNode struct {
	NextURL *string
	CurrURL *string
	PrevURL *string
}

var (
	nextURL, _ = url.JoinPath(consts.BASE_URL, consts.AREA_PATH)
	urlNode    = &URLNode{
		NextURL: &nextURL,
	}
	cache = pokecache.NewCache()
)

func GetPokedexMapGen(isFwd bool) func() error {
	return func() error {
		var currURL string

		if isFwd && urlNode.NextURL != nil {
			currURL = *urlNode.NextURL
		}

		if !isFwd && urlNode.PrevURL != nil {
			currURL = *urlNode.PrevURL
		}

		result, nextNode, err := getResult(currURL, cache, urlNode)
		if err != nil {
			return err
		}

		for _, loc := range result {
			fmt.Println(loc)
		}

		urlNode = nextNode
		return nil
	}
}

func getResult(currURL string, cache *pokecache.Cache, urlNode *URLNode) ([]string, *URLNode, error) {
	body, ok := cache.Get(currURL)
	if !ok {
		resp, err := http.Get(currURL)
		if err != nil {
			return nil, urlNode, err
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, urlNode, err
		}
	}

	var data *AreaResp
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, urlNode, err
	}

	acc := make([]string, 0, len(data.Results))
	for _, loc := range data.Results {
		acc = append(acc, loc.Name)
	}

	nextNode := &URLNode{
		NextURL: &data.Next,
		CurrURL: &currURL,
		PrevURL: &data.Previous,
	}

	return acc, nextNode, nil
}
