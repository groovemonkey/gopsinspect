// inspect gathers information on running processes in Linux, and maybe other operating systems
package main

import (
	"fmt"
	"os"
	"runtime"
)

// These types are what's actually going to get exported
type ProcessList []Process
type Process interface {
	Name() string
	PID() string
	State() string
	// User?
	// In kilobytes
	VirtualMemorySize() string //TODO int
	// Point-in-time CPU usage
	CPU() string // TODO float32
	// Total CPU time used, seconds
	CPUTimeTotal() string // TODO int

	// Any OS-specific but useful information that we want to add to the process
	ExtraInfo() map[string]interface{}
}

func Processes() ([]Process, error) {
	if runtime.GOOS == "windows" {
		return getProcesses_windows()
	} else {
		return getProcesses_unix()
	}
}

// This is meant to be used as a library, so main() is only a quick-and-dirty test
func main() {
	os.Exit(realMain())
}

func realMain() int {
	processList, err := Processes()
	if err != nil {
		// TODO log error
		return 1
	}
	fmt.Println(processList)
	return 0
}
