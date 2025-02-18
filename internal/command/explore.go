package command

import (
	"fmt"
	"net/url"

	"github.com/Jh123x/pokedex/internal/consts"
	"github.com/Jh123x/pokedex/internal/utils"
)

type EncounterRate struct {
	Name string `json:"walk"`
	URL  string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetail struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterRate   `json:"encounter_method"`
	VersionDetails  []VersionDetail `json:"version_details"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Encounter struct {
	Pokemon        Pokemon         `json:"pokemon"`
	VersionDetails []VersionDetail `json:"version_details"`
}

type LocationInfo struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	Location             Location              `json:"location"`
	PokemonEncounters    []Encounter           `json:"pokemon_encounters"`
}

func CommandExplore(args []string) error {
	if len(args) != 1 {
		fmt.Println("explore only takes location as the argument")
	}

	area := args[0]
	urlPath, _ := url.JoinPath(consts.BASE_URL, consts.AREA_PATH, area)

	res, err := utils.GetResult[LocationInfo](urlPath)
	if err != nil {
		return err
	}

	for _, pkm := range res.PokemonEncounters {
		fmt.Println(pkm.Pokemon.Name)
	}

	return nil
}
