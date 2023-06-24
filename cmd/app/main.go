package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()

	if err != nil {
		fmt.Errorf("%s", err)
	}

	fmt.Printf(dir + "\\xavier")
}