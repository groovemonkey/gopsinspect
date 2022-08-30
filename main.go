// inspect gathers information on running processes in Linux, and maybe other operating systems
package main

import (
	"fmt"
	"os"
)

// Process is a generic interface for representing a system process in an OS-agnostic way
// type Process interface {
// 	// OS() string
// 	Name() string
// 	PID() int
// 	PPID() int
// 	State() string

// 	// File stuff
// 	FileHandles() []FileHandle
// 	// TODO(string versions of these, too?)
// 	UID() int
// 	GID() int

// 	// Linux Only (blank everywhere else)
// 	Umask() int
// 	Groups() []string
// 	Seccomp() int
// 	FDsize() int
// }

func main() {
	processList, err := linux.getProcessList()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(len(processList))
	fmt.Println(processList)
}
