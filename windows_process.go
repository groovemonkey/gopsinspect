// This file parses Windows Get-Process output, as described in windows_getprocess_data.md
// It does not pull *all* available data, just enough to mimic the data returned by 'ps' on Unix-like systems

package gopsinspect

import (
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

type windowsProcess struct {
	// The number of handles that the process has opened.
	handles string
	// Non-paged memory pool (kilobytes)
	npm string
	// Paged pool (kilobytes)
	pm string
	// Working set (kilobytes)
	ws string
	// Virtual memory (in megabytes)
	vm string
	// CPU time (in seconds)
	cpu string
	// Process ID
	id string

	starttime string

	// What is this? Shown on my W10 test machine but not documented here: https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.management/get-process?view=powershell-7.2#parameters
	// SI int

	username string
	// ProcessName
	name string
}

func (p *windowsProcess) Name() string {
	return p.name
}

func (p *windowsProcess) PID() string {
	return p.id
}

func (p *windowsProcess) State() string {
	return "TODO"
}

func (p *windowsProcess) StartTime() string {
	return p.starttime
}

func (p *windowsProcess) VirtualMemorySize() string {
	return p.vm
}

func (p *windowsProcess) CPU() string {
	return "TODO"
}

func (p *windowsProcess) CPUTimeTotal() string {
	return p.cpu
}

func (p *windowsProcess) ExtraInfo() map[string]interface{} {
	return make(map[string]interface{})
}

func newWindowsProcess(data string) (*windowsProcess, error) {
	proc := windowsProcess{}
	lines := strings.Split(data, "\n")

	// Iterate through each line of the data (process "attribute : value")
	/* ALTERNATE IMPLEMENTATION, if line ordering is consistent:
	proc.foobar = windowsPsVal(lines[12], idx)
	*/
	for _, l := range lines {
		attributes := strings.Fields(l)

		// Only handling simple attr : value pairs right now
		// This also skips e.g. "attr : {val1, val2, val3}"
		if len(attributes) != 3 {
			continue
		}

		// Switch on Name
		value := attributes[2]
		switch attributes[0] {

		case "Name":
			proc.name = value
		case "Id":
			proc.id = value
		case "Handles":
			proc.handles = value
		case "NPM":
			proc.npm = value
		case "PM":
			proc.pm = value
		case "WS":
			proc.ws = value
		case "VM":
			proc.vm = value
		case "CPU":
			proc.cpu = value
		case "StartTime":
			proc.starttime = value
		case "UserName":
			proc.username = value
		}
	}
	return &proc, nil
}

func windowsProcessList(data string) ([]Process, error) {
	procs := make([]Process, 0)

	// Split off multi-line process chunks
	chunks := strings.Split(data, "\n\n")

	for _, chunk := range chunks {
		if chunk == "" || chunk == "\n" {
			continue
		}

		proc, err := newWindowsProcess(chunk)
		if err != nil {
			// TODO log error
			// Ignore process errors
			continue
		}
		procs = append(procs, proc)
	}
	return procs, nil
}

func processes_windows() ([]Process, error) {
	if runtime.GOOS != "windows" {
		return nil, errors.New("non-windows platform passed to winGetProcess")
	}
	bts, err := exec.Command("powershell", "\"Get-Process | Format-List *\"").Output()
	if err != nil {
		return nil, err
	}

	procs, err := windowsProcessList(string(bts))
	if err != nil {
		return nil, err
	}
	return procs, nil
}
