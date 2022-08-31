package main

import (
	"os/exec"
	"strings"
)

func PsAuxDirect() (PsListing, error) {
	cmd := "ps"
	args := "auxh"
	bts, err := exec.Command(cmd, args).Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bts), "\n")
	pslisting := make(PsListing, len(lines))

	for i, l := range lines {
		if l == "" {
			continue
		}
		// Split on whitespace
		line := strings.Fields(l)
		psinfo := PsInfo{
			// ["root"  "1"  "0.0"  "0.0"   "4492"  "3556" "pts/0"  "Ss+" "06:02"  "0:00" "bash"]
			User:              line[0],
			PID:               forceInt(line[1]),
			CPUPercent:        line[2],
			MemPercent:        line[3],
			VirtualMemorySize: line[4],
			RSS:               line[5],
			TTY:               line[6],
			State:             line[7],
			Start:             line[8],
			Time:              line[9],
			Command:           line[10],
		}
		pslisting[i] = psinfo
	}
	return pslisting, nil
}
