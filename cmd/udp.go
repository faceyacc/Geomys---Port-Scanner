/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// udpCmd represents the udp command
var udpCmd = &cobra.Command{
	Use:   "udp",
	Short: "Returns a opened ports on UDP. Requies root privileges.",
	Long:  `Get a list of exposed ports on UDP. This command takes two arguments: hostname(s) and port(s). It returns if the port(s) are open or closed.`,
	Args:  cobra.MatchAll(),
	Run: func(cmd *cobra.Command, args []string) {
		hostnames := args[0]
		ports := args[1]
		UDPScan(hostnames, ports)
	},
}

func init() {
	rootCmd.AddCommand(udpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// udpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// udpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
