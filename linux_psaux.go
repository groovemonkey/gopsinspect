package main

// PsListing is analogous to the data returned by running 'ps aux'
type PsListing []PsInfo

// PsInfo is analogous to one line of 'ps aux'
type PsInfo struct {
	User string
	// UID        int
	PID        int
	CPUPercent string // float32
	MemPercent string // float32
	// virtual memory size of the process in KiB (1024-byte units)
	VirtualMemorySize string
	// resident set size, the non-swapped physical memory that a task has used (in kilobytes)
	RSS     string
	TTY     string
	State   string
	Start   string
	Time    string
	Command string
}

// PsInfo creates one PsInfo struct, analogous to one (process-specific) line of 'ps aux' output
// TODO return optional error
func (p LinuxProcess) PsInfo() PsInfo {
	return PsInfo{
		// TODO: lookup user from ID
		User: p.UID,
		// UID:               1600, // TODO p.UID,
		PID:               p.Pid,
		CPUPercent:        "TODO",
		MemPercent:        "TODO",
		VirtualMemorySize: p.VmSize,
		RSS:               p.VmRSS, // TODO is this correct? Or add together the other RSS values?
		TTY:               "TODO",
		State:             p.State,
		Start:             "TODO",
		Time:              "TODO",
		Command:           p.Name,
	}
}

// Ps takes a LinuxProcessList and uses it to create a PsListing, which is a slice of PsInfo entries
func (l LinuxProcessList) Ps() PsListing {
	pslisting := make(PsListing, len(l))
	for i, p := range l {
		pslisting[i] = p.PsInfo()
	}
	pslisting.Sort()
	return pslisting
}

// TODO PsListing.new?

// Sort modifies a PsListing by sorting it (currently, by CPU usage)
func (l *PsListing) Sort() {
	// TODO sort by CPU usage
}
