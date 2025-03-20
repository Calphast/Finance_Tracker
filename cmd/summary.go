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
	Short: "Returns a summary of all entries in log",
	Long: `returns a summary of all entries, the total income and the total expense,
	in addition to total balance
	
	- log summary`,
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
