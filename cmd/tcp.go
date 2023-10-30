/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// tcpCmd represents the tcp command
var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "Returns a opened ports on tcp",
	Long:  `Get a list of exposed ports on TCP. This command takes two arguments: hostname(s) and port(s). It returns if the port(s) are open or closed`,
	Args:  cobra.MatchAll(),
	Run: func(cmd *cobra.Command, args []string) {
		hostnames := args[0] // TODO: hostnames should take all args except for the last arg.
		ports := args[1]
		NMAPScan(hostnames, ports)
		// fmt.Println(res)
		// fmt.Println("tcp called")
	},
}

func init() {
	rootCmd.AddCommand(tcpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tcpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tcpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
