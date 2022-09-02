package gopsinspect

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

func processes_unix() ([]Process, error) {
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

func (p *UnixProcess) StartTime() string {
	return p.start
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
