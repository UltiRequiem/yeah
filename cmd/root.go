package cmd

import (
	"fmt"
	"os"
	"strings"
)

func Init() {
	// If you do not pass any argument
	if len(os.Args) == 1 {
		for {
			fmt.Println("y")
		}
	}

	for {
		fmt.Println(strings.Join(os.Args[1:], " "))
	}

}
