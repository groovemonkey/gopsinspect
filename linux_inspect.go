package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type FileHandle struct {
	Descriptor int
	Path       string
}

type NetworkSocket struct {
	Type string
	Path string
}

// LinuxProcessList is used for calculating or displaying metrics from many LinuxProcesses
type LinuxProcessList []LinuxProcess

// LinuxProcess represents process information
type LinuxProcess struct {
	/*
		First Batch of values is from /proc/[pid]/status
		https://man7.org/linux/man-pages/man5/proc.5.html
	*/
	Name                 string
	Umask                string
	State                string
	Tgid, Ngid           int
	Pid, PPid, TracerPid int
	// TODO(make these an integer slice? 4 values each: 1000    1000    1000    1000)
	UID, GID string
	FDSize   int
	// TODO(make this an integer slice)
	Groups                       string
	NStgid, NSpid, NSpgid, NSsid int
	// TODO(memory stuff should be an int of kbs? Not sure. Is it always KB?)
	VmPeak, VmSize, VmLck, VmPin, VmHWM               string
	VmRSS, RssAnon, RssFile, RssShmem                 string
	VmData, VmStk, VmExe, VmLib, VmPTE, VmPMD, VmSwap string
	HugetlbPages                                      string
	CoreDumping                                       int // TODO(bool?)

	Threads int
	// some skipped values here
	NoNewPrivs                  int
	Seccomp                     int
	SeccompFilters              string
	SpeculationStoreBypass      string
	CPUsAllowed                 string
	CPUsAllowedList             string
	MemsAllowed                 string
	MemsAllowedList             string
	VoluntaryContextSwitches    int
	NonVoluntaryContextSwitches int

	// Additional values, not from /proc/[pid]/status
	MemPercent float32
	CPUPercent float32
	Username   string
	Groupname  string

	// File Handles
	NumOpenFileHandles    int
	NumOpenNetworkSockets int
	FileHandles           []FileHandle
	NetworkSockets        []NetworkSocket
}

func getProcessList() (LinuxProcessList, error) {
	procDirs, err := os.ReadDir("/proc/")
	if err != nil {
		return nil, err
	}
	// only take directories that represent a PID
	validProcDirs := make([]fs.DirEntry, 0)
	for _, d := range procDirs {
		if isInt(d.Name()) && d.IsDir() {
			validProcDirs = append(validProcDirs, d)
		}
	}

	processList := make(LinuxProcessList, len(validProcDirs))

	for i, d := range validProcDirs {
		fmt.Println(d.Name())
		statPath := fmt.Sprintf("/proc/%s/status", d.Name())

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
	lines := strings.Split(statData, "\n")

	// This is terrible and I'll fix it once I have a working prototype
	// Strings.Split() strands an empty item at the end of the slice (missing \n before EOF?)
	if lines[56] != "" && len(lines) != 57 {
		return LinuxProcess{}, fmt.Errorf("expecting 57 lines exactly (last one should be empty) but got %d", len(lines))
	}

	// Drop the empty last item (SO BRITTLE OMG FIX ALL OF THIS)
	lines = lines[0 : len(lines)-1]

	// Grab just what we need from each line
	data := make([]string, len(lines))
	for i, l := range lines {
		// Split the line on whitespace: "Name:		bash foo bar"
		line := strings.Fields(l)

		// Guard against lines with no value ("Name: "), len=1
		if len(line) < 2 {
			line = []string{""}
		} else {
			// join all but the first element together and add to data
			line = line[1:]
		}

		data[i] = strings.Join(line, " ")
	}

	process := LinuxProcess{
		Name:  data[0],
		Umask: data[1],

		State:     data[2],
		Tgid:      forceInt(data[3]),
		Ngid:      forceInt(data[4]),
		Pid:       forceInt(data[5]),
		PPid:      forceInt(data[6]),
		TracerPid: forceInt(data[7]),
		// TODO(make these an integer slice? 4 values each: 1000    1000    1000    1000)
		UID:    data[8],
		GID:    data[9],
		FDSize: forceInt(data[10]),
		Groups: data[11],
		NStgid: forceInt(data[12]),
		NSpid:  forceInt(data[13]),
		NSpgid: forceInt(data[14]),
		NSsid:  forceInt(data[15]),
		// TODO(memory stuff should be an int of kbs? Not sure. Is it always KB?)
		VmPeak:       data[16],
		VmSize:       data[17],
		VmLck:        data[18],
		VmPin:        data[19],
		VmHWM:        data[20],
		VmRSS:        data[21],
		RssAnon:      data[22],
		RssFile:      data[23],
		RssShmem:     data[24],
		VmData:       data[25],
		VmStk:        data[26],
		VmExe:        data[27],
		VmLib:        data[28],
		VmPTE:        data[29],
		VmPMD:        data[30],
		VmSwap:       data[31],
		HugetlbPages: data[32],
		CoreDumping:  forceInt(data[33]),

		Threads: forceInt(data[34]),
		// 12 skipped values here
		NoNewPrivs:                  forceInt(data[46]),
		Seccomp:                     forceInt(data[47]),
		SeccompFilters:              data[48],
		SpeculationStoreBypass:      data[49],
		CPUsAllowed:                 data[50],
		CPUsAllowedList:             data[51],
		MemsAllowed:                 data[52],
		MemsAllowedList:             data[53],
		VoluntaryContextSwitches:    forceInt(data[54]),
		NonVoluntaryContextSwitches: forceInt(data[55]),
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
	_, err := strconv.Atoi(s)
	return err == nil
}
