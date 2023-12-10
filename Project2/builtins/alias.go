package builtins

import (
	"fmt"
	"strings"
)

func Alias(args ...string) error {

	alias := strings.Join(args, " ")
	fmt.Println("Alias: " + alias)
	return nil
}
