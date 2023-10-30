/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wideScan bool
var shortScan bool

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans for the number of open TCP and UDP ports",
	Long: `Get the number of open TCP and UDP ports. 
		   This command takes one argument: hostname. 
		   It returns a the number of open TCP and UDP ports.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hostname := args[0]

		if shortScan {
			tcpResult, udpResult := InitalScan(hostname)
			fmt.Printf("Number of opened TCP: %d\n", tcpResult)
			fmt.Printf("Number of opened UDP: %d\n", udpResult)

		}

		if wideScan {
			tcpResult, udpResult := WideScan(hostname)
			fmt.Printf("Number of opened TCP: %d\n", tcpResult)
			fmt.Printf("Number of opened UDP: %d\n", udpResult)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	scanCmd.PersistentFlags().BoolVarP(&wideScan, "wideScan", "w", false, "Number of opened TCP and UDP ports ranging from port 1024 to 49152")
	scanCmd.PersistentFlags().BoolVarP(&shortScan, "shortscan", "s", false, "Number of opened TCP and UDP ports ranging from port 0 to 1023")
}
