// Output a string repeatedly until killed.
package cmd

import (
	"os"
	"strings"
)

func Init() {
	toPrintUntilKilled := []byte(map[bool]string{true: strings.Join(os.Args[1:], " "), false: "y"}[len(os.Args) > 1] + "\n")

	for {
		os.Stdout.Write(toPrintUntilKilled)
	}
}
