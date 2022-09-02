// inspect gathers information on running processes in Linux, and maybe other operating systems
package gopsinspect

import (
	"runtime"
)

// These types are what's actually going to get exported
type Process interface {
	Name() string
	PID() string
	State() string
	StartTime() string
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
		return processes_windows()
	} else {
		return processes_unix()
	}
}
