package main

import (
	"fmt"
	"os"
	"strings"
)

func isEmpty(s string) {
	if len(strings.TrimSpace(s)) == 0 {

		fmt.Print("Error ! Filed is empty.")
		os.Exit(0)
	}
}
