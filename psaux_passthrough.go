package main

import (
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

// psAuxDirect runs 'ps auxh' command on Unix-like systems
func psAuxDirect() ([]Process, error) {
	if runtime.GOOS == "windows" {
		return nil, errors.New("unsupported platform windows passed to psAuxDirect")
	}
	bts, err := exec.Command("ps", "auxh").Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bts), "\n")
	processes := make([]Process, len(lines))

	for i, l := range lines {
		if l == "" {
			continue
		}
		// Split on whitespace
		line := strings.Fields(l)
		proc := &UnixProcess{
			// ["root"  "1"  "0.0"  "0.0"   "4492"  "3556" "pts/0"  "Ss+" "06:02"  "0:00" "bash"]
			user:              line[0],
			pid:               line[1],
			cpupercent:        line[2],
			mempercent:        line[3],
			virtualmemorysize: line[4],
			rss:               line[5],
			tty:               line[6],
			state:             line[7],
			start:             line[8],
			time:              line[9],
			command:           line[10],
		}
		processes[i] = proc
	}
	return processes, nil
}
