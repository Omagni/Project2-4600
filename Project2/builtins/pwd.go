package builtins

import (
	"fmt"
	"os"
)

func PWD() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("cwd:", cwd)
	return nil
}
