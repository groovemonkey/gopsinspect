package main

// UnixProcess is analogous to one line of 'ps aux'
type UnixProcess struct {
	user string
	// UID        int
	pid        string
	cpupercent string // float32
	mempercent string // float32
	// virtual memory size of the process in KiB (1024-byte units)
	virtualmemorysize string
	// resident set size, the non-swapped physical memory that a task has used (in kilobytes)
	rss     string
	tty     string
	state   string
	start   string
	time    string
	command string
}

func getProcesses_unix() ([]Process, error) {
	return psAuxDirect()
}

func (p *UnixProcess) Name() string {
	return p.command
}
func (p *UnixProcess) PID() string {
	return p.pid
}
func (p *UnixProcess) State() string {
	return p.state
}
func (p *UnixProcess) VirtualMemorySize() string {
	return p.virtualmemorysize
}
func (p *UnixProcess) CPU() string {
	return p.cpupercent
}
func (p *UnixProcess) CPUTimeTotal() string {
	return p.time
}
func (p *UnixProcess) ExtraInfo() map[string]interface{} {
	return make(map[string]interface{})
}

// // UnixProcess creates one UnixProcess struct, analogous to one (process-specific) line of 'ps aux' output
// // TODO return optional error
// func (p LinuxProcess) UnixProcess() UnixProcess {
// 	return UnixProcess{
// 		// TODO: lookup user from ID
// 		User: p.UID,
// 		// UID:               1600, // TODO p.UID,
// 		PID:               p.Pid,
// 		CPUPercent:        "TODO",
// 		MemPercent:        "TODO",
// 		VirtualMemorySize: p.VmSize,
// 		RSS:               p.VmRSS, // TODO is this correct? Or add together the other RSS values?
// 		TTY:               "TODO",
// 		State:             p.State,
// 		Start:             "TODO",
// 		Time:              "TODO",
// 		Command:           p.Name,
// 	}
// }

// // Ps takes a LinuxProcessList and uses it to create a PsListing, which is a slice of PsInfo entries
// func (l LinuxProcessList) Ps() PsListing {
// 	pslisting := make(PsListing, len(l))
// 	for i, p := range l {
// 		pslisting[i] = p.UnixProcess()
// 	}
// 	pslisting.Sort()
// 	return pslisting
// }

// // TODO PsListing.new?

// // Sort modifies a PsListing by sorting it (currently, by CPU usage)
// func (l *PsListing) Sort() {
// 	// TODO sort by CPU usage
// }
