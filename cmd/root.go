package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Init() {
	var toPrintUntilKilled string = "y"

	if len(os.Args) == 1 {
		toPrintUntilKilled = strings.Join(os.Args[1:], " ")
	}

	for {
		fmt.Println(toPrintUntilKilled)
	}
}
