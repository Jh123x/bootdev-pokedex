package command

import (
	"fmt"
	"net/url"

	"github.com/Jh123x/pokedex/internal/consts"
	"github.com/Jh123x/pokedex/internal/utils"
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
)

func GetPokedexMapGen(isFwd bool) consts.Command {
	return func(_ []string, _ *consts.PlayerInfo) error {
		var currURL string

		if isFwd && urlNode.NextURL != nil {
			currURL = *urlNode.NextURL
		}

		if !isFwd && urlNode.PrevURL != nil {
			currURL = *urlNode.PrevURL
		}

		result, nextNode, err := getResult(currURL, urlNode)
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

func getResult(currURL string, urlNode *URLNode) ([]string, *URLNode, error) {
	data, err := utils.GetResult[AreaResp](currURL)
	if err != nil {
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
