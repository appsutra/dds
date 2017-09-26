package main

import (
	"github.com/appsutra/dds/commands"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// make the release version and commit info available to the version command
	commands.ReleaseVersion = ReleaseVersion

	commands.Execute()
}
