/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"finance-tracker/db"

	"fmt"

	"github.com/spf13/cobra"
)

// addIncomeCmd represents the addIncome command
var addIncomeCmd = &cobra.Command{
	Use:   "addIncome",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve the flag value correctly
		amountRetrieved, err := cmd.Flags().GetFloat64("amount")
		if err != nil {
			fmt.Println("Error retrieving amount flag:", err)
			return
		}

		err = db.AddTransactions("income", amountRetrieved)
		if err != nil {
			fmt.Println("Error storing into database:", err)
			return
		}

		fmt.Printf("Income added: $%.2f\n", amountRetrieved)
	},
}

func init() {
	logCmd.AddCommand(addIncomeCmd)
	addIncomeCmd.Flags().Float64("amount", 0.0, "Please specify amount.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addIncomeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addIncomeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
