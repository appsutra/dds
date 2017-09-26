package agent

import (
	"fmt"
)

var (
	dds *DDS
)

// Start the agent
func Start() chan error {
	fmt.Println("\nStarting DDS")
	// Create a new instance of dds
	dds := new(DDS)
	dds.Init()
	failure := dds.Start()
	return failure
}

// stop the agent
func Stop() {
	fmt.Println("\nStopping DDS")
	dds.Stop()
}
