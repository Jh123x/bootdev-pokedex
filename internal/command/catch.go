package command

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/Jh123x/pokedex/internal/consts"
	"github.com/Jh123x/pokedex/internal/utils"
)

func CommandCatch(args []string) error {
	if len(args) != 1 {
		fmt.Println("catch command takes in the name of a pokemon")
		return nil
	}

	pkmName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pkmName)
	urlPath, _ := url.JoinPath(consts.BASE_URL, consts.POKEMON_SPECIES_PATH, pkmName)
	res, err := utils.GetResult[consts.PokemonSpecies](urlPath)
	if err != nil {
		return err
	}

	playerRate := rand.Intn(256)
	if playerRate > res.CaptureRate {
		fmt.Printf("%s was caught!\n", pkmName)
	} else {
		fmt.Printf("%s escaped!\n", pkmName)
	}

	return nil
}
