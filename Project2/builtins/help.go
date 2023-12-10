package builtins

import (
	"fmt"
)

func Help() error {
	fmt.Println("Echo: display text on the terminal")

	fmt.Println("Test: evaluate conditional expressions")

	fmt.Println("pwd: prints working directory")

	fmt.Println("Alias: create shortcuts or alternate names for commands")

	return nil
}
