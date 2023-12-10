package builtins

import (
	"fmt"
	"strings"
)

func Echo(args ...string) error {
	// join the arguments
	message := strings.Join(args, " ")

	// print results
	fmt.Println(message)
	return nil
}
