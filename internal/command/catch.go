package command

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/Jh123x/pokedex/internal/consts"
	"github.com/Jh123x/pokedex/internal/utils"
)

func CommandCatch(args []string, playerInfo *consts.PlayerInfo) error {
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
		if err := storePokemonInfo(pkmName, playerInfo); err != nil {
			return err
		}
	} else {
		fmt.Printf("%s escaped!\n", pkmName)
	}

	return nil
}

func storePokemonInfo(pkmName string, playerInfo *consts.PlayerInfo) error {
	urlPath, _ := url.JoinPath(consts.BASE_URL, consts.POKEMON_PATH, pkmName)
	res, err := utils.GetResult[consts.PokemonInfo](urlPath)
	if err != nil {
		return err
	}

	s := &consts.Stats{}

	for _, v := range res.Stats {
		switch v.Stat.Name {
		case "speed":
			s.Speed = v.BaseStat
		case "hp":
			s.HP = v.BaseStat
		case "defense":
			s.Defense = v.BaseStat
		case "special-attack":
			s.SpecialAttack = v.BaseStat
		case "special-defense":
			s.SpecialDefense = v.BaseStat
		case "attack":
			s.Attack = v.BaseStat
		default:
			panic(v.Stat.Name)
		}
	}

	types := make([]string, 0, 3)
	for _, v := range res.Types {
		types = append(types, v.Type.Name)
	}

	playerInfo.CaughtPokemons[pkmName] = &consts.PokemonInspectInfo{
		Name:   res.Name,
		Height: res.Height,
		Weight: res.Weight,
		Stats:  *s,
		Types:  types,
	}

	return nil
}
