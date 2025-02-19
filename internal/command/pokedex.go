package command

import (
	"fmt"

	"github.com/Jh123x/pokedex/internal/consts"
)

func CommandPokedex(_ []string, playerInfo *consts.PlayerInfo) error {
	fmt.Println("Your Pokedex:")
	for k := range playerInfo.CaughtPokemons {
		fmt.Printf("  - %s\n", k)
	}

	return nil
}
