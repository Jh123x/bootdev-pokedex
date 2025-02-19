package command

import (
	"fmt"
	"os"

	"github.com/Jh123x/pokedex/internal/consts"
)

func CommandExit(_ []string, _ *consts.PlayerInfo) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
