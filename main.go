package main

import (
	"fmt"
	"github.com/b-nova-openhub/fisicus/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}
