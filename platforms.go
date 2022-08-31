package main

import (
	"fmt"
	"runtime"
)

type Platform struct {
	Name string
	// e.g. []{"ps", "aux"} or []{"tasklist" "windowsargs"}
	PsCommand []string
}

var Platforms = map[string]Platform{
	"windows": Platform{
		Name:      "windows",
		PsCommand: []string{"ps", "auxh"},
	},
	"linux": Platform{
		Name:      "windows",
		PsCommand: []string{"ps", "auxh"},
	},
	"darwin": Platform{
		Name:      "darwin",
		PsCommand: []string{"ps", "auxh"},
	},
}

func GetPlatform() (Platform, error) {
	p, ok := Platforms[runtime.GOOS]
	if !ok {
		return p, UnhandledPlatformError{goos: runtime.GOOS}
	}
	return p, nil
}

type UnhandledPlatformError struct {
	goos string
}

func (e UnhandledPlatformError) Error() string {
	return fmt.Sprintf("%s not found in supported platforms", runtime.GOOS)
}
