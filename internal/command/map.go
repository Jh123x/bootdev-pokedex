package command

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Jh123x/pokedex/internal/consts"
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

var (
	currMap  = make([][]string, 0, 51)
	currPage = -1
)

func GetPokedexMapGen(isFwd bool) func() error {
	return func() error {
		if isFwd {
			currPage += 1
		} else {
			if currPage > 0 {
				currPage -= 1
			}
		}

		var result []string
		if currPage < len(currMap) {
			result = currMap[currPage]
		} else {
			var err error
			result, err = getResult()
			if err != nil {
				return err
			}
			currMap = append(currMap, result)
		}

		for _, loc := range result {
			fmt.Println(loc)
		}

		return nil
	}
}

var (
	currURL, _ = url.JoinPath(consts.BASE_URL, consts.AREA_PATH)
)

func getResult() ([]string, error) {
	resp, err := http.Get(currURL)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *AreaResp
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	acc := make([]string, 0, len(data.Results))
	for _, loc := range data.Results {
		acc = append(acc, loc.Name)
	}

	currURL = data.Next

	return acc, nil
}
