// Output a string repeatedly until killed.
package cmd

import (
	"os"
	"strings"
)

func Init() {

	toPrintUntilKilled := []byte("y\n")

	if len(os.Args) > 1 {
		toPrintUntilKilled = []byte(strings.Join(os.Args[1:], " ") + "\n")
	}

	for {
		os.Stdout.Write(toPrintUntilKilled)
	}
}
