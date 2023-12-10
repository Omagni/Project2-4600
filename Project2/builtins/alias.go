package builtins

import (
	"fmt"
	"strings"
)

func Alias(args ...string) error {
	// join the arguments
	alias := strings.Join(args, " ")

	// print results
	fmt.Println("Alias: " + alias)
	return nil
}
