package util

import (
	"fmt"
	"os"
)

func HandleErr(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
		os.Exit(1)
	}
}
