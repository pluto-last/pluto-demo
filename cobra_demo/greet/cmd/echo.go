/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echoes the provided strings",
	Long:  `Takes a sequence of strings and prints them back to the console.`,
	// 在这里添加参数验证器
	Args: cobra.MinimumNArgs(1),

	//cobra.NoArgs: 不允许任何参数。
	//cobra.ExactArgs(n): 必须有 n 个参数。
	//cobra.MinimumNArgs(n): 至少要有 n 个参数。
	//cobra.MaximumNArgs(n): 最多只能有 n 个参数。
	//cobra.RangeArgs(min, max): 参数个数必须在 min 和 max 之间。

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
