package agent

import ()

var (
	dds *DDS
)

// Start the agent
func Start() {
	// Create a new instance of dds
	dds := new(DDS)
	dds.Init()
	dds.Start()
}

// stop the agent
func Stop() {
	dds.Stop()
}
