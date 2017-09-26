package main

import (
	"github.com/appsutra/dns-ddos-sutra/commands"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// make the release version and commit info available to the version command
	commands.ReleaseVersion = ReleaseVersion

	commands.Execute()
}
