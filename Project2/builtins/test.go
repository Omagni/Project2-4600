package builtins

import (
	"fmt"
)

func Test(args ...string) error {
	// Retrieve command-line arguments excluding the program name

	if args[0] > args[1] {
		fmt.Printf("%s > %s\n", args[0], args[1])
	}

	if args[0] < args[1] {
		fmt.Printf("%s < %s\n", args[0], args[1])
	}

	if args[0] == args[1] {
		fmt.Printf("%s == %s\n", args[0], args[1])
	}

	return nil
}
