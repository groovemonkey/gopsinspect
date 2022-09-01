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

	// What is this? Shown on my W10 test machine but not documented here: https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.management/get-process?view=powershell-7.2#parameters
	// SI int

	username string
	// ProcessName
	name string
}

func getProcesses_windows() ([]Process, error) {
	if runtime.GOOS != "windows" {
		return nil, errors.New("non-windows platform passed to winGetProcess")
	}
	bts, err := exec.Command("powershell", "\"Get-Process | Format-List *\"").Output()
	if err != nil {
		return nil, err
	}

	// Split lines
	lines := strings.Split(string(bts), "\n")
	// Clean up headers -- the first two are garbage, I think?
	lines = lines[2:]

	pslisting := make([]Process, len(lines))
	for i, l := range lines {
		if l == "" {
			continue
		}
		// Split on whitespace
		line := strings.Fields(l)
		getpsinfo := &windowsProcess{
			// TODO show example output line here
			handles:  line[0],
			npm:      line[1],
			pm:       line[2],
			ws:       line[3],
			vm:       line[4],
			cpu:      line[5],
			id:       line[6],
			username: line[7],
			name:     line[8],
		}
		pslisting[i] = getpsinfo
	}
	return pslisting, nil
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
