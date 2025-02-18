package command

import (
	"fmt"
	"os"
)

func CommandExit(_ []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
