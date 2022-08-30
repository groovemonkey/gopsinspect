package linux

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type FileHandle struct {
	Descriptor int
	Path       string
}

type LinuxProcess struct {
	// Basics
	Name           string // `json:"name"`
	PID            int
	PPID           int
	State          string
	TimeSinceStart time.Duration

	// A bit more advanced
	Umask   int
	Groups  []string
	Seccomp int
	FDsize  int

	// CPU Stuff
	CPUUsageCurrent int
	CPUTimeTotal    time.Duration

	// Memory Stuff
	VMPeak int
	VMSize int

	// File stuff
	FileHandles []FileHandle
	// TODO(string versions of these, too?)
	UID int
	GID int
}

func getProcessList() ([]LinuxProcess, error) {
	procDirs, err := os.ReadDir("/proc/")
	if err != nil {
		return nil, err
	}
	// only take directories that represent a PID
	validProcDirs := make([]fs.DirEntry, 0)
	for _, d := procDirs {
		if isInt(d.Name()) and d.IsDir() {
			validProcDirs = append(validProcDirs, d)
		}
	}

	processList := make([]LinuxProcess, len(validProcDirs))

	for i, d := range validProcDirs {
		fmt.Println(d.Name())
		statPath := fmt.Sprintf("/proc/%d/stat", d.Name())

		dataBytes, err := ioutil.ReadFile(statPath)
		if err != nil {
			return nil, err
		}

		proc, err := linuxProcess(string(dataBytes))
		if err != nil {
			return nil, err
		}
		processList[i] = proc
	}
	return processList, nil
}

// process parses /proc/$PID/stat fields according to man 5 proc, and returns a LinuxProcess
func linuxProcess(statData string) (LinuxProcess, error) {
	// see https://stackoverflow.com/questions/39066998/what-are-the-meaning-of-values-at-proc-pid-stat
	stat := strings.Split(statData, " ")
	process := LinuxProcess{
		PID:  forceInt(stat[0]),
		Name: stat[1],
	}
	return process, nil
}

// forceInt force-converts a string to an int. Errors get ignored
// TODO(this is a bad idea, handle the error)
func forceInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// isInt returns true if a string can be converted to a valid integer, false otherwise
func isInt(s string) bool {
	fmt.Println("isInt:", s)

	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
