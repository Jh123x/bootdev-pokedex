package command

import (
	"fmt"
	"net/url"

	"github.com/Jh123x/pokedex/internal/consts"
	"github.com/Jh123x/pokedex/internal/utils"
)

func CommandExplore(args []string, _ *consts.PlayerInfo) error {
	if len(args) != 1 {
		fmt.Println("explore only takes location as the argument")
	}

	area := args[0]
	urlPath, _ := url.JoinPath(consts.BASE_URL, consts.AREA_PATH, area)

	res, err := utils.GetResult[consts.LocationInfo](urlPath)
	if err != nil {
		return err
	}

	for _, pkm := range res.PokemonEncounters {
		fmt.Println(pkm.Pokemon.Name)
	}

	return nil
}
