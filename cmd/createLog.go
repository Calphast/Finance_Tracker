/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// createLogCmd represents the createLog command
var createLogCmd = &cobra.Command{
	Use:   "createLog",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createLog called")
	},
}

func init() {
	rootCmd.AddCommand(createLogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createLogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createLogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	csvFile, err := os.Create("csv/finance-log.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
}
