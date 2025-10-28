/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// 定义一个变量来接收持久化标志的值
	verbose bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra_demo",
	Short: "A brief description of your application",
	Long:  `A longer description...`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 这个函数会在任何子命令执行前运行
		fmt.Println("I am the ROOT PersistentPreRun hook! I run before every command.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
}
