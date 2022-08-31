// inspect gathers information on running processes in Linux, and maybe other operating systems
package main

import (
	"fmt"
	"os"
	// "github.com/groovemonkey/gopsinspect/linux"
)

func main() {
	// processList, err := getProcessList()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println("number of processes:", len(processList))
	// fmt.Println("Groups of the first process:", processList[0].Groups)
	// fmt.Println("CPUsAllowedList of the first process:", processList[0].CPUsAllowedList)
	// fmt.Println(processList)

	//// Homemade ps aux
	// fmt.Println("\n\nTesting ps aux view!")
	// psaux := processList.Ps()
	// fmt.Printf("%+v\n", psaux)

	// passthrough/exec ps aux
	fmt.Println("\n\nTesting PsAuxDirect()! (direct command execution of ps aux)")
	psauxDirect, err := PsAuxDirect()
	if err != nil {
		fmt.Println("error while calling PsAuxDirect", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", psauxDirect)
}
