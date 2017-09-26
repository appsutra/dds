package commands

import (
	"github.com/spf13/cobra"
)

//MainCmd is Singularity's root command. Every other command is it's child.
var MainCmd = &cobra.Command{
	Use:   "dds",
	Short: "DDS is a DNS DDOS Protector",
	Long:  `DNS-DDOS-Protector is a DNS proxy that mitigates common DNS DDOS attacks`,
}

//Execute adds all subcommands to the root command MainCmd
func Execute() {
	AddSubcommands()
	MainCmd.Execute()
}

//AddSubcommands adds child commands to the root command MainCmd.
func AddSubcommands() {
	MainCmd.AddCommand(versionCmd)
	MainCmd.AddCommand(startCmd)
}
