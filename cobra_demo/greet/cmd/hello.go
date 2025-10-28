/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string // 用于存储 flag 的值的变量

func init() {
	rootCmd.AddCommand(helloCmd)

	// 添加 --name flag
	// &name: 将 flag 的值绑定到 name 变量
	// "name": flag 的长名称 --name
	// "n": flag 的短名称 -n
	// "": flag 的默认值
	// "Name to greet": flag 的帮助信息
	helloCmd.Flags().StringVarP(&name, "name", "n", "", "Name to greet")
}

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints a hello message",
	Long:  `Prints a friendly hello message to the console.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the HELLO PreRun hook, just before Run.")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- This is the main Run function for HELLO ---")
		if name != "" {
			fmt.Printf("Hello, %s!\n", name)
		} else {
			fmt.Println("Hello, World!")
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the HELLO PostRun hook, right after Run.")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
