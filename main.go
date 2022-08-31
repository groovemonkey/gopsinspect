// inspect gathers information on running processes in Linux, and maybe other operating systems
package main

import (
	"fmt"
	"os"
	// "github.com/groovemonkey/gopsinspect/linux"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	platform, err := GetPlatform()
	if err != nil {
		// TODO log error
		return 1
	}

	processList, err := PsAuxDirect(platform)
	if err != nil {
		// TODO log error
		return 1
	}
	fmt.Println(processList)

	//// Homemade ps aux
	// fmt.Println("\n\nTesting ps aux view!")
	// psaux := processList.Ps()
	// fmt.Printf("%+v\n", psaux)

	return 0
}
