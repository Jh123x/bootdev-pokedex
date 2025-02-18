package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Jh123x/pokedex/internal/pokecache"
)

var (
	cache = pokecache.NewCache()
)

func GetResult[T any](currURL string) (*T, error) {
	body, ok := cache.Get(currURL)
	if !ok {
		resp, err := http.Get(currURL)
		if err != nil {
			return nil, err
		}

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		cache.Add(currURL, body)
	}

	var data *T
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
