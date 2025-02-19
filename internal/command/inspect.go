package command

import (
	"fmt"

	"github.com/Jh123x/pokedex/internal/consts"
)

func CommandInspect(args []string, playerInfo *consts.PlayerInfo) error {
	if len(args) != 1 {
		fmt.Println("inspect takes in a pokemon name")
		return nil
	}

	pkmName := args[0]

	res, ok := playerInfo.CaughtPokemons[pkmName]
	if !ok {
		fmt.Printf("You have not caught %s yet\n", pkmName)
		return nil
	}

	fmt.Println(res.String())
	return nil
}
