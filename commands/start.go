package commands

import (
	"fmt"
	"github.com/appsutra/dns-ddos-sutra/agent"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the DDS",
	Long:  `All software needs starting. So does DDS`,
	Run: func(cmd *cobra.Command, args []string) {

		shutdownChannel := makeShutdownChannel()

		failure := agent.Start()

		//we block on this channel
		select {
		case msg := <-shutdownChannel:
			fmt.Println("Signal: ", msg)

		case msg := <-failure:
			fmt.Println(msg)
		}

		agent.Stop()

		fmt.Println("DDS stopped ")
	},
}

func makeShutdownChannel() chan os.Signal {
	//channel for catching signals of interest
	signalCatchingChannel := make(chan os.Signal)

	//catch Ctrl-C and Kill -9 signals
	signal.Notify(signalCatchingChannel, syscall.SIGINT, syscall.SIGTERM)

	return signalCatchingChannel
}
