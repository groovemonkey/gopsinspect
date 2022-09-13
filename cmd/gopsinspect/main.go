package main

import (
	"fmt"

	"github.com/groovemonkey/gopsinspect/v2"
)

func main() {
	processes, err := gopsinspect.Processes()
	if err != nil {
		fmt.Println("Error getting processes:", err)
	}

	for _, p := range processes {
		fmt.Printf("\n%+v", p)
	}

	fmt.Println("found processes:", len(processes))

}
