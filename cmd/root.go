/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest-cmd",
	Short: "A CLI for making REST API calls",
	Long: `rest-cmd is a command-line tool designed to simplify
interacting with RESTful APIs directly from your terminal.

You can use it to perform GET, POST, PUT, and DELETE requests.`,
	// The Run function for the root command is often used to display
	// help information if no subcommand is provided.
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// No need to print the error here, as Cobra already does it.
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// These will be "persistent" flags, available to all subcommands.
	// Example: rootCmd.PersistentFlags().StringP("url", "u", "", "The target URL for the request")
}