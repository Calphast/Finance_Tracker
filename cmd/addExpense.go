/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"finance-tracker/db"
	"fmt"

	"github.com/spf13/cobra"
)

// addExpenseCmd represents the addExpense command
var addExpenseCmd = &cobra.Command{
	Use:   "addExpense",
	Short: "Adds a expense entry to log.",
	Long: `Generates a new expense log:
	
	- log addExpense --amount <amount>`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the flag value correctly
		amountRetrieved, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			fmt.Println("Error retrieving amount flag:", err)
			return
		}

		err = db.AddTransactions("expense", amountRetrieved)
		if err != nil {
			fmt.Println("Error storing into database:", err)
			return
		}

		fmt.Printf("Expense added: $%.2f\n", amountRetrieved)
	},
}

func init() {
	logCmd.AddCommand(addExpenseCmd)
	addExpenseCmd.Flags().Float64("amount", 0.0, "Please specify amount.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addExpenseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addExpenseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
