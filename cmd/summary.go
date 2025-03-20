/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"finance-tracker/db"
	"fmt"

	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		income, expenses, balance, err := db.GetSummary()
		if err != nil {
			fmt.Println("Log empty, try adding income and/or expense")
			return
		}

		fmt.Printf("Total Income: $%.2f\nTotal Expenses: $%.2f\nBalance: $%.2f\n", income, expenses, balance)
	},
}

func init() {
	logCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
